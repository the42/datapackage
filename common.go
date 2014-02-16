package datapackage

import (
	"crypto/md5"
	"fmt"
)

const Filename = "datapackage.json"

func MD5SumforData(data []byte) string {
	b := md5.Sum(data)
	return fmt.Sprintf("%x", b)
}
