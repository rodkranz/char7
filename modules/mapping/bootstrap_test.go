package mapping

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/rodkranz/char7/modules/setting"
)

func Test_C7Config_Should_Exists_In_Home(t *testing.T) {

	mapPath := setting.HomeDir + "/" + setting.MapCharset
	if _, err := os.Stat(mapPath); os.IsNotExist(err) {
		t.Errorf("Expected that %v exists in %v.", setting.MapCharset, setting.HomeDir)
	}
}

func Test_C7Config_Should_Return_Default_Mapping(t *testing.T) {
	// Create Fake Map in temp dir
	f, err := ioutil.TempFile("", "charset-test")
	if err != nil {
		t.Fatal(err)
	}
	defText := []byte(`[ { "key": "»", "value": "&#187;" }, { "key": "¼", "value": "&#188;"} ]`)
	filename := f.Name()
	if err := ioutil.WriteFile(filename, defText, 0644); err != nil {
		t.Fatalf("WriteFile %s: %v", filename, err)
	}
	f.Close()

	// Mirror C7Map of mapping file-
	expectedC7Map := C7Map{
		Map: []c7{
			{Key: "»", Value: "&#187;"},
			{Key: "¼", Value: "&#188;"},
		},
	}

	//
	c7m, err := ReadCharSetJSON(filename)
	if err != nil {
		t.Error("Expected got a error message but got a nil")
	}

m1:
	for _, c7 := range c7m.Map {
		found := false
		for _, c7F := range expectedC7Map.Map {
			if c7F.Key == c7.Key {
				found = true
				if c7F.Value != c7.Value {
					t.Errorf("Map %v expected %v but got %v.", c7.Key, c7.Value, c7F.Value)
					continue m1
				}
			}
		}
		if !found {
			t.Errorf("Expected to find %v with value %v but got nothing.", c7.Key, c7.Value)
		}
	}
}

//func Test_C7Config_Should_Read_From_Home(t *testing.T) {
//
//}
