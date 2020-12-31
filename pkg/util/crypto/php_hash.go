package crypto

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Make(str string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hashed), err
}

func Check(str string, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(str)) == nil
}

const (
	key = "asdeszqsc_v1.0"
)

func mKey() string {
	mHash := md5.New()
	mHash.Write([]byte(key))
	return hex.EncodeToString(mHash.Sum(nil))
}

func Encrypt(data string) string {
	hKey := mKey()
	x, char := 0, make([]byte, 0, len(data))
	lenKey := len(hKey)
	for i := 0; i < len(data); i++ {
		if x == lenKey {
			x = 0
		}
		char = append(char, hKey[x])
		x++
	}
	var finishData []byte
	for i, v := range data {
		finishData = append(finishData, byte(v+int32(char[i])%256))
	}
	return base64.StdEncoding.EncodeToString(finishData)
}

func Decrypt(data string) (string, error) {
	hKey := mKey()
	bData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", fmt.Errorf("decode: %w", err)
	}
	x, char := 0, make([]byte, 0, len(bData))
	lenKey := len(hKey)
	for range bData {
		if x == lenKey {
			x = 0
		}
		char = append(char, hKey[x])
		x++
	}
	var finishData []byte
	for k, v := range bData {
		if v < char[k] {
			finishData = append(finishData, byte(int32(v)+256-int32(char[k])))
		} else {
			finishData = append(finishData, v-char[k])
		}
	}
	return string(finishData), nil
}
