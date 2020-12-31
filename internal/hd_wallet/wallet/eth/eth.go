package eth

import (
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"
	"wallet/internal/hd_wallet/external/api"
	"wallet/internal/hd_wallet/types"
	"wallet/pkg/util"

	"github.com/ethereum/go-ethereum/rlp"

	"github.com/btcsuite/btcutil/hdkeychain"
	ethCommon "github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	ethParams "github.com/ethereum/go-ethereum/params"
)

const (
	transferGasLimit      = 21000
	tokenTransferGasLimit = 65000
)

//go:generate mockgen -destination eth_mock.go -source eth.go -package eth
type backend interface {
	Api() api.Api
}

type Eth struct {
	token   *Token
	backend backend
}

func New(backend backend) *Eth {
	e := &Eth{
		backend: backend,
		token:   NewToken(backend),
	}
	return e
}

func (e *Eth) GetAddressAt(_ types.Coin, key *hdkeychain.ExtendedKey) (string, error) {
	privateKey, err := privateKey(key)
	if err != nil {
		return "", fmt.Errorf("privateKey: %w", err)
	}
	return addressFromPrivate(privateKey).Hex(), nil
}

func (e *Eth) Transfer(coin types.Coin, key *hdkeychain.ExtendedKey, amount, fee, addr string,
	isFull bool) (string, error) {
	feePrice, err := util.StringToUint64(fee, 9)
	if err != nil {
		return "", fmt.Errorf("fee(%s): %w", fee, err)
	}
	gasPrice := new(big.Int).SetUint64(feePrice)
	fromPrivate, err := privateKey(key)
	if err != nil {
		return "", fmt.Errorf("privateKey: %w", err)
	}
	fromAddr := addressFromPrivate(fromPrivate)
	toAddr := ethCommon.HexToAddress(addr)
	bAmount, err := amountToBigInt(amount, coin)
	if err != nil {
		return "", fmt.Errorf("amount: %w", err)
	}
	var tx *ethTypes.Transaction
	if coin == types.ETHCoin {
		tx, err = e.transfer(fromAddr, toAddr, bAmount, gasPrice, isFull)
	} else {
		tx, err = e.token.transfer(fromAddr, toAddr, bAmount, gasPrice)
	}
	if err != nil {
		return "", fmt.Errorf("transfer, coin(%s): %w", coin, err)
	}
	tx, err = ethTypes.SignTx(
		tx,
		ethTypes.MakeSigner(ethParams.MainnetChainConfig, nil),
		fromPrivate,
	)
	if err != nil {
		return "", fmt.Errorf("signTx: %w", err)
	}
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return "", fmt.Errorf("encode: %w", err)
	}
	if err = e.backend.Api().SendTx(coin, data); err != nil {
		return "", fmt.Errorf("sendTx: %w", err)
	}
	return tx.Hash().Hex(), nil
}

func (e *Eth) transfer(from, to ethCommon.Address, amount, gasPrice *big.Int, isFull bool) (
	*ethTypes.Transaction, error) {
	nonce, err := e.backend.Api().Nonce(from)
	if err != nil {
		return nil, fmt.Errorf("nonce: %w", err)
	}
	gas := new(big.Int).Mul(big.NewInt(transferGasLimit), gasPrice)
	if amount.Cmp(gas) < 1 {
		return nil, fmt.Errorf("amount too low")
	}
	if isFull {
		amount = amount.Sub(amount, gas)
		if amount.Cmp(gas) < 1 {
			return nil, fmt.Errorf("amount too low")
		}
	}
	return ethTypes.NewTransaction(nonce, to, amount, transferGasLimit, gasPrice, nil), nil
}

func privateKey(key *hdkeychain.ExtendedKey) (*ecdsa.PrivateKey, error) {
	privateKey, err := key.ECPrivKey()
	if err != nil {
		return nil, fmt.Errorf("eth private: %w", err)
	}
	return privateKey.ToECDSA(), nil
}

func addressFromPrivate(privateKey *ecdsa.PrivateKey) ethCommon.Address {
	return ethCrypto.PubkeyToAddress(privateKey.PublicKey)
}

func amountToBigInt(amount string, coin types.Coin) (*big.Int, error) {
	var bAmount *big.Int
	decimals := coin.Decimals()
	if decimals > 9 {
		uAmount, err := util.StringToUint64(amount, 9)
		if err != nil {
			return nil, fmt.Errorf("amount(%s) wrong", amount)
		}
		bAmount = new(big.Int).Mul(
			new(big.Int).SetUint64(uAmount),
			big.NewInt(int64(math.Pow10(decimals-9))),
		)
	} else {
		uAmount, err := util.StringToUint64(amount, decimals)
		if err != nil {
			return nil, fmt.Errorf("amount(%s) wrong", amount)
		}
		bAmount = new(big.Int).SetUint64(uAmount)
	}
	if bAmount.Cmp(big0) != 1 {
		return nil, fmt.Errorf("amount too low")
	}
	return bAmount, nil
}
