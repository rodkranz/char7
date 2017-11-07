package files

import "os"

// Delete remove file from HD
func Delete(src string) error {
	return os.Remove(src)
}
