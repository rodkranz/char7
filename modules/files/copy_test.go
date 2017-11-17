package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	f, err := ioutil.TempFile(os.TempDir(), "origin-file-")
	if err != nil {
		t.Fatalf("expected none error, got %v", err.Error())
	}
	defer f.Close()

	mockTest := "Lorem Ipsum dolor set Amet"
	fmt.Fprintf(f, mockTest)
	f.Sync()

	dstFile := f.Name() + "copied"

	if err := Copy(f.Name(), dstFile); err != nil {
		t.Fatalf("expected none error, got %v", err.Error())
	}

	if _, err := os.Open(dstFile); os.IsNotExist(err) {
		t.Errorf("expected none error, got %v", err)
	}
}
