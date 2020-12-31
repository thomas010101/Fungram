package api

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	url2 "net/url"
	"time"
	"wallet/internal/hd_wallet/types"
	gTypes "wallet/types"

	"github.com/btcsuite/btcutil"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

//go:generate mockgen -destination api_mock.go -source api.go -package api
type Api interface {
	Nonce(address ethCommon.Address) (uint64, error)
	SendTx(coin types.Coin, data []byte) error
	Utxos(address btcutil.Address) ([]Coin, error)
}

type Option = func(a *api)

func OptWithClient(client *http.Client) Option {
	return func(a *api) {
		a.client = client
	}
}

type api struct {
	client *http.Client
	domain string
}

func New(domain string, opts ...Option) Api {
	a := &api{
		client: http.DefaultClient,
		domain: domain,
	}
	for _, v := range opts {
		v(a)
	}
	return a
}

func (a *api) Nonce(address ethCommon.Address) (uint64, error) {
	var res nonceRes
	if err := get(a.domain, a.client, nonceUrlKind, map[string]string{
		"address": address.Hex(),
	}, &res); err != nil {
		return 0, err
	}
	return res.Nonce.Uint64(), nil
}

func (a *api) SendTx(coin types.Coin, data []byte) error {
	return post(a.domain, a.client, sendTxUrlKind, sendTxReq{
		Symbol: gTypes.NewUpperString(coin.String()),
		Data:   hex.EncodeToString(data),
	}, nil)
}

func (a *api) Utxos(address btcutil.Address) ([]Coin, error) {
	var coins Coins
	if err := get(a.domain, a.client, utxoUrlKind, map[string]string{
		"address": address.String(),
	}, &coins); err != nil {
		return nil, err
	}
	return coins, nil
}

func do(client *http.Client, method, url string, body io.Reader, data interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	request, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}
	request.Header.Add("Imei", "hd_wallet")
	resp, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("do: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("resp statusCode(%d)", resp.StatusCode)
	}
	var res response
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return fmt.Errorf("decode body: %w", err)
	}
	if res.ErrNo != 0 {
		return fmt.Errorf("msg: %s", res.Msg)
	}
	if data == nil {
		return nil
	}
	if err = json.Unmarshal(res.Data, data); err != nil {
		return fmt.Errorf("decode data: %w", err)
	}
	return nil
}

func get(domain string, client *http.Client, kind urlKind, params map[string]string, data interface{}) error {
	u := url2.Values{}
	for k, v := range params {
		u.Set(k, v)
	}
	return do(client, http.MethodGet, fmt.Sprintf("%s?%s", kind.Url(domain), u.Encode()), nil, data)
}

func post(domain string, client *http.Client, kind urlKind, req interface{}, data interface{}) error {
	b, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	return do(client, http.MethodPost, kind.Url(domain), bytes.NewReader(b), data)
}
