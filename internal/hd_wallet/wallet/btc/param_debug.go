// +build !release

package btc

import "github.com/btcsuite/btcd/chaincfg"

var (
	NetParams = &chaincfg.TestNet3Params
)
