package files

import "os"

func Delete(src string) error {
	return os.Remove(src)
}
