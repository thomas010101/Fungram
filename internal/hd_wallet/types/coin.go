package types

import (
	"fmt"
	"strings"
)

type Coin uint8

const (
	_ Coin = iota
	BTCCoin
	ETHCoin
	OmniUSDTCoin
	Erc20USDTCoin
)

var (
	coins = [...]Coin{
		BTCCoin,
		ETHCoin,
		OmniUSDTCoin,
		Erc20USDTCoin,
	}
)

func ToCoin(s string) (Coin, error) {
	s = strings.ToUpper(s)
	for _, v := range coins {
		if v.String() == s {
			return v, nil
		}
	}
	return 0, fmt.Errorf("coin(%s) not found", s)
}

func (c Coin) Child() Coin {
	switch c {
	case BTCCoin, OmniUSDTCoin:
		return BTCCoin
	case ETHCoin, Erc20USDTCoin:
		return ETHCoin
	}
	panic(fmt.Sprintf("coin(%d)", c))
}

func (c Coin) CoinTypeIndex() uint32 {
	switch c.Child() {
	case BTCCoin:
		return 0
	case ETHCoin:
		return 60
	}
	panic(fmt.Sprintf("coin(%d) coin type index", c))
}

func (c Coin) String() string {
	switch c {
	case BTCCoin:
		return "BTC"
	case OmniUSDTCoin:
		return "OMNI-USDT"
	case ETHCoin:
		return "ETH"
	case Erc20USDTCoin:
		return "ERC20-USDT"
	}
	return fmt.Sprintf("<coin(%d) not found>", c)
}
