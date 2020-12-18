package licenseplate

// release
type TpaasLicenseUuid struct {
	Uuid    string
	MagicNo string
}

type TpaasLicenseMeta struct {
	Corporation string `json:"corporation"`
	Quota       int64  `json:"quota"`
	ExpiredTime string `json:"expired_time"`
	//Country string
	//Province string
	//City string
	Extension     string `json:"extension"`
	Version       string `json:"version"`
	ServicePeriod string `json:"service_period"`
	HomeLicense   string `json:"home_license"`
}

type TpaasLicenseInfo struct {
	TpaasLicenseMeta
	UuidHash string
}

type TpaasLicenseHeader []byte

type TpaasLicenseBody []byte

type TpaasLicense struct {
	Header   TpaasLicenseHeader
	SignBody TpaasLicenseBody
}

//verification
