package files

import (
	"io/ioutil"
)

func Write(dst string, bs []byte) error {
	return ioutil.WriteFile(dst, bs, 0644)
}