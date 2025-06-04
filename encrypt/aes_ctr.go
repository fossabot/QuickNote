package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

func AESCTREncrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize) // 16
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(data))
	cipher.NewCTR(block, iv).XORKeyStream(ciphertext, data)

	// contains iv + encrypted data
	out := make([]byte, 0, len(iv)+len(ciphertext))
	out = append(out, iv...)
	out = append(out, ciphertext...)

	return out, nil
}

func AESCTRDecrypt(encrypted, key []byte) ([]byte, error) {
	if len(encrypted) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := encrypted[aes.BlockSize:]

	plaintext := make([]byte, len(ciphertext))
	// 2: iv
	stream := cipher.NewCTR(block, encrypted[:aes.BlockSize])
	stream.XORKeyStream(plaintext, ciphertext)
	return plaintext, nil
}
