package files

import (
	"os"
	"path/filepath"
	"strings"
)

var files []string

// Filter define options of search
type Filter struct {
	FileName string
	Exts     []string
	Dir      string
}

var optFilter *Filter

func init() {
	optFilter = new(Filter)
}

// SearchFiles return the list of files that were found by filter specification
func SearchFiles(filter *Filter) ([]string, error) {
	optFilter = filter

	files = make([]string, 0)
	return files, filepath.Walk(filter.Dir, walk)
}

func walk(path string, info os.FileInfo, _ error) error {
	if info.IsDir() {
		return nil
	}

	if len(optFilter.FileName) != 0 && !strings.Contains(info.Name(), optFilter.FileName) {
		return nil
	}

	if !allowExtToContinue(filepath.Ext(info.Name()), optFilter.Exts) {
		return nil
	}

	files = append(files, path)
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
