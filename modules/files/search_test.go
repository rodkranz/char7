package files

import (
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"testing"
	"time"
)

type TestSetup struct {
	ListFiles []*os.File
	Dir       string
}

func Setup(t *testing.T, prefix map[string]string, n int) *TestSetup {
	files = []string{}
	dir := os.TempDir()
	list := make([]*os.File, 0, n*len(prefix))

	// make temp files
	for pref, suffix := range prefix {
		f, err := TempFile(dir, pref, suffix)
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

func TestSearchFiles(t *testing.T) {
	testExpected := Setup(t, map[string]string{"aa.": ".find", "bb.": ".ignore"}, 5)

	filter := Filter{
		Dir:  testExpected.Dir,
		Exts: []string{".find"},
	}

	files, err := SearchFiles(&filter)
	if err != nil {
		t.Fatalf("expected none error, got %v", err)
	}

	for _, fileExpected := range testExpected.ListFiles {
		found := false

		if !allowExtToContinue(filepath.Ext(fileExpected.Name()), filter.Exts) {
			continue
		}

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

func TestWalkOk(t *testing.T) {
	// make temp files
	testExpected := Setup(t, map[string]string{"aa.": "", "bb.": ""}, 5)

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
		t.Fatalf("expected none error, got %v", err)
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

	t.Run("ignoreFolder", func(t *testing.T) {
		files = []string{}

		f, err := os.OpenFile(testExpected.Dir, os.O_RDONLY, 0755)
		if err != nil {
			t.Fatalf("expected none error, got %v", err)
		}

		s, err := f.Stat()
		if err != nil {
			t.Fatalf("expected none error, got %v", err)
		}

		err = walk(testExpected.Dir, s, nil)
		if err != nil {
			t.Fatalf("expected none error, got %v", err)
		}

		if len(files) != 0 {
			t.Fatalf("expected none files, got %v", len(files))
		}
	})

	t.Run("ignoreExtension", func(t *testing.T) {
		files = []string{}
		f := testExpected.ListFiles[0]

		optFilter = &Filter{
			FileName: "test",
			Exts:     []string{".tmp"},
		}

		s, err := f.Stat()
		if err != nil {
			t.Fatalf("expected none error, got %v", err)
		}

		err = walk(testExpected.Dir, s, nil)
		if err != nil {
			t.Fatalf("expected none error, got %v", err)
		}

		if len(files) != 0 {
			t.Fatalf("expected none files, got %v", len(files))
		}
	})

}

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

func TempFile(dir, prefix, suff string) (f *os.File, err error) {
	if dir == "" {
		dir = os.TempDir()
	}

	nconflict := 0
	for i := 0; i < 10000; i++ {
		name := filepath.Join(dir, prefix+nextSuffix()+suff)
		f, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
		if os.IsExist(err) {
			if nconflict++; nconflict > 10 {
				randmu.Lock()
				rand = reseed()
				randmu.Unlock()
			}
			continue
		}
		break
	}
	return
}

var rand uint32
var randmu sync.Mutex

func reseed() uint32 {
	return uint32(time.Now().UnixNano() + int64(os.Getpid()))
}

func nextSuffix() string {
	randmu.Lock()
	r := rand
	if r == 0 {
		r = reseed()
	}
	r = r*1664525 + 1013904223 // constants from Numerical Recipes
	rand = r
	randmu.Unlock()
	return strconv.Itoa(int(1e9 + r%1e9))[1:]
}
