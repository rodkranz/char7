package files

import (
	"bitbucket.org/rkranz/tmp/modules/settings"
	"path/filepath"
	"os"
)


var filePaths []string

func SearchFiles() ([]string, error) {
	filePaths = make([]string, 0)
	return filePaths, filepath.Walk(settings.Folder, walk)

}

func walk(path string, info os.FileInfo, _ error) error {
	if info.IsDir()  {
		return nil
	}

	if !allowExtToContinue(filepath.Ext(info.Name()), settings.ExtFile) {
		return nil
	}

	filePaths = append(filePaths, path)
	return nil
}

func allowExtToContinue(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
