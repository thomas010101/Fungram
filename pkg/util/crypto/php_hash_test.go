package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"testing"

	"github.com/pborman/uuid"

	"github.com/stretchr/testify/require"
)

func TestMake_Check(t *testing.T) {
	h := md5.New()
	h.Write([]byte("aa888888"))
	data := hex.EncodeToString(h.Sum(nil))
	mData, err := Make(data)
	require.Nil(t, err)
	require.True(t, Check(data, mData))
}

func TestEncrypt_Decrypt(t *testing.T) {
	for i := 0; i < 1000; i++ {
		s := uuid.New()
		data := Encrypt(s)
		out, err := Decrypt(data)
		require.Nil(t, err)
		require.True(t, s == out)
	}
}
