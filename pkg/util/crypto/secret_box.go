package crypto

import (
	"crypto/rand"
	"fmt"
	"io"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/nacl/secretbox"
)

const (
	saltLength  = 16
	nonceLength = 24
	keyLength   = 32
)

func deriveKey(passphrase []byte, salt []byte) [keyLength]byte {
	// argon2被用来从一个密码来派生加密密钥
	secretKeyBytes := argon2.IDKey(passphrase, salt, 4, 64*1024, 4, 32)
	var secretKey [keyLength]byte
	copy(secretKey[:], secretKeyBytes)
	return secretKey
}

func Seal(passphrase, in []byte) ([]byte, error) {
	var nonce [nonceLength]byte
	// 生成一个随机nonce
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		return nil, err
	}
	salt := make([]byte, saltLength)
	// 生成一个随机salt
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	// 根据密码、盐，派生一个密钥
	secretKey := deriveKey(passphrase, salt)
	// 前缀为：salt + nonce
	prefix := append(salt, nonce[:]...)
	// 密文：加密的身份验证的in的拷贝被追加打prefix尾部，并返回。
	cipherText := secretbox.Seal(prefix, []byte(in), &nonce, &secretKey)
	return cipherText, nil
}

func Open(passphrase, out []byte) ([]byte, error) {
	if len(out) < nonceLength+saltLength {
		return nil, fmt.Errorf("data wrong")
	}
	salt := make([]byte, saltLength)
	// 从密文中取出salt
	copy(salt, out[:saltLength])
	var nonce [nonceLength]byte
	// 从密文中取出nonce
	copy(nonce[:], out[saltLength:nonceLength+saltLength])
	// 根据密码、盐，派生一个密钥
	secretKey := deriveKey(passphrase, salt)
	// 解密剩余部分
	decrypted, ok := secretbox.Open(nil, out[nonceLength+saltLength:], &nonce, &secretKey)
	if !ok {
		return nil, fmt.Errorf("failed to decrypt")
	}
	return decrypted, nil
}
