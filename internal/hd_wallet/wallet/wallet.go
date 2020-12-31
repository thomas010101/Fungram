package wallet

import (
	"fmt"
	"wallet/internal/hd_wallet/external/api"
	"wallet/internal/hd_wallet/types"
	"wallet/internal/hd_wallet/wallet/btc"
	"wallet/internal/hd_wallet/wallet/eth"

	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

type Option func(w *Wallet)

func OptWithEth() Option {
	return func(w *Wallet) {
		w.ls[types.ETHCoin] = eth.New(w)
	}
}

func OptWithBtc() Option {
	return func(w *Wallet) {
		w.ls[types.BTCCoin] = btc.New(w)
	}
}

type Walletter interface {
	GetAddressAt(coin types.Coin, key *hdkeychain.ExtendedKey) (string, error)
	Transfer(coin types.Coin, key *hdkeychain.ExtendedKey, amount, fee, addr string, isFull bool) (string, error)
}

type backend interface {
	Api() api.Api
}

type Wallet struct {
	backend backend
	ls      map[types.Coin]Walletter
}

func New(backend backend, opts ...Option) *Wallet {
	w := &Wallet{
		backend: backend,
		ls:      make(map[types.Coin]Walletter),
	}
	for _, v := range opts {
		v(w)
	}
	return w
}

func (w *Wallet) load(coin types.Coin, fun func(walletter Walletter) error) error {
	wallet, ok := w.ls[coin.Child()]
	if !ok {
		return fmt.Errorf("wallet not found")
	}
	return fun(wallet)
}

func (w *Wallet) GetAddress(mnemonic string, coin types.Coin) (string, error) {
	addressKey, err := addressKey(mnemonic, coin)
	if err != nil {
		return "", fmt.Errorf("key: %w", err)
	}
	var address string
	err = w.load(coin, func(walletter Walletter) error {
		address, err = walletter.GetAddressAt(coin, addressKey)
		return err
	})
	return address, err
}

func (w *Wallet) Transfer(mnemonic string, coin types.Coin, amount, fee, addr string, isFull bool) (string, error) {
	addressKey, err := addressKey(mnemonic, coin)
	if err != nil {
		return "", fmt.Errorf("key: %w", err)
	}
	var hash string
	err = w.load(coin, func(walletter Walletter) error {
		hash, err = walletter.Transfer(coin, addressKey, amount, fee, addr, isFull)
		return err
	})
	return hash, err
}

func (w *Wallet) Api() api.Api {
	return w.backend.Api()
}

func addressKey(mnemonic string, coin types.Coin) (*hdkeychain.ExtendedKey, error) {
	seed := bip39.NewSeed(mnemonic, "")
	masterKey, err := hdkeychain.NewMaster(seed, btc.NetParams)
	if err != nil {
		return nil, fmt.Errorf("master: %w", err)
	}
	purposeKey, err := masterKey.Child(44)
	if err != nil {
		return nil, fmt.Errorf("purpose: %w", err)
	}
	coinKey, err := purposeKey.Child(coin.CoinTypeIndex())
	if err != nil {
		return nil, fmt.Errorf("coin: %w", err)
	}
	accountKey, err := coinKey.Child(0)
	if err != nil {
		return nil, fmt.Errorf("account: %w", err)
	}
	changeKey, err := accountKey.Child(0)
	if err != nil {
		return nil, fmt.Errorf("change: %w", err)
	}
	addressKey, err := changeKey.Child(0)
	if err != nil {
		return nil, fmt.Errorf("address: %w", err)
	}
	return addressKey, nil
}
