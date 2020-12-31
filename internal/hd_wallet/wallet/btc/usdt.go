package btc

import (
	"encoding/binary"
	"fmt"

	"github.com/btcsuite/btcd/txscript"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

const (
	defaultAmount = 546 // btc默认最小转账金额
	propertyId    = 31
)

type usdt struct {
	backend backend
}

func newUsdt(backend backend) usdt {
	return usdt{
		backend: backend,
	}
}

func (u usdt) transfer(amount, feeRate int64, private *btcec.PrivateKey, toAddr, fromAddr btcutil.Address) (
	*wire.MsgTx, error) {
	toPkScript, err := txscript.PayToAddrScript(toAddr)
	if err != nil {
		return nil, fmt.Errorf("payToAddr: %w", err)
	}
	fromPkScript, err := txscript.PayToAddrScript(fromAddr)
	if err != nil {
		return nil, fmt.Errorf("payFromAddr: %w", err)
	}
	utxos, err := u.backend.Api().Utxos(fromAddr)
	if err != nil {
		return nil, fmt.Errorf("utxos: %w", err)
	}
	txOuts := []*wire.TxOut{
		wire.NewTxOut(defaultAmount, toPkScript),
		wire.NewTxOut(0, u.amountToScript(amount)),
	}
	utxos, remainingAmount, err := selectUtxo(645, btcutil.Amount(feeRate), len(txOuts), utxos)
	if err != nil {
		return nil, fmt.Errorf("fee: %w", err)
	}
	tx := wire.NewMsgTx(1)
	for _, v := range utxos {
		tx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(v.Hash, v.Index), nil, nil))
	}
	tx.TxOut = txOuts
	if remainingAmount > 0 {
		tx.AddTxOut(&wire.TxOut{
			Value:    int64(remainingAmount),
			PkScript: fromPkScript,
		})
	}
	kdb := newKeyDB(fromAddr, private)
	for i, txIn := range tx.TxIn {
		txScript, err := txscript.SignTxOutput(NetParams, tx, i, fromPkScript,
			txscript.SigHashAll, kdb, kdb, nil)
		if err != nil {
			return nil, fmt.Errorf("signTxOutput: %w", err)
		}
		txIn.SignatureScript = txScript
	}
	return tx, nil
}

func (u usdt) amountToScript(amount int64) []byte {
	pkScriptPrefix := make([]byte, 22)
	pkScriptPrefix[0] = 106
	pkScriptPrefix[1] = 20
	pkScriptPrefix[2] = 111
	pkScriptPrefix[3] = 109
	pkScriptPrefix[4] = 110
	pkScriptPrefix[5] = 105
	binary.BigEndian.PutUint64(pkScriptPrefix[6:14], propertyId)
	binary.BigEndian.PutUint64(pkScriptPrefix[14:22], uint64(amount))
	return pkScriptPrefix
}
