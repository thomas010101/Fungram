package hd_wallet

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"wallet/internal/hd_wallet/external/api"
	"wallet/internal/hd_wallet/types"
	"wallet/internal/hd_wallet/wallet"
	"wallet/pkg/util/crypto"

	"github.com/tyler-smith/go-bip39"
)

//go:generate mockgen -destination wallet_mock.go -source wallet.go -package hd_wallet
type database interface {
	ReadAll() ([]byte, error)
	io.Writer
}

type HdWallet struct {
	mnemonic string
	db       database
	wallet   *wallet.Wallet
	api      api.Api
}

func NewWithFile(domain string, path string) *HdWallet {
	if len(path) == 0 {
		return New(domain, emptyWrite{})
	}
	return New(domain, writeFile{path: path})
}

func New(domain string, db database) *HdWallet {
	hw := &HdWallet{
		db: db,
		api: api.New(domain, api.OptWithClient(&http.Client{
			Transport: &http.Transport{},
		})),
	}
	hw.wallet = wallet.New(hw, wallet.OptWithBtc(), wallet.OptWithEth())
	return hw
}

func (hw *HdWallet) Generate(pwd string) (string, error) {
	mnemonic, err := newMnemonic()
	if err != nil {
		return "", fmt.Errorf("new mnemonic: %w", err)
	}
	out, err := hw.ImportMnemonic(pwd, mnemonic)
	if err != nil {
		return "", fmt.Errorf("write: %w", err)
	}
	return out, nil
}

func (hw *HdWallet) ExportMnemonic() (string, error) {
	if err := hw.checkMnemonic(hw.mnemonic); err != nil {
		return "", err
	}
	return hw.mnemonic, nil
}

func (hw *HdWallet) ImportMnemonic(pwd, mnemonic string) (string, error) {
	if err := hw.checkMnemonic(mnemonic); err != nil {
		return "", err
	}
	passphrase, err := getKeyByAt(pwd)
	if err != nil {
		return "", fmt.Errorf("passphrase: %w", err)
	}
	out, err := crypto.Seal(passphrase, []byte(mnemonic))
	if err != nil {
		return "", fmt.Errorf("encrypt: %w", err)
	}
	if _, err = hw.db.Write(out); err != nil {
		return "", fmt.Errorf("fileWrite: %w", err)
	}
	hw.mnemonic = mnemonic
	return hex.EncodeToString(out), nil
}

func (hw *HdWallet) RandomMnemonic() ([]string, error) {
	if err := hw.checkMnemonic(hw.mnemonic); err != nil {
		return nil, err
	}
	list := strings.Split(hw.mnemonic, " ")
	rand.Seed(time.Now().UnixNano())
	for i := len(list) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		list[i], list[num] = list[num], list[i]
	}
	return list, nil
}

func (hw *HdWallet) LoadMnemonic(pwd, data string) error {
	passphrase, err := getKeyByAt(pwd)
	if err != nil {
		return fmt.Errorf("key: %w", err)
	}
	var outData []byte
	if len(data) == 0 {
		outData, err = hw.db.ReadAll()
		if err != nil {
			return fmt.Errorf("readAll: %w", err)
		}
	} else {
		outData, err = hex.DecodeString(data)
		if err != nil {
			return fmt.Errorf("decode data: %w", err)
		}
	}
	out, err := crypto.Open(passphrase, outData)
	if err != nil {
		return fmt.Errorf("decrypt: %w", err)
	}
	mnemonic := string(out)
	if err := hw.checkMnemonic(mnemonic); err != nil {
		return err
	}
	hw.mnemonic = mnemonic
	return nil
}

func (hw *HdWallet) checkMnemonic(mnemonic string) error {
	if len(mnemonic) == 0 {
		return fmt.Errorf("mnemonic cannot be empty")
	}
	if !bip39.IsMnemonicValid(mnemonic) {
		return fmt.Errorf("failed mnemonic")
	}
	return nil
}

func (hw *HdWallet) GetAddress(s string) (string, error) {
	coin, err := types.ToCoin(s)
	if err != nil {
		return "", fmt.Errorf("toCoin: %w", err)
	}
	if err := hw.checkMnemonic(hw.mnemonic); err != nil {
		return "", err
	}
	return hw.wallet.GetAddress(hw.mnemonic, coin)
}

func (hw *HdWallet) Transfer(sCoin, amount, fee, addr string, isFull bool) (string, error) {
	coin, err := types.ToCoin(sCoin)
	if err != nil {
		return "", fmt.Errorf("toCoin: %w", err)
	}
	if err := hw.checkMnemonic(hw.mnemonic); err != nil {
		return "", err
	}
	return hw.wallet.Transfer(hw.mnemonic, coin, amount, fee, addr, isFull)
}

func (hw *HdWallet) Api() api.Api {
	return hw.api
}

func getKeyByAt(pwd string) ([]byte, error) {
	if len(pwd) == 0 {
		return nil, fmt.Errorf("password cannot be empty")
	}
	h := hmac.New(sha512.New512_256, []byte("wallet"))
	h.Write([]byte(pwd))
	return h.Sum(nil), nil
}

func newMnemonic() (string, error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return "", fmt.Errorf("new entropy: %w", err)
	}
	return bip39.NewMnemonic(entropy)
}
