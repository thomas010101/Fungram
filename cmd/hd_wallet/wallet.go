package hd_wallet

import (
	"encoding/json"
	mErrors "wallet/errors"
	"wallet/internal/hd_wallet"
	"wallet/types"
)

type response struct {
	ErrNo int64       `json:"errNo"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

type generateRes struct {
	Data string `json:"data"`
}

type Wallet struct {
	w        *hd_wallet.HdWallet
	language types.Language
}

func GenerateWallet(domain, path string, language string) *Wallet {
	return &Wallet{
		w:        hd_wallet.NewWithFile(domain, path),
		language: types.ToLanguage(language),
	}
}

func (w *Wallet) Generate(pwd string) string {
	return w.response(w.generate(pwd))
}

func (w *Wallet) generate(pwd string) (generateRes, error) {
	out, err := w.w.Generate(pwd)
	if err != nil {
		return generateRes{}, err
	}
	return generateRes{
		Data: out,
	}, nil
}

func (w *Wallet) importMnemonic(pwd, mnemonic string) (generateRes, error) {
	out, err := w.w.ImportMnemonic(pwd, mnemonic)
	if err != nil {
		return generateRes{}, err
	}
	return generateRes{
		Data: out,
	}, nil
}

func (w *Wallet) ImportMnemonic(pwd, mnemonic string) string {
	return w.response(w.importMnemonic(pwd, mnemonic))
}

func (w *Wallet) ExportMnemonic() string {
	return w.response(w.w.ExportMnemonic())
}

func (w *Wallet) RandomMnemonic() string {
	return w.response(w.w.RandomMnemonic())
}

func (w *Wallet) LoadMnemonic(pwd, data string) string {
	err := w.w.LoadMnemonic(pwd, data)
	return w.response(nil, err)
}

func (w *Wallet) GetAddress(coin string) string {
	return w.response(w.w.GetAddress(coin))
}

func (w *Wallet) Transfer(coin, amount, fee, addr string, isFull bool) string {
	return w.response(w.w.Transfer(coin, amount, fee, addr, isFull))
}

func (w *Wallet) response(data interface{}, err error) string {
	mErr := mErrors.As(err)
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	response := &response{
		ErrNo: mErr.Code(),
		//Msg:   mErr.Msg(w.language),
		Msg:  msg,
		Data: mErr.Data(data),
	}
	buf, err := json.Marshal(response)
	if err != nil {
		panic("unknown error")
	}
	return string(buf)
}
