package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUpperString_JSON(t *testing.T) {
	s := NewUpperString("btc")
	require.True(t, s.String() == "BTC")
	b, err := json.Marshal(s)
	require.Nil(t, err)
	require.True(t, string(b) == `"BTC"`)
	var s2 UpperString
	err = json.Unmarshal(b, &s2)
	require.Nil(t, err)
	require.True(t, s2.String() == "BTC")
}
