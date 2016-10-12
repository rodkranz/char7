package charset

import (
	"os"
	"fmt"
	"bufio"
	"bytes"
	"errors"
	"strings"
	"encoding/json"

	"github.com/rodkranz/char7/modules/files"
	"github.com/rodkranz/char7/modules/settings"
	"github.com/rodkranz/char7/modules/chatdata"
)

type c7 struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type c7Map struct {
	Map []c7
}

var HasChange bool

func (cm *c7Map) Replace(line string) string {

	for _, c := range cm.Map {
		if strings.Contains(line, c.Key) {
			HasChange = true
			line = strings.Replace(line, c.Key, c.Value, -1)
		}
	}

	return fmt.Sprintf("%s\n", line)
}

func readCharSetJson(src string) (c7Map, error) {
	var jsonParser *json.Decoder
	charsetMap, err := os.Open(src)
	if err != nil {
		arrBytes := chatdata.MustAsset(".charset")
		bsRead := bytes.NewReader(arrBytes)
		jsonParser = json.NewDecoder(bsRead)
	} else {
		jsonParser = json.NewDecoder(charsetMap)
	}

	mapC7 := make([]c7, 0)
	if err = jsonParser.Decode(&mapC7); err != nil {
		return c7Map{}, errors.New(fmt.Sprintf("Parsing %s file: %s", src, err.Error()))
	}

	return c7Map{Map: mapC7}, nil
}

func CharSet(src string) (err error) {
	HasChange = false

	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	c7map, err := readCharSetJson(settings.MapCharset)
	if err != nil {
		return err
	}

	var buff bytes.Buffer
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		buff.WriteString(c7map.Replace(line))
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return files.Write(src, buff.Bytes())
}
