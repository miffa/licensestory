package cmd

import (
	"fmt"

	license "code.troila.io/license/licenseplate"
)

const (
	PrivateKey = "_private.pem"
	PublicKey  = "_public.pem"
)

var (
	uuid        string
	licensename string
	prefix      string
	corporation string
)

func GenerateKey(p string) error {
	err := license.GenRsaKeyFile(p)
	if err != nil {
		fmt.Println("生成RSA密钥对出错:%s", err)
		return err
	}
	fmt.Println("生成RSA密钥对成功")
	return nil
}
