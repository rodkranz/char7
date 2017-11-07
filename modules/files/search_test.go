package files

import (
	"testing"
	"os"
	"io/ioutil"
	"path/filepath"
)

type TestSetup struct {
	ListFiles []*os.File
	Dir       string
}

func Setup(t *testing.T, prefix []string, n int) *TestSetup {
	dir := os.TempDir()
	list := make([]*os.File, 0, n*len(prefix))

	// make temp files
	for _, pref := range prefix {
		f, err := ioutil.TempFile(dir, pref)
		if err != nil {
			t.Fatalf("Error to create temp file: %v", err)
		}
		list = append(list, f)
	}

	return &TestSetup{
		Dir:       dir,
		ListFiles: list,
	}
}

func TestWalkOk(t *testing.T) {
	// make temp files
	testExpected := Setup(t, []string{"aa.", "bb."}, 5)

	// get files
	exts := make([]string, 0)
	for _, f := range testExpected.ListFiles {
		exts = append(exts, filepath.Ext(f.Name()))
	}

	// define files as we have for allow in extension
	optFilter = &Filter{
		Exts: exts,
	}

	err := filepath.Walk(testExpected.Dir, walk)
	if err != nil {
		t.Fatalf("expected none error, got %v", err.Error())
	}

	if len(testExpected.ListFiles) != len(files) {
		t.Fatalf("Expected %d files, got %d files", len(testExpected.ListFiles), len(files))
	}

	for _, fileExpected := range testExpected.ListFiles {
		found := false
		for _, fileActual := range files {
			if fileActual == fileExpected.Name() {
				found = true
			}
		}

		if !found {
			t.Errorf("Expected found file %v, got none", fileExpected.Name())
		}
	}
}

//func helpMd5File(src string) (string, error) {
//	f, err := os.Open(src)
//	if err != nil {
//		return "", err
//	}
//	defer f.Close()
//
//	h := md5.New()
//	if _, err := io.Copy(h, f); err != nil {
//		return "", err
//	}
//
//	return fmt.Sprintf("%x", h.Sum(nil)), nil
//}

func TestAllowExtToContinue(t *testing.T) {

	tests := []struct {
		param1   string
		param2   []string
		Expected bool
	}{
		{
			param1: "/var/tmp/lorem3",
			param2: []string{
				"/var/tmp/lorem1",
				"/var/tmp/lorem2",
				"/var/tmp/lorem3",
			},
			Expected: true,
		},
		{
			param1: "/var/tmp/lorem",
			param2: []string{
				"/var/tmp/lorem1",
				"/var/tmp/lorem2",
				"/var/tmp/lorem3",
			},
			Expected: false,
		},
	}

	for _, test := range tests {
		actual := allowExtToContinue(test.param1, test.param2)
		if actual != test.Expected {
			t.Errorf("Expected %t, but got %v", test.Expected, actual)
		}
	}
}
