// Code generated by go-bindata.
// sources:
// data/VERSION
// data/ca.pem
// DO NOT EDIT!

package shared

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

var _dataVersion = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x32\xd0\x33\xd2\x33\xe4\x02\x04\x00\x00\xff\xff\x07\x82\x30\x85\x06\x00\x00\x00")

func dataVersionBytes() ([]byte, error) {
	return bindataRead(
		_dataVersion,
		"data/VERSION",
	)
}

func dataVersion() (*asset, error) {
	bytes, err := dataVersionBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/VERSION", size: 6, mode: os.FileMode(420), modTime: time.Unix(1455656522, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataCaPem = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x93\xcb\x92\xba\x38\x1c\x85\xf7\x3c\xc5\xec\xad\x29\x01\x2f\x2d\x8b\x59\xfc\x12\xc2\x4d\x41\x03\xe1\x12\x76\x0d\x28\x20\x2a\x63\x2b\x04\x78\xfa\xb1\xbb\xab\x66\xf1\xef\x45\x67\xf9\x9d\xd4\xa9\xef\x54\x25\x7f\x7f\x1e\x44\x4c\xdb\xfb\x0b\x13\x9f\xd9\x86\x8d\x81\x91\x2f\x2a\xb9\xb6\x8d\xed\x09\x63\xe0\x02\x63\x8a\x0d\x96\x19\x6f\x89\x17\xf4\xfd\x19\x3c\x54\x36\xf7\xaa\xa9\x4d\x4d\xc8\x08\x68\x68\x80\x8e\x62\x97\x3e\x04\xa6\x5c\x8f\x28\x35\x89\x70\x90\x14\x31\xc2\xdc\x57\x39\x28\x21\xc1\x20\xb6\xa1\xaa\x3d\xd3\x40\x61\x85\x69\xc8\x29\x23\x95\x8b\x36\xdf\x59\x29\x78\x10\xaf\xe4\x34\x71\xba\x34\xa1\x65\xa4\x5e\x9a\x54\xbd\xc8\x52\x3e\x22\x5a\x58\x4d\xc9\x2c\xbf\x71\xe9\x46\xe8\xdf\xed\xba\x0e\xe9\xeb\xb2\x27\x73\x95\x08\xab\xca\x3d\x97\x51\xe1\x4d\x64\xe1\x9e\xdd\xc1\x9b\xc2\x21\xfe\x64\x67\x2a\xa4\x17\x94\xff\x87\xe7\x9f\x8a\xbf\x19\x4a\xbf\x29\xfe\x66\x28\x7d\x2a\x96\xe9\x57\xe0\x6c\xdb\xd4\xae\xfa\xdc\x03\x4a\x10\xa2\xa0\x97\x5c\x06\xd7\x36\x1d\x68\x4d\x04\x2e\x82\x85\x88\x8e\xcc\xb4\x52\x5f\xf1\xc7\xe3\xbc\x9e\x39\x9b\x77\x55\xca\xbb\x55\x99\x08\x23\x5a\xb3\xdb\x3e\x9b\x4c\xf7\x2d\xdb\x19\x78\xbf\xaa\xb5\x7e\x79\xdd\xb5\xcb\xda\x98\x0e\x0f\xeb\x1a\xeb\xe1\xb0\x0d\x16\x7c\x53\x27\x59\x27\x1f\x39\xba\xec\x5a\x8f\xcb\x97\xa0\x97\x06\x7d\x58\x2c\xf6\x33\xbd\x13\xee\x35\xf2\x66\xfa\x14\xe8\x33\xd2\x16\xda\x74\xea\xc5\x3e\xb6\x1e\xef\x13\xc7\x29\x53\x0f\x95\x37\xdc\xa2\xd8\xee\x9c\x24\xaa\xf9\xd2\x01\x2b\x34\xdb\xe2\x3c\xcc\x24\x0a\xf5\x7a\xc4\xd0\x7d\x84\xa9\xd9\xf3\xc4\x2c\xd3\x00\x41\xf9\xf2\x05\xf2\xe7\x2c\xe3\x7b\x16\x01\xda\xd2\x48\xdf\x64\xa2\x78\xc6\x8e\x04\xb7\xda\x44\x22\xc7\xa7\x3b\x1f\xad\xd9\x7c\xfb\x2f\x28\x32\xe4\xc7\xcc\x89\x22\x36\xf2\x8c\xf7\xb6\x46\x37\xc5\xda\x0f\xa2\x6e\x79\x48\x47\x6d\x0f\x3b\xcb\x9f\xd3\x8f\x18\x15\x9c\xc1\x38\x49\x27\x0f\xae\xb9\x9a\x5f\x8b\x26\xf0\x9d\xe9\x5c\xd9\xef\xfb\xc7\xf3\x46\x47\xc5\x99\x19\x8d\xcc\xb4\x64\xec\xe9\xf3\xbe\xd4\xb4\xd1\xc8\xee\xda\xd0\x96\x61\x74\x21\xa6\xb5\x56\x87\xc3\xba\xb7\x64\x2e\xad\xba\x66\x3b\x9f\x17\xf8\x00\x13\x28\x0a\xef\x12\xf5\xe3\x76\x1c\x64\xa7\xea\x18\x3d\xe5\xf6\x3f\xd2\xd7\x8b\x27\x9e\xfe\xf3\x17\xfc\x17\x00\x00\xff\xff\xc6\x69\x3c\xb1\x22\x03\x00\x00")

func dataCaPemBytes() ([]byte, error) {
	return bindataRead(
		_dataCaPem,
		"data/ca.pem",
	)
}

func dataCaPem() (*asset, error) {
	bytes, err := dataCaPemBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/ca.pem", size: 802, mode: os.FileMode(420), modTime: time.Unix(1446813769, 0)}
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
	"data/VERSION": dataVersion,
	"data/ca.pem": dataCaPem,
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
	"data": &bintree{nil, map[string]*bintree{
		"VERSION": &bintree{dataVersion, map[string]*bintree{}},
		"ca.pem": &bintree{dataCaPem, map[string]*bintree{}},
	}},
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

