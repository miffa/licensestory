package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/pkg/errors"
)

const (
	k = 16
)

func AesEncrypt(plaintext, key []byte) ([]byte, error) {
	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key[:k])
	if err != nil {
		return nil, errors.WithMessage(err, "aes.NewCipher")
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.WithMessage(err, "cipher.NewGCM")
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, errors.WithMessage(err, "rand.Reader")

	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	return aesGCM.Seal(nonce, nonce, plaintext, nil), nil
}

func AesDecrypt(enc, key []byte) ([]byte, error) {

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key[:k])
	if err != nil {
		return nil, errors.WithMessage(err, "aes.NewCipher")
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.WithMessage(err, "cipher.NewGCM")
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "GCM.Open")
	}
	return plaintext, nil
}
