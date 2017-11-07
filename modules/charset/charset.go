package charset

import (
	"bufio"
	"bytes"
	"os"

	"github.com/rodkranz/char7/modules/files"
	"github.com/rodkranz/char7/modules/mapping"
)

// HasChange the global variable that indicate has changes
var HasChange bool

// CharSet function that search things that needs to replace
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
