// Code generated by go-bindata.
// sources:
// .charset
// DO NOT EDIT!

package chatdata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _Charset = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\xd7\x49\x8b\x13\x51\x14\x05\xe0\x7d\x7e\x45\x28\xc1\x55\x36\xa2\x60\xd4\x95\x84\x40\xdc\x24\x20\x71\x21\xe2\xa2\xa8\x7a\x26\x0f\x2b\x09\xd6\x10\x88\x22\x38\xcf\xf3\x6c\xcf\xf3\x3c\xcf\x73\x43\xfa\x8f\x75\x27\x8b\xee\xcd\x3d\xdc\x03\x5d\x8b\x5e\x54\xf3\xd1\x75\xce\xad\x7a\xaf\xdf\xbd\x54\xfa\xe4\x7a\xd2\xfd\xd9\xb9\x9c\x87\xa6\xe5\x5c\x4f\x3b\xed\x5d\x27\x73\x76\xb3\xe9\x06\x89\xe9\xdc\xbe\x78\xe1\x52\xf6\xea\x0d\xa7\xfb\x9b\xa7\x19\x64\xf7\xa0\xcd\xaa\x76\x1f\xda\x6b\xaa\x3d\x90\xed\x83\xd0\xf5\x2e\x5f\x51\xf5\xa1\xac\xed\xa3\xc4\x44\xb1\xa6\x8f\xfe\xca\x3a\xb6\x35\x13\xa9\x78\x53\xc6\xbe\x6d\x5a\xdf\xa8\xfa\x99\xac\x6f\x56\x42\xb7\xa9\xeb\xe7\x40\xbb\x5e\x12\xeb\xfa\x05\xd0\x9e\x0d\x3d\x15\xbf\x04\x38\xb6\x01\x11\xfb\x35\xd0\xa1\xad\x57\x54\xfc\xea\x3c\xf8\x0d\xc0\xf9\xc0\xea\xf8\xad\x8c\x73\x9e\xf1\x6d\xa0\xea\x77\xb2\xce\x73\xb3\x7e\x0f\x34\x37\xeb\x0f\x40\xe7\x98\x59\x7f\x04\x38\xa9\xe9\xa1\x3f\xc9\xf6\x16\x17\xfa\x33\xd0\x5c\xe8\x2f\x40\x53\x2f\xf8\x57\x80\x99\xd0\xdf\x40\x61\xe5\x82\x4a\xbf\xcb\xb4\xc8\x7d\x57\x3f\x64\x5d\xe2\xda\xfe\x09\x34\xd7\xf6\x2f\xa0\xa9\xb6\x7f\x03\xcc\xc5\xfe\x03\x34\x33\xab\x7f\xc0\x46\x81\x1b\x55\x55\xfd\x5f\xd6\x77\xb8\xc2\x7b\x80\xe6\x0a\xef\x05\x9a\x2a\xbc\x0f\x60\xa6\xb2\x7e\xd9\xde\xe5\x1e\x7b\x40\xd6\xe5\x42\xe9\x76\x51\xc5\x83\x32\x8e\x1e\x33\xcb\xf7\x90\x8c\x5d\x6e\x58\xc3\x40\x73\xa9\x47\x80\xa6\x86\x35\x0a\x30\xf7\x75\x8c\x03\x4d\xed\x97\x63\x32\x66\x5e\x93\x09\xf0\x77\x0d\x33\xab\x49\x19\x7b\xdc\x56\x3b\x25\x6b\xc3\x4d\x7a\x1a\x68\x6e\xd2\x33\x40\x53\x93\x9e\x05\x98\xe9\x7b\x0e\xfc\x13\xcc\x85\x9e\x07\x9a\x0b\xbd\x00\x34\x15\x7a\x11\x60\x26\xf4\x12\x28\x2c\xd6\xd7\xee\x65\x99\xd6\xb9\xaf\x6a\x45\xd6\x0d\xae\xed\x55\xa0\xb9\xb6\xd7\x80\xa6\xda\x5e\x07\x98\x8b\xbd\x01\x34\x33\xab\x2d\x60\xb9\xad\x76\x1b\xac\x44\x5c\xe1\x3b\x40\x73\x85\x83\x13\x75\x42\x15\x0e\x8e\xd4\x09\x53\x19\x38\x52\xb7\xb8\xc7\x06\x87\xea\xb8\xda\x08\xeb\x2a\x06\x67\xea\x16\xf1\xd8\x6d\xd0\x76\x23\xf4\x6b\xaa\x05\x4b\x49\x68\x4e\x37\x8d\xd4\xfd\xd4\x71\x00\x00\x00\xff\xff\xca\x16\x8f\x2c\xfc\x10\x00\x00")

func CharsetBytes() ([]byte, error) {
	return bindataRead(
		_Charset,
		".charset",
	)
}

func Charset() (*asset, error) {
	bytes, err := CharsetBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".charset", size: 4348, mode: os.FileMode(420), modTime: time.Unix(1476170433, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	".charset": Charset,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	".charset": &bintree{Charset, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

