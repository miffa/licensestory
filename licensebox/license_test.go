package licensebox

import (
	"code.troila.io/licenseutils"
	"os"
	"testing"
)

var (
	PublicKeyFile = "./licenseplate/rsa_public_key.pem"

	testuuid = "zse45f48ybjy69h,j06lho9ioi28udjr"
)

func TestLicenseVerify(t *testing.T) {
	os.Setenv(utils.EnvUuid, testuuid)

	mylicense := `QnpRdGcrd3AxT2luaWdvWGpTK0ZUeGFnV1pmTU9SWGVzNkFNNXp3RFNXY2wwUEE3Vk1iLzBFTUkwaVpYTk4ydTdycTFVL09ZTWJ3TjVpbjZnM2dhTzFOSEpyZUdiRjhmc2t1UHkrZkdRQ0lHZXFDZG5jWFdwTlowL3E0cmhCUGRxa1FNdkJNM1pTT0srMTNuUW9OWS9xUHU5aG1sYkRFSkl2VUZBT2ExZ1U4TXpWSlJwVHhGckpxZmZ0Z1ZGeGd1cGxnVlJqbXRhMHl4ekNlcSBYSmtwL0pDU3c5T3NlN0I4MzA0ZitVaDNlV0ZhaS9YMzlPMS93K2VFblFLVlZ5T09ncGdsYWt3N24zblh5SmhMUG4wOUl4RXUvbDBDYmUyYUdTRjVxb2JnU1ZGelZQYUlkandqdmVwdlcwUnlkZnpGd0FGOTFpWCt0ZFlsaVFYWTZiRHd6RytwdnprZEJZTnhmSUt1UmNUM0hzY0ZBa05TdzZ2S3BZajVGcWs9`

	vv, err := NewLVerifier(utils.GetUuidFromEnv(), PublicKeyFile)
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
