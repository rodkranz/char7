package files

import (
	"io/ioutil"
)

// Write create a new file with content data
func Write(dst string, bs []byte) error {
	return ioutil.WriteFile(dst, bs, 0644)
}
