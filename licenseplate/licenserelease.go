package licenseplate

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strings"

	//"encoding/hex"
	"encoding/json"

	rsalib "code.miffa.io/license/crypt"

	"github.com/pkg/errors"
)

const (
	LicenseSeperator       = " "
	LicensePartQuantity    = 2
	LicensePartHeaderIndex = 0
	LicensePartBodyIndex   = 1

	EnvUuid = "TPAAS_UUID"
	PemKey  = "license_pem"
)

//  !!!! make a new magic number for our customers
var (
	//test magic number
	magicNumber []byte = []byte{0x00, 0x22, 0x58, 0x30, 0x15, 0x16, 0xff, 0xee}

	//magicNumber []byte = []byte{...}
)

func getMagicNumber() string {
	return hex.EncodeToString(magicNumber)
}

func NewLicenseFactory(pub, pri []byte) *LicenseFactory {
	return &LicenseFactory{Publickey: pub, PrivateKey: pri}
}

type LicenseFactory struct {
	Publickey  []byte
	PrivateKey []byte
}

func (l *LicenseFactory) Release(uuid TpaasLicenseUuid, meta TpaasLicenseMeta) (string, error) {
	// step 1.1 get uuidhash
	uuidhash := l.shaUuid(uuid)

	//step 1.2 package uuid_hash and meta and sign it
	//         generate license body
	m, err := json.Marshal(meta)
	if err != nil {
		return "", errors.WithMessage(err, "encrypt meta")
	}

	lbody, err := l.sign(l.pack(uuidhash, m))
	if err != nil {
		return "", errors.WithMessage(err, "sign license body")
	}
	licensebody := base64.StdEncoding.EncodeToString(lbody)

	//step 2.1 encrypt meta
	//         generate license header
	lheader, err := l.encrypt(m)
	if err != nil {
		return "", errors.WithMessage(err, "encrypt meta")
	}
	licenseheader := base64.StdEncoding.EncodeToString(lheader)

	//step 3  package encrypted meta and sign data
	return base64.StdEncoding.EncodeToString([]byte(licenseheader + LicenseSeperator + licensebody)), nil
}

func (l *LicenseFactory) Verify(uuid TpaasLicenseUuid, lice string) (*TpaasLicenseMeta, error) {
	//
	licensedata, err := base64.StdEncoding.DecodeString(lice)
	if err != nil {
		return nil, errors.WithMessage(err, "base64 decode license")
	}

	licenseitems := strings.Split(string(licensedata), LicenseSeperator)
	if len(licenseitems) != LicensePartQuantity {
		return nil, errors.New("invalid license")
	}

	licenseheader := licenseitems[LicensePartHeaderIndex]
	licensebody := licenseitems[LicensePartBodyIndex]

	lheader, err := base64.StdEncoding.DecodeString(licenseheader)
	if err != nil {
		return nil, errors.WithMessage(err, "base64 decode license header")
	}
	lbody, err := base64.StdEncoding.DecodeString(licensebody)
	if err != nil {
		return nil, errors.WithMessage(err, "base64 decode license body")
	}
	//
	m, err := l.decrypt(lheader)
	if err != nil {
		return nil, errors.WithMessage(err, "decrypt header")
	}

	// stepi 1.1 get uuidhash
	uuidhash := l.shaUuid(uuid)
	err = l.verify(l.pack(uuidhash, m), lbody)
	if err != nil {
		return nil, errors.WithMessage(err, "verify the signature")
	}

	var lc TpaasLicenseMeta
	err = json.Unmarshal(m, &lc)
	if err != nil {
		return nil, errors.WithMessage(err, "meta info not json")
	}
	return &lc, nil
}

func (l *LicenseFactory) shaUuid(uuid TpaasLicenseUuid) []byte {
	d := sha256.Sum256([]byte(uuid.Uuid + LicenseSeperator + uuid.MagicNo))
	return d[:]
	//return hex.EncodeToString(sha256.Sum256([]byte(uuid.Uuid + "|" + uuid.MagicNo))[:])
}

func (l *LicenseFactory) RsaEncrypt(d []byte) ([]byte, error) {
	return rsalib.RsaEncrypt(l.Publickey, d)
}

func (l *LicenseFactory) RsaDecrypt(d []byte) ([]byte, error) {
	return rsalib.RsaDecrypt(l.PrivateKey, d)
}

func (l *LicenseFactory) sign(d []byte) ([]byte, error) {
	return rsalib.RsaSign(l.PrivateKey, d)
}

func (l *LicenseFactory) verify(d, s []byte) error {
	return rsalib.RsaSignVer(l.Publickey, d, s)
}

func (l *LicenseFactory) encrypt(d []byte) ([]byte, error) {
	return rsalib.AesEncrypt(d, l.Publickey)
}

func (l *LicenseFactory) decrypt(d []byte) ([]byte, error) {
	return rsalib.AesDecrypt(d, l.Publickey)
}

func (l LicenseFactory) pack(d ...[]byte) []byte {
	buf := bytes.NewBuffer(nil)
	for _, v := range d {
		buf.Write(v)
	}
	return buf.Bytes()
}

func (l LicenseFactory) Marshal(d interface{}) (string, error) {
	m, err := json.Marshal(d)
	if err != nil {
		return "", errors.WithMessage(err, "json marshal")
	}
	return base64.StdEncoding.EncodeToString(m), nil
}

func (l LicenseFactory) HexEncode(d []byte) string {
	return base64.StdEncoding.EncodeToString(d)
}

func (l LicenseFactory) HexDecode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
