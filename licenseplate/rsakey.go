package licenseplate

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/pkg/errors"
)

func GenRsaKey() ([]byte, []byte, error) {
	size := 1024
	priv, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return nil, nil, errors.WithMessage(err, "Generate Rsa Key")
	}

	prk := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	publ, err := x509.MarshalPKIXPublicKey(priv.Public())
	if err != nil {
		return nil, nil, errors.WithMessage(err, "MarshalPKCS1PrivateKey Rsa Key")
	}
	puk := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: publ,
		},
	)
	return puk, prk, nil
}

func GenRsaKeyFile(prefix string) error {
	size := 1024
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return err

	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	file, err := os.Create(prefix + "_private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	publicKey := &privateKey.PublicKey

	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create(prefix + "_public.pem")
	if err != nil {
		return err
	}

	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}
