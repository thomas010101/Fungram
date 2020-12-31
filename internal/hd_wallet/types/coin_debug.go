// +build !release

package types

import "fmt"

const (
	USDTContractAddress = "0x553f41f3cc8633a5239762109672a4cb4822b34a"
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
		return 18
	}
	panic(fmt.Sprintf("coin(%d) decimals not found", c))
}
