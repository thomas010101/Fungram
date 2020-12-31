package btc

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"wallet/internal/hd_wallet/external/api"
	"wallet/internal/hd_wallet/types"
	"wallet/pkg/util"

	"github.com/btcsuite/btcd/chaincfg/chainhash"

	"github.com/btcsuite/btcd/btcec"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
)

type Coin struct {
	hash       *chainhash.Hash
	index      uint32
	value      int64
	pkScript   []byte
	createTime int64
}

//go:generate mockgen -destination btc_mock.go -source btc.go -package btc
type backend interface {
	Api() api.Api
}

type Btc struct {
	backend backend
	usdt    usdt
}

func New(backend backend) *Btc {
	return &Btc{
		backend: backend,
		usdt:    newUsdt(backend),
	}
}

func (b *Btc) GetAddressAt(_ types.Coin, key *hdkeychain.ExtendedKey) (string, error) {
	addr, err := addressFrom(key)
	if err != nil {
		return "", fmt.Errorf("address: %w", err)
	}
	return addr.String(), nil
}

func (b *Btc) Transfer(coin types.Coin, key *hdkeychain.ExtendedKey, amount, fee, addr string,
	isFull bool) (string, error) {
	iAmount, err := util.StringToInt64(amount, 8)
	if err != nil {
		return "", fmt.Errorf("amount(%s): %w", amount, err)
	}
	iFee, err := strconv.ParseInt(fee, 10, 64)
	if err != nil {
		return "", fmt.Errorf("parse fee: %w", err)
	}
	toAddress, err := btcutil.DecodeAddress(addr, NetParams)
	if err != nil {
		return "", fmt.Errorf("decode address(%s): %w", addr, err)
	}
	fromAddress, err := addressFrom(key)
	if err != nil {
		return "", fmt.Errorf("from address: %w", err)
	}
	private, err := privateKeyFrom(key)
	if err != nil {
		return "", fmt.Errorf("private: %w", err)
	}
	var tx *wire.MsgTx
	switch coin {
	case types.BTCCoin:
		tx, err = b.transfer(iAmount, iFee, private, toAddress, fromAddress, isFull)
	case types.OmniUSDTCoin:
		tx, err = b.usdt.transfer(iAmount, iFee, private, toAddress, fromAddress)
	default:
		err = fmt.Errorf("coin(%s) not found in the btc", coin)
	}
	if err != nil {
		return "", fmt.Errorf("generate tx: %w", err)
	}
	buf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
	if err = tx.Serialize(buf); err != nil {
		return "", fmt.Errorf("serialize: %w", err)
	}
	if err = b.backend.Api().SendTx(coin, buf.Bytes()); err != nil {
		return "", fmt.Errorf("sendTx: %w", err)
	}
	return tx.TxHash().String(), nil
}

type keyDB struct {
	address string
	private *btcec.PrivateKey
}

func newKeyDB(address btcutil.Address, key *btcec.PrivateKey) keyDB {
	return keyDB{
		address: address.String(),
		private: key,
	}
}

func (k keyDB) GetKey(addr btcutil.Address) (*btcec.PrivateKey, bool, error) {
	if k.address != addr.String() {
		return nil, false, fmt.Errorf("private(%s) not found", addr.String())
	}
	return k.private, true, nil
}

func (k keyDB) GetScript(btcutil.Address) ([]byte, error) {
	return nil, fmt.Errorf("not script")
}

func (b *Btc) transfer(iAmount, iFeeRate int64, private *btcec.PrivateKey, toAddr, fromAddr btcutil.Address,
	isFull bool) (
	*wire.MsgTx, error) {
	toPkScript, err := txscript.PayToAddrScript(toAddr)
	if err != nil {
		return nil, fmt.Errorf("payToAddrScript(%s): %w", toAddr, err)
	}
	pkScript, err := txscript.PayToAddrScript(fromAddr)
	if err != nil {
		return nil, fmt.Errorf("payFromAddrScript(%s): %w", fromAddr, err)
	}
	utxos, err := b.backend.Api().Utxos(fromAddr)
	if err != nil {
		return nil, fmt.Errorf("utxo: %w", err)
	}
	feeRate := btcutil.Amount(iFeeRate)
	var txOuts []*wire.TxOut
	if isFull {
		var amount int64
		utxos, amount, err = selectAllUtxo(feeRate, utxos)
		if err != nil {
			return nil, fmt.Errorf("select utxo: %w", err)
		}
		txOuts = []*wire.TxOut{wire.NewTxOut(amount, toPkScript)}
	} else {
		txOuts = []*wire.TxOut{wire.NewTxOut(iAmount, toPkScript)}
		amount := btcutil.Amount(iAmount)
		var remainingAmount int64
		utxos, remainingAmount, err = selectUtxo(amount, feeRate, len(txOuts), utxos)
		if err != nil {
			return nil, fmt.Errorf("fee: %w", err)
		}
		if remainingAmount > 0 {
			txOuts = append(txOuts, wire.NewTxOut(remainingAmount, pkScript))
		}
	}
	msgTx := wire.NewMsgTx(1)
	for _, v := range utxos {
		msgTx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(v.Hash, v.Index), nil, nil))
	}
	msgTx.TxOut = txOuts
	kdb := newKeyDB(fromAddr, private)
	for i, txIn := range msgTx.TxIn {
		txScript, err := txscript.SignTxOutput(NetParams, msgTx, i, pkScript,
			txscript.SigHashAll, kdb, kdb, nil)
		if err != nil {
			return nil, fmt.Errorf("signTxOutput: %w", err)
		}
		txIn.SignatureScript = txScript
	}
	return msgTx, nil
}

func addressFrom(key *hdkeychain.ExtendedKey) (btcutil.Address, error) {
	addr, err := key.Address(NetParams)
	if err != nil {
		return nil, fmt.Errorf("address of the index: %w", err)
	}
	return addr, nil
}

func privateKeyFrom(key *hdkeychain.ExtendedKey) (*btcec.PrivateKey, error) {
	return key.ECPrivKey()
}

// 获取选择的utxo, 找零金额
// return
//  0: 打包的utxo
//  1: 找零金额
//  2: 错误
func selectUtxo(amount, feeRate btcutil.Amount, outLen int, coins []api.Coin) (
	[]api.Coin, int64, error) {
	var (
		aAmount  btcutil.Amount
		fee      btcutil.Amount
		utxoSize btcutil.Amount
		isFull   bool
	)
	tmpOutLen := btcutil.Amount(outLen + 1)
	sort.SliceIsSorted(coins, func(i, j int) bool {
		return coins[i].Value < coins[j].Value
	})
	for _, v := range coins {
		utxoSize++
		aAmount += v.Value
		fee = (utxoSize*148 + 34*tmpOutLen + 10) * feeRate
		if aAmount >= fee+amount {
			isFull = true
			break
		}
	}
	if !isFull {
		return nil, 0, fmt.Errorf("balance too low, (amount: %d), feeRate(%d), "+
			"outlen(%d), coinsLen(%d)", amount, feeRate, outLen, len(coins))
	}
	return coins[:utxoSize], int64(aAmount - (fee + amount)), nil
}

// 全部转出
// return
//  0: 交易输入
//  1: 转出金额
//  2: 错误
func selectAllUtxo(feeRate btcutil.Amount, coins []api.Coin) ([]api.Coin, int64, error) {
	fee := btcutil.Amount(len(coins)*148+34*1+10) * feeRate
	var amount btcutil.Amount
	for _, v := range coins {
		amount += v.Value
	}
	if fee <= amount {
		return nil, 0, fmt.Errorf("amount too low")
	}
	return coins, int64(amount - fee), nil
}
