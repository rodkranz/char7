package charset

import (
	"os"
	"bufio"
	"fmt"
	"encoding/json"
	"errors"
	"strings"
	"bytes"
	"github.com/rodkranz/char7/modules/settings"
)

type c7 struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type c7Map struct {
	Map []c7
}

func (cm *c7Map) Replace(line string) (string) {
	var newLine string

	for _, c := range cm.Map {
		if strings.Contains(line, c.Key) {
			newLine = strings.Replace(line, c.Key, c.Value, -1)
		}
	}

	return newLine
}

func readCharSetJson(src string) (c7Map, error) {
	charsetMap, err := os.Open(src)
	if err != nil {
		return c7Map{}, errors.New(fmt.Sprintf("Opening %s file: %s", src, err.Error()))
	}

	mapC7 := make([]c7, 0)
	jsonParser := json.NewDecoder(charsetMap)
	if err = jsonParser.Decode(&mapC7); err != nil {
		return c7Map{}, errors.New(fmt.Sprintf("Parsing %s file: %s", src, err.Error()))
	}

	return c7Map{Map: mapC7}, nil
}

func CharSet(src string) (oldString, newString bytes.Buffer, err error) {

	f, err := os.Open(src)
	if err != nil {
		return oldString, newString, err
	}
	defer f.Close()

	c7map, err := readCharSetJson(settings.MapCharset)
	if err != nil {
		return oldString, newString, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		oldString.WriteString(line + "\n")
		newString.WriteString(c7map.Replace(line))
	}

	if err := scanner.Err(); err != nil {
		return oldString, newString, err
	}

	return oldString, newString, nil
}
