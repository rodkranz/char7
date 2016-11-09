// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package files

import (
	"os"
	"fmt"
	"testing"
	"io/ioutil"
	"bytes"
)

func TestCopyShouldCreateNewFileWithoutError(t *testing.T) {
	text := []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit. " +
		"Suspendisse eu pulvinar ipsum. Vestibulum elit erat, blandit non lacus id, " +
		"eleifend finibus purus. Interdum et malesuada fames ac ante ipsum primis in faucibus. " +
		"Interdum et malesuada fames ac ante ipsum primis in faucibus. " +
		"Praesent sit amet mi ac mi lobortis venenatis. Ut sagittis dui sed iaculis convallis. " +
		"Vestibulum interdum tellus in dolor ullamcorper cursus. " +
		"Sed placerat sapien eget rutrum malesuada. " +
		"In lacinia pellentesque eros ut venenatis.")

	var f *os.File
	var err error

	if f, err = ioutil.TempFile("", "copy-file"); err != nil {
		t.Fatal(err)
	}

	filename := f.Name()
	if err := ioutil.WriteFile(filename, text, 0644); err != nil {
		t.Fatal(fmt.Errorf("WriteFile %s: %v", filename, err))
	}
	f.Close()

	output := "/tmp/_test_"
	if err = Copy(f.Name(), output); err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(text, b) {
		t.Fatal("The copy has no the same text from origin")
	}
}