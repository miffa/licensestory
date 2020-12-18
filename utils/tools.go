package utils

import (
	"encoding/binary"
	"io/ioutil"
	"os"
	"time"
)

const (
	EnvUuid                 = "TPAAS_UUID"
	WEB_FORMAT              = "20060102150405"
	WEB_STD_FORMAT          = "2006-01-02 15:04:05"
	WEB_1970_FORMAT         = "2006-01-02T15:04:05Z"
	WEB_1970_FORMAT_DEFAULT = "0000-01-01T00:00:00Z"
	DB_STD_FORMAT           = "2006-01-02 15:04:05"
)

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf

}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))

}

func LoadFile(filename string) ([]byte, error) {
	idata, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return idata, nil
}

func GetUuidFromEnv() string {
	return os.Getenv(EnvUuid)
}

func Timestamp() string {
	return time.Now().Format(WEB_FORMAT)
}

func Str2Time(str string) (time.Time, error) {
	return time.Parse(WEB_FORMAT, str)
}

func StdStr2Time(str string) (time.Time, error) {
	return time.Parse(WEB_STD_FORMAT, str)
}
