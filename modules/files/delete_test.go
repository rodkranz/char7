package files

import (
	"testing"
	"os"
	"io/ioutil"
)

func TestDelete(t *testing.T) {
	f, err := ioutil.TempFile(os.TempDir(), "")
	if err != nil {
		t.Fatalf("expected none error, got %v", err.Error())
	}
	defer f.Close()

	if err := Delete(f.Name()); err != nil {
		t.Fatalf("expected none error, got %v", err.Error())
	}

	if _, err := os.Open(f.Name()); os.IsExist(err) {
		t.Errorf("expected none error, got %v", err)
	}
}
