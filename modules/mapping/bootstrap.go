package mapping

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/rodkranz/char7/modules/chatdata"
	"github.com/rodkranz/char7/modules/setting"
)

type c7 struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// C7Map is map of things that need to replace
type C7Map struct {
	Map []c7
}

// Replace find text that needs to replace
func (cm *C7Map) Replace(line string) (bool, string) {
	HasChange := false
	for _, c := range cm.Map {
		if strings.Contains(line, c.Key) {
			HasChange = true
			line = strings.Replace(line, c.Key, c.Value, -1)
		}
	}
	return HasChange, fmt.Sprintf("%s\n", line)
}

// ReadCharSetJSON read file ".c7map" and load the configuration
func ReadCharSetJSON(src string) (C7Map, error) {
	var jsonParser *json.Decoder
	charsetMap, err := os.Open(src)
	if err != nil {
		arrBytes := chatdata.MustAsset(setting.MapCharset)
		bsRead := bytes.NewReader(arrBytes)
		jsonParser = json.NewDecoder(bsRead)
	} else {
		jsonParser = json.NewDecoder(charsetMap)
	}

	mapC7 := make([]c7, 0)
	if err = jsonParser.Decode(&mapC7); err != nil {
		return C7Map{}, fmt.Errorf("parsing %s file: %s", src, err.Error())
	}

	return C7Map{Map: mapC7}, nil
}

// GetMapping returns map of things that must be replaced
func GetMapping() (C7Map, error) {
	return ReadCharSetJSON(setting.MapCharset)
}
