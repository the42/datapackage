package datapackage

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
)

// the canonical data package file name
const Filename = "datapackage.json"

func MD5SumforData(data []byte) string {
	b := md5.Sum(data)
	return fmt.Sprintf("%x", b)
}

// lazyGetData will read the contents of a file only if the input parameter
// data is nil, otherwise it will return data
func lazyGetData(path string, data *[]byte) (*[]byte, error) {
	if data != nil {
		return data, nil
	}
	rdata, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &rdata, nil
}
