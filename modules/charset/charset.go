// Package charset is for transform character from files.
package charset

import (
	"os"
	"bufio"
	"bytes"

	"github.com/rodkranz/char7/modules/files"
	"github.com/rodkranz/char7/modules/mapping"
)

// HasChange returns boolean if latest file has character to change
// and did a transformation.
var HasChange bool

// CharSet convert character para o code mapped in .c7map file
// also, create a backup file for security if you need to recovery
// Usage:
//    CharSet("tmp/index.html") would return nil error.
func CharSet(src string) (err error) {
	HasChange = false

	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	c7map, err := mapping.GetMapping()
	if err != nil {
		return err
	}

	var buff bytes.Buffer
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		has, wrapper := c7map.Replace(line)

		if !HasChange && has {
			HasChange = true
		}
		buff.WriteString(wrapper)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return files.Write(src, buff.Bytes())
}
