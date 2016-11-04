// Package mapping create a map list of characters that will use for transformation
package mapping

import (
	"os"
	"fmt"
	"bytes"
	"errors"
	"strings"
	"encoding/json"

	"github.com/rodkranz/char7/modules/files"
	"github.com/rodkranz/char7/modules/chatdata"
	"github.com/rodkranz/char7/modules/settings"
)

// C7 is an entity contains key and value that will looking for and replace
// when execute the transformation.
type C7 struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// C7Map is a map with C7
type c7Map struct {
	Map []C7
}

// Replace returns bool (if has changed) and string (new line changed)
// receive a original line and search for word to change
// and returns if has any change and the line changed.
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

// ReadCharSetJson returns c7Map with list of c7
// it reads .c7map and fill c7map with those values, returns
// c7map filled and error.
// Usage:
//   ReadCharSetJson("/home/kranz/.c7map") would return string (new line with changes) and error
//   c7map, err := ReadCharSetJson("/home/kranz/.c7map")
//   if err != nil {
//      fmt.Errorf("Error to try read file: %v", err.Error())
//   }
//   line := "coração";
//   if has, newLine := c7map.Replace(line); has {
//      // Must return "cora&ccedil;&atilde;o"
//      fmt.Printf("[Changed] : %v", newLine)
//   } else {
//      fmt.Printf("[Original]: %v", line)
//   }
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

	mapC7 := make([]C7, 0)
	if err = jsonParser.Decode(&mapC7); err != nil {
		return c7Map{}, errors.New(fmt.Sprintf("Parsing %s file: %s", src, err.Error()))
	}

	return c7Map{Map: mapC7}, nil
}

// GetMapping return C7map filled and error
// but this function you don't specify the path
// the path get from settings, by default is from "$HOME/.c7map"
func GetMapping() (c7Map, error) {
	return ReadCharSetJson(settings.MapCharset)
}

func init() {
	mapPath := settings.HomeDir + "/" + settings.MapCharset
	if _, err := os.Stat(mapPath); os.IsNotExist(err) {
		files.Write(mapPath, chatdata.MustAsset(settings.MapCharset))
	}
}