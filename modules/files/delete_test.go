// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package files

import (
	"os"
	"io/ioutil"
	"fmt"
	"testing"
)

func TestDeleteFileWithoutError(t *testing.T) {
	var f *os.File
	var err error
	var text = []byte("Lorem Ipsum Dolor Amet")

	if f, err = ioutil.TempFile("", "copy-file"); err != nil {
		t.Fatal(err)
	}

	filename := f.Name()
	if err := ioutil.WriteFile(filename, text, 0644); err != nil {
		t.Fatal(fmt.Errorf("WriteFile %s: %v", filename, err))
	}
	f.Close()

	if err = Delete(filename); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(filename); err == nil {
		t.Fatal(fmt.Errorf("File %s should be removed", filename))
	}
}
