package crypto

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOpen_seal(t *testing.T) {
	passphrase := []byte("passphrase_1")
	in := []byte("start foam first mule inherit calm catch cool enrich arrow voice color")
	out, err := Seal(passphrase, in)
	require.Nil(t, err)
	in2, err := Open(append(passphrase, '2'), out)
	require.EqualError(t, err, "failed to decrypt")
	in2, err = Open(passphrase, out)
	require.Nil(t, err)
	require.True(t, reflect.DeepEqual(in, in2))
}
