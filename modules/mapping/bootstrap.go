package mapping

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/rodkranz/char7/modules/chatdata"
	"github.com/rodkranz/char7/modules/files"
	"github.com/rodkranz/char7/modules/settings"
)

type c7 struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type c7Map struct {
	Map []c7
}

func (cm *c7Map) Replace(line string) (bool, string) {
	HasChange := false
	for _, c := range cm.Map {
		if strings.Contains(line, c.Key) {
			HasChange = true
			line = strings.Replace(line, c.Key, c.Value, -1)
		}
	}
	return HasChange, fmt.Sprintf("%s\n", line)
}

func ReadCharSetJson(src string) (c7Map, error) {
	var jsonParser *json.Decoder
	charsetMap, err := os.Open(src)
	if err != nil {
		arrBytes := chatdata.MustAsset(settings.MapCharset)
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

func GetMapping() (c7Map, error) {
	return ReadCharSetJson(settings.MapCharset)
}

func init() {
	mapPath := settings.HomeDir + "/" + settings.MapCharset
	if _, err := os.Stat(mapPath); os.IsNotExist(err) {
		files.Write(mapPath, chatdata.MustAsset(settings.MapCharset))
	}
}
