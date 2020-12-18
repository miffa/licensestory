package licensebox

import (
	"encoding/json"
	"code.troila.io/licenselicenseplate"
)

// LVerifier
type LVerifier interface {
	Verify(string) (*licenseplate.TpaasLicenseMeta, error)
}

func NewLVerifier(f, p string) (LVerifier, error) {
	return licenseplate.NewLicenseVerifier(f, p)
}

// CLVerifier
type CLVerifier interface {
	Verify(string) (string, error)
}

func NewCLVerifier(id, pk string) (CLVerifier, error) {
	var err error
	v := new(cLvVerifier)
	v.v, err = licenseplate.NewLicenseVerifier(id, pk)
	if err != nil {
		return nil, err
	}
	return v, nil
}

type cLvVerifier struct {
	v *licenseplate.LicenseVerifier
}

func (v *cLvVerifier) Verify(tx string) (string, error) {
	c, err := v.Verify(tx)
	if err != nil {
		return "", err
	}
	cp, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(cp), nil
}
