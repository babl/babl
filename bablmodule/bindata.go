// Code generated by go-bindata.
// sources:
// data/ca.pem
// DO NOT EDIT!

package bablmodule

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

var _caPem = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x93\x4b\x8f\xb2\x48\x14\x86\xf7\xfc\x8a\xd9\x9b\x89\x88\x97\x96\xc5\x2c\x4e\x15\xc5\x4d\x29\x2d\x29\x2e\x55\xbb\x06\x14\x11\x95\xb1\x15\x0a\xf8\xf5\x93\xee\x4e\x66\xf1\xf5\xa2\xcf\xf2\x79\x93\x93\xe7\x4d\xce\xf9\xfb\x73\x10\x71\x3c\xfa\x17\x26\x07\xee\xd9\x1e\x06\x4e\xbe\xa8\x16\x78\x1e\xf6\x46\x8c\x41\x28\x8c\x19\xb6\x79\x66\xbf\xa5\x34\xec\xba\x0b\x50\x54\xd6\x8f\x73\x5d\x39\xa6\xd2\x11\xb0\xc8\x06\x0b\x25\x01\x7b\x2a\xcc\x84\x15\x33\xe6\x10\xe5\x23\x2d\xe6\x84\x07\x88\x38\x30\x8b\x08\x06\xb5\x89\x0c\xf3\x25\xc3\x19\x2f\x1c\x5b\x97\x9c\x9c\x03\xb4\xfe\xce\x4a\x25\xc2\x64\xa9\xcb\xd4\x6f\x65\xca\xca\xd8\xb8\xd6\xd2\xb8\xea\x5a\x3e\x20\x56\xb8\x75\xc9\xdd\x43\x1d\xb0\xb5\xb2\xbe\xb7\x5b\x16\x48\x5d\xa6\x54\x17\x06\x51\xee\x39\xa7\x01\x67\x8a\x8e\x64\x1e\x5c\x82\x9e\x8e\x51\x9f\x7c\xb2\x0b\x53\x1a\x1d\x89\xfe\x3f\xbc\xfc\x54\xfc\xcd\x50\xfb\x4d\xf1\x37\x43\xed\x53\xb1\x94\x5f\x81\xbf\x69\xa4\x77\xee\x72\x0a\x8c\x20\xc4\xc0\x2a\x85\x0e\x81\xe7\xf8\xd0\x38\x08\x02\x04\x73\x15\x1f\xb9\xe3\xca\xc3\xec\x30\x1c\xa7\xd5\xc4\x5f\xbf\x1b\x5a\xde\x2e\xcb\x54\xd9\xf1\x8a\xdf\x77\xd9\xe8\x04\x6f\xd9\xd6\xc6\xbb\x65\x65\x76\x8b\xdb\xb6\x59\x54\xf6\xb8\x7f\xba\xb7\xc4\x8a\xfa\x4d\x38\x17\xeb\x2a\xcd\x5a\xfd\x28\xd0\x75\xdb\x50\xa1\x5f\xc3\x4e\xeb\xad\x7e\x3e\xdf\x4d\xac\x56\x05\xb7\x98\x4e\xac\x31\xb4\x26\xa4\x29\xcc\xf1\xd4\xa9\x5d\xe2\x3e\xdf\x47\x81\x25\x37\xf6\x67\xda\xdf\xe3\xc4\x6b\xfd\x34\xae\xc4\xc2\x07\x37\x72\x9a\xe2\xd2\x4f\x34\x06\xd5\x6a\xc0\xd0\x7e\x44\xd2\xe9\x44\xea\x94\x32\x44\x50\x06\x08\x80\xfc\x59\xcb\xfe\xae\x45\x80\x35\x2c\xb6\xd6\x99\x2a\x5e\x89\xaf\xc1\xbd\x72\x90\xca\xf1\xe9\x21\x06\x77\x32\xdd\xfc\x0b\x33\x1d\xf2\x63\xe6\xc7\x31\x1f\x44\x26\x3a\xcf\x64\xeb\x62\x75\x08\xe3\x76\xb1\x97\x83\xb9\x83\xad\x7b\x98\xb2\x8f\x04\x15\x82\xc3\x30\x6a\x27\x0a\xb7\xdc\xc8\x6f\x45\x1d\x1e\xfc\xf1\x72\xf6\xde\x77\xcf\xd7\x9d\x0d\x33\x7f\x62\xd7\x3a\x37\xd3\xa1\x63\xaf\xc7\xc2\x34\x07\x3b\x7b\x98\x7d\x53\x46\xf1\x95\x38\xee\xca\xe8\xf7\xab\xce\xd5\x85\xb6\x6c\xeb\xcd\x74\x5a\xe0\x3d\x8c\x30\x9b\x89\x36\x35\x3e\xee\xc7\x5e\xf7\xcf\x2d\x67\xa7\xdc\xfb\x47\xfb\xba\x78\x42\xad\x9f\x5f\xf0\x5f\x00\x00\x00\xff\xff\xc6\x69\x3c\xb1\x22\x03\x00\x00")

func caPemBytes() ([]byte, error) {
	return bindataRead(
		_caPem,
		"ca.pem",
	)
}

func caPem() (*asset, error) {
	bytes, err := caPemBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ca.pem", size: 802, mode: os.FileMode(420), modTime: time.Unix(1475170859, 0)}
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
	"ca.pem": caPem,
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
	"ca.pem": &bintree{caPem, map[string]*bintree{}},
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

