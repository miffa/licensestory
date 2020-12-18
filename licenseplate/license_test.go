package licenseplate

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"code.miffa.io/licenseutils"
	"os"
	"testing"

	"github.com/pkg/errors"
)

var (
	PrivateKeyFile = "rsa_private_key.pem"
	PublicKeyFile  = "rsa_public_key.pem"

	testuuid = "zse45f48ybjy69h,j06lho9ioi28udjr"
)

func loadPem(filename string) ([]byte, error) {
	idata, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return idata, nil
}

func initRsa() ([]byte, []byte, error) {
	prikey, err := loadPem(PrivateKeyFile)
	if err != nil {
		return nil, nil, errors.WithMessage(err, "init private key")
	}
	pubkey, err := loadPem(PublicKeyFile)
	if err != nil {
		return nil, nil, errors.WithMessage(err, "init public key")
	}

	return pubkey, prikey, nil
}

func TestLicense(t *testing.T) {

	t.Logf("1. init rsa key from file")
	pu, pi, err := initRsa()
	if err != nil {
		t.Errorf("load pem file err:%v\n", err)
		return
	}

	t.Logf("2: magic no:%s\n", string(magicNumber))
	t.Logf("2: magic no:%s\n", getMagicNumber())

	t.Logf("3. new license factory\n")
	Nike := NewLicenseFactory(pu, pi)

	mayday := TpaasLicenseMeta{
		Corporation: "fucar.com.ltd",
		Quota:       40,
		ExpiredTime: "2022-02-03 23:59:59",
		Extension:   "thi is a test",
		Version:     "v2.34",
	}

	maydayuuid := TpaasLicenseUuid{
		Uuid: testuuid,
		//MagicNo: string(MagicNumber),
		MagicNo: getMagicNumber(),
	}

	itd, _ := json.Marshal(mayday)

	t.Logf("4. generate license for %s\n", string(itd))
	License, err := Nike.Release(
		maydayuuid,
		mayday,
	)
	if err != nil {
		t.Errorf("sign license err:%v\n", err)
		return
	}
	t.Logf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n")
	t.Logf("my license: %s\n", License)
	t.Logf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n")

	t.Logf("5. verify licenes\n")

	maydayuuidcopy := TpaasLicenseUuid{
		Uuid: testuuid,
		//MagicNo: getMagicNumber(),
		MagicNo: "fdgfdgdfgfdgd",
	}
	id, err := Nike.Verify(maydayuuidcopy, License)
	if err != nil {
		t.Logf("sign license err:%v expected\n", err)
	}

	maydayuuidcopy1 := TpaasLicenseUuid{
		Uuid:    testuuid,
		MagicNo: getMagicNumber(),
	}
	id, err = Nike.Verify(maydayuuidcopy1, License)
	if err != nil {
		t.Logf("verify license err:%v \n", err)
		return
	}
	t.Logf("6. verify licenes ok :%v\n", *id)
}

func TestLicenseMgr(t *testing.T) {

	t.Logf("1. ======================================================")
	//ss, err := NewLicenseSigner(testuuid, PublicKeyFile, PrivateKeyFile)
	ss, err := NewLicenseSigner(testuuid, "mytest_public.pem", "mytest_private.pem")
	if err != nil {
		t.Errorf("init signer err:%v", err)
		return
	}

	mayday := TpaasLicenseMeta{
		Corporation: "badluckin.com.ltd",
		Quota:       40,
		ExpiredTime: "2022-02-03 23:59:59",
		Extension:   "thi is a test",
		Version:     "v2.34",
	}
	LicenseText, err := ss.Sign(mayday)
	if err != nil {
		t.Errorf("init signer err:%v", err)
		return
	}
	t.Logf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n")
	t.Logf("my license: %s\n", LicenseText)
	t.Logf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n")

	id, err := ss.Verify(LicenseText)
	if err != nil {
		t.Logf("verify license err:%v \n", err)
		return
	}
	ep, _ := json.Marshal(id)
	t.Logf("6. verify licenes ok :%#v\n", string(ep))
	t.Logf("7. ======================================================")
}

func TestLicenseVerify(t *testing.T) {
	os.Setenv(utils.EnvUuid, testuuid)

	mylicense := `QnpRdGcrd3AxT2luaWdvWGpTK0ZUeGFnV1pmTU9SWGVzNkFNNXp3RFNXY2wwUEE3Vk1iLzBFTUkwaVpYTk4ydTdycTFVL09ZTWJ3TjVpbjZnM2dhTzFOSEpyZUdiRjhmc2t1UHkrZkdRQ0lHZXFDZG5jWFdwTlowL3E0cmhCUGRxa1FNdkJNM1pTT0srMTNuUW9OWS9xUHU5aG1sYkRFSkl2VUZBT2ExZ1U4TXpWSlJwVHhGckpxZmZ0Z1ZGeGd1cGxnVlJqbXRhMHl4ekNlcSBYSmtwL0pDU3c5T3NlN0I4MzA0ZitVaDNlV0ZhaS9YMzlPMS93K2VFblFLVlZ5T09ncGdsYWt3N24zblh5SmhMUG4wOUl4RXUvbDBDYmUyYUdTRjVxb2JnU1ZGelZQYUlkandqdmVwdlcwUnlkZnpGd0FGOTFpWCt0ZFlsaVFYWTZiRHd6RytwdnprZEJZTnhmSUt1UmNUM0hzY0ZBa05TdzZ2S3BZajVGcWs9`

	vv, err := NewLicenseVerifier(utils.GetUuidFromEnv(), PublicKeyFile)
	if err != nil {
		t.Errorf("init verifier err:%v", err)
		return
	}

	id, err := vv.Verify(mylicense)
	if err != nil {
		t.Logf("verify license err:%v \n", err)
		return
	}

	t.Logf("6. the verifier verify licenes ok :%#v\n", *id)
}

func TestLicenseGenKey(t *testing.T) {
	pu, pk, err := GenRsaKey()
	if err != nil {
		t.Errorf("init rsa key err:%v", err)
		return
	}

	t.Logf("public: %s\n", hex.EncodeToString(pu))
	t.Logf("private: %s\n", hex.EncodeToString(pk))
	t.Logf("public: %s\n", base64.StdEncoding.EncodeToString(pu))
	t.Logf("private: %s\n", base64.StdEncoding.EncodeToString(pk))

	//GenRsaKeyFile("mytest")

}
