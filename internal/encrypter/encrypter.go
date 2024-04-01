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

func (em *EncryptionManager) EncryptByte(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(em.createKey(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func (em *EncryptionManager) DecryptByte(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(em.createKey(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func (em *EncryptionManager) EncryptString(text, key string) (string, error) {
	byteEncrypted, err := em.EncryptByte([]byte(text), []byte(key))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(byteEncrypted), nil
}

func (em *EncryptionManager) DecryptString(text, key string) (string, error) {
	b, err := hex.DecodeString(text)
	if err != nil {
		return "", err
	}
	res, err := em.DecryptByte(b, []byte(key))
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func (em *EncryptionManager) createKey(key []byte) []byte {
	em.hasher.Write(key)
	b := em.hasher.Sum(nil)
	em.hasher.Reset()
	return b
}
