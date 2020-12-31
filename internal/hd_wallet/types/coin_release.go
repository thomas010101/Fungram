// +build release

package types

import "fmt"

const (
	USDTContractAddress = "0xdac17f958d2ee523a2206206994597c13d831ec7"
)

func (c Coin) Decimals() int {
	switch c {
	case BTCCoin:
		return 8
	case OmniUSDTCoin:
		return 8
	case ETHCoin:
		return 18
	case Erc20USDTCoin:
		return 6
	}
	panic(fmt.Sprintf("coin(%d) decimals not found", c))
}
