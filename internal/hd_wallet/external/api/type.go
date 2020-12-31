package api

import (
	"encoding/json"
	"fmt"
	"wallet/pkg/util"
	"wallet/types"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcutil"
)

const (
	url = "http://192.168.0.230:9001"
)

type urlKind uint8

const (
	_ urlKind = iota
	utxoUrlKind
	sendTxUrlKind
	nonceUrlKind
)

func (u urlKind) Url(domain string) string {
	var path string
	switch u {
	case utxoUrlKind:
		path = "/api/v1/hdWallet/getBtcUnspent"
	case sendTxUrlKind:
		path = "/api/v1/hdWallet/sendTx"
	case nonceUrlKind:
		path = "/api/v1/hdWallet/txNonce"
	default:
		panic("url not found")
	}
	return fmt.Sprintf("%s%s", domain, path)
}

type Coin struct {
	Hash  *chainhash.Hash
	Index uint32
	Value btcutil.Amount
}

type jsonCoin struct {
	TxHash   string `json:"tx_hash"`
	OutputNo uint32 `json:"output_no"`
	Value    string `json:"value"`
}

func NewCoin(hash *chainhash.Hash, index uint32, value btcutil.Amount) Coin {
	return Coin{
		Hash:  hash,
		Index: index,
		Value: value,
	}
}

type response struct {
	ErrNo int64           `json:"errNo"`
	Msg   string          `json:"msg"`
	Data  json.RawMessage `json:"data"`
}

type Coins []Coin

func (c *Coins) UnmarshalJSON(b []byte) error {
	var coins []jsonCoin
	if err := json.Unmarshal(b, &coins); err != nil {
		return err
	}
	for _, v := range coins {
		hash, err := chainhash.NewHashFromStr(v.TxHash)
		if err != nil {
			return fmt.Errorf("hash(%s)", v.TxHash)
		}
		value, err := util.StringToInt64(v.Value, 8)
		if err != nil {
			return fmt.Errorf("value(%s)", v.Value)
		}
		*c = append(*c, Coin{
			Hash:  hash,
			Index: v.OutputNo,
			Value: btcutil.Amount(value),
		})
	}
	return nil
}

type UtxosRes []Coin

type sendTxReq struct {
	Symbol types.UpperString `json:"symbol"`
	Data   string            `json:"data"`
}

type nonceRes struct {
	Nonce types.Uint64 `json:"nonce"`
}
