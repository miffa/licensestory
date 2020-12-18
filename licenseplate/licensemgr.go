package licenseplate

import (
	"fmt"

	"code.troila.io/license/utils"

	"github.com/pkg/errors"
)

type LicenseVerifier struct {
	//
	uuid          string
	magicNumber   string
	licenseWorker *LicenseFactory
}

//
func NewLicenseVerifier(uuid, pemkeypath string) (*LicenseVerifier, error) {
	mgr := &LicenseVerifier{
		uuid: uuid,
		//uuid:        os.Getenv(EnvUuid),
		magicNumber: getMagicNumber(),
	}
	if mgr.uuid == "" {
		return nil, errors.New(fmt.Sprintf("License service init err : no ENV:%s in host", EnvUuid))
	}

	//pemkeypath := beego.AppConfig.String(PemKey)
	if pemkeypath == "" {
		return nil, errors.New("License service init err : public key in config")
	}

	pubkey, err := utils.LoadFile(pemkeypath)
	if err != nil {
		pubkey = []byte("1234567890123456")

		//return nil, errors.WithMessage(err, "read public key")
	}

	// the length  of the  key uesd for aes encryptipon must be more than 16
	if len(pubkey) < 16 {
		return nil, errors.New("License service init err : public key data invalid")
	}

	mgr.licenseWorker = NewLicenseFactory(pubkey, []byte("be tricked"))
	return mgr, nil
}

func (e *LicenseVerifier) Verify(licentext string) (*TpaasLicenseMeta, error) {
	uuid := TpaasLicenseUuid{Uuid: e.uuid, MagicNo: e.magicNumber}
	return e.licenseWorker.Verify(uuid, licentext)
}

func (e *LicenseVerifier) Close() error {
	return nil
}

////////
type LicenseSigner struct {
	//
	uuid          string
	magicNumber   string
	licenseWorker *LicenseFactory
}

func NewLicenseSigner(uid, puk, prk string) (*LicenseSigner, error) {

	if uid == "" {
		return nil, errors.New("uid is invalid")
	}

	mgr := &LicenseSigner{
		uuid:        uid,
		magicNumber: getMagicNumber(),
	}
	pubkey, err := utils.LoadFile(puk)
	if err != nil {
		return nil, errors.WithMessage(err, "read public key")
	}
	prikey, err := utils.LoadFile(prk)
	if err != nil {
		return nil, errors.WithMessage(err, "read private key")
	}

	// the length  of the  key uesd for aes encryptipon must be more than 16
	if len(pubkey) < 16 {
		return nil, errors.New("public key data invalid")
	}

	mgr.licenseWorker = NewLicenseFactory(pubkey, prikey)
	return mgr, nil
}

func (e *LicenseSigner) Sign(m TpaasLicenseMeta) (string, error) {
	uuid := TpaasLicenseUuid{Uuid: e.uuid, MagicNo: e.magicNumber}
	return e.licenseWorker.Release(uuid, m)
}

func (e *LicenseSigner) Verify(licentext string) (*TpaasLicenseMeta, error) {
	uuid := TpaasLicenseUuid{Uuid: e.uuid, MagicNo: e.magicNumber}
	return e.licenseWorker.Verify(uuid, licentext)
}

func (e *LicenseSigner) Close() error {
	return nil
}

////////
type LicenseDecryptor struct {
	licenseWorker *LicenseFactory
}

func NewLicenseDecryptor(prk string) (*LicenseDecryptor, error) {

	prikey, err := utils.LoadFile(prk)
	if err != nil {
		return nil, errors.WithMessage(err, "read private key")
	}

	mgr := &LicenseDecryptor{}

	mgr.licenseWorker = NewLicenseFactory(nil, prikey)
	return mgr, nil
}

func (e *LicenseDecryptor) Decrypt(s string) (string, error) {

	b, err := e.licenseWorker.HexDecode(s)
	if err != nil {
		return "", errors.WithMessage(err, "base64 decode")
	}
	bb, err := e.licenseWorker.RsaDecrypt(b)
	if err != nil {
		return "", errors.WithMessage(err, "rsa decrypt")
	}
	return string(bb), nil
}
