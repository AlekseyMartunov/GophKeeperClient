package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"hash"
	"io"
)

type EncryptionManager struct {
	hasher hash.Hash
}

func NewEncryptionManager() *EncryptionManager {
	hasher := md5.New()

	return &EncryptionManager{
		hasher: hasher,
	}
}

func (em *EncryptionManager) Encrypt(text, key string) (string, error) {
	block, err := aes.NewCipher(em.createKey(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", nil
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(text), nil)
	return hex.EncodeToString(ciphertext), nil
}

func (em *EncryptionManager) Decrypt(text, key string) (string, error) {
	block, err := aes.NewCipher(em.createKey(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", nil
	}

	data, err := hex.DecodeString(text)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil

}

func (em *EncryptionManager) createKey(key string) []byte {
	em.hasher.Write([]byte(key))
	b := em.hasher.Sum(nil)
	em.hasher.Reset()
	return b
}
