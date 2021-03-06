// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/run
// templates/server.properties
package mcconfig

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

var _templatesRun = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\x41\x6f\x9b\x40\x10\x85\xef\xf9\x15\x54\x95\x7a\xa9\x0c\xc6\x49\xa5\x36\x12\x95\x2a\xdb\xa2\x95\x42\x83\x30\x69\x73\x9d\x2c\xe3\x65\x92\x65\x17\xed\xec\x62\xdc\x28\xff\xbd\x02\x83\x55\x59\xbd\xb1\xef\x7d\xf3\x1e\x9a\x79\xff\x2e\x7a\x22\x1d\x71\x7d\x75\xf5\x0c\x1d\x04\x8b\xc7\x86\x5f\x5f\x83\x30\xc3\xc6\xd8\xe3\x1d\x35\xe4\x82\xb7\xb7\x6c\xd0\xfb\xff\xeb\x8f\xb7\x1f\x1f\x18\xd3\x38\x5d\x9f\x1e\x39\x58\x50\x0a\x55\x81\xfb\xdc\x1a\xb1\xd5\xf0\xa4\xb0\x1a\xbd\x0c\xfa\x74\x9d\x83\x67\xcc\x48\x29\xe2\x64\xb5\x5c\x4e\x09\x5a\x19\xf1\xb2\xed\x5b\xb4\xd4\xa0\x76\xa0\x7e\x65\xf7\xad\x23\xa3\xf9\x04\x6c\x88\x87\x9c\x6d\xdf\x2a\x12\xe4\xe6\xb2\x6f\xea\x00\x47\xce\x2d\x96\xc6\x8b\x7a\xd4\xd2\xf8\x27\x1e\x76\xf4\x07\x73\xb4\x02\xb5\x4b\xae\x97\x93\x9e\x41\x7f\x61\xdd\xcc\xd6\x77\x84\xb6\x40\x49\x46\x0f\x76\xf2\x39\x9b\xf4\x02\x19\x6d\x77\xe6\x57\xff\xf2\xbf\x81\xdd\xd9\xf9\x34\x77\x50\x8f\x55\xba\x5e\x1b\xaf\x5d\x09\x56\xa2\x4b\x6e\x46\xeb\x87\x26\x47\xe0\x48\xcb\x61\xf6\x5e\x08\xdf\x82\x16\xc7\x79\x3e\xbe\x08\xb8\xa3\x0e\xcb\xda\x22\xd7\x46\x55\x33\xf4\x65\xae\x2f\x76\xe8\x1e\xda\x6a\x8c\x1b\xf7\x59\x52\x73\xf1\x2b\x3b\x6f\x3b\xea\x8c\x2d\xc0\x91\x49\xae\x57\xd3\x71\xd0\xee\xa7\x55\xee\x6a\xb0\x58\x65\xd8\xcc\xa7\x29\x51\x7b\x4b\x5a\x9e\x6b\x93\x38\x58\x6c\x3c\x93\x96\x21\xd0\x0b\x58\x0e\xf7\x0a\x24\x27\xb5\x73\x2d\xdf\x46\x51\x23\xc6\x77\x88\x8d\x08\x25\x07\x8b\xcd\x44\x69\x3c\x4c\xa4\xb3\x1e\x83\xc5\x33\xd8\x20\xe2\x96\xa4\x71\xe1\xf0\xad\x8d\xf4\x14\xac\xbe\x7e\x88\xff\x06\x00\x00\xff\xff\xf0\xa1\x43\x62\x7e\x02\x00\x00")

func templatesRunBytes() ([]byte, error) {
	return bindataRead(
		_templatesRun,
		"templates/run",
	)
}

func templatesRun() (*asset, error) {
	bytes, err := templatesRunBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/run", size: 638, mode: os.FileMode(493), modTime: time.Unix(1608272086, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesServerProperties = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x56\xc1\x72\xdb\x38\x0f\xbe\xeb\x29\xfc\x02\xec\x38\x99\x26\xff\xbf\x07\x1f\xda\x74\xdb\x66\xa6\xce\x7a\xd7\x99\xed\x99\xa6\x20\x8b\x6b\x8a\xe4\x92\x90\x6c\xd7\xa3\x77\xdf\xa1\x48\x4a\xa4\xe3\xc8\x27\xe3\x83\xe6\xfb\x40\x10\x04\x60\x35\x3d\x4a\xa2\x8d\x42\x60\xc8\x95\x5c\x5d\x2e\x0b\x5e\x2d\x3e\x6c\x1d\xbe\x19\xe1\x45\xdf\x5f\x2e\xb7\x51\xb2\x00\x61\x61\xd1\xf7\x77\x8f\x83\x21\xcb\x45\xdf\x17\x0d\x3d\x11\xe4\xec\x40\x90\x37\xb0\x7a\x5c\x2e\x97\xcb\xe2\xdf\x16\xcc\xf9\x83\x56\x06\xa3\xca\x9f\x0e\xd9\x28\x83\x81\x3f\xb7\x47\xe6\xfb\x87\x87\xc7\x87\x40\xde\xf7\xc5\x1e\x24\x18\x8a\xca\x10\x0b\x88\x5c\xee\x6d\xe4\xfb\x16\x3d\xdb\xe0\x08\xbc\xb7\xf1\x31\x58\x7b\x96\x8c\xb0\xba\x95\x07\x72\x34\x1c\xc1\xae\xd0\xb4\x50\x54\xca\x30\x20\x7b\xda\x40\xa3\x4a\x58\x5d\x2e\x4e\xe2\xab\x03\xbf\x05\xcc\xb3\x67\x50\x1e\x79\x45\x85\x85\x29\x72\x2a\x84\x3a\x12\x09\x58\x83\xf1\x1a\x20\xbd\xca\xb1\xe6\x08\x82\x5b\x0c\x32\xbf\x7b\xfc\x67\x84\xbd\xd2\x35\x9a\x88\xbd\xd1\xba\x8a\x3b\x0f\xf9\x46\xb4\x7d\x6f\x5b\xd3\xf1\x8e\x8a\x89\x63\x67\x14\x2d\x19\xb5\x48\x98\x92\x56\x09\x20\xa8\x88\xd2\x36\xc6\x4e\x77\x02\xc8\x70\xad\x63\xd8\x0e\x1a\xae\x31\x46\x3c\x02\x73\xc1\x6a\x41\xcf\x60\x08\x2f\x9d\x04\x6f\x40\xb5\x31\x11\x9b\xc1\xf3\x5c\x0a\x78\xf5\xb8\xe7\x7d\x03\x67\xec\xcb\x89\x19\xe1\x84\xa4\xe2\x02\xc1\x70\xb9\x77\xe7\xa8\xf8\x7e\x55\x94\xbc\xaa\x38\x6b\x05\xc6\xc8\xbf\x8c\x80\x17\x98\xec\x8c\x19\xa8\x3d\xdf\xca\x8f\x61\x4a\x66\xc9\xf1\x0f\xab\x51\xd2\x22\x18\x1b\x44\x86\xf7\xb3\x0e\x58\xf2\xa4\x22\x94\x49\x39\x9a\x49\x4a\x69\xa2\xc1\x34\xdc\x5a\xae\x24\x11\xd0\x81\x58\x7d\x2c\x74\xa7\x63\x9e\x3a\x1d\x32\xd3\xe9\x19\x1a\x90\xc8\xf1\x4c\x92\xc0\xa9\xdc\x83\xa3\x66\x20\x91\xee\x61\x75\xb7\x5c\x16\x56\x2a\xa5\xc1\x10\x7f\xc3\xe5\x6a\xb8\xad\x62\x10\x25\x78\xd6\xb1\xa6\x7e\x38\xe0\xf5\xac\x43\x51\x8d\x66\xa6\x5f\x42\x45\x5b\x81\x69\x08\x43\xd9\x58\xa4\xd8\xda\xac\x6e\xb6\x03\x94\x16\x8e\x47\x66\xce\x53\x53\x53\x32\x65\x62\x44\xdf\x83\xe9\x39\xa2\x35\x57\x79\x21\x1a\xa6\x9a\x86\xca\x92\xec\x84\x62\x87\x2c\xa8\x27\xef\xf9\xec\x1c\x69\x68\x29\x3e\x27\x20\x01\x8f\xca\x1c\x9c\x82\x36\xe0\xaf\x0f\x6b\x03\xb6\x56\xa2\x0c\x4a\x2f\xfe\x9b\xa7\xe9\x93\xd7\xf8\x85\x97\x9c\xf9\x20\xd3\xbe\x7f\x78\x9c\x94\x5d\x0f\xf6\x0f\x2b\xa6\x79\x4d\x4f\xfe\xe1\x84\x24\x4f\x76\xce\xb2\xcc\x49\x8e\xca\x88\x92\x58\xfe\x0b\x26\x9e\x9f\x0e\xdb\xf2\x5f\x30\x32\x8d\x48\xce\xf5\x9b\xfb\xfd\xff\xe3\xc4\x68\xc0\xaa\xd6\xf5\x3c\x4d\xd9\x81\xd8\x9a\xde\x05\xd6\xbf\x82\x63\x43\xd9\x61\x5b\xd3\x3b\xcf\x7c\x8d\x4e\xad\xbb\xef\x8b\xaa\x95\xc3\x24\x7a\xfb\x36\xee\x0b\xf7\x24\xe3\xb4\x19\xe8\x99\x92\x6e\xb6\x04\xda\x60\x5d\xa5\xef\xe1\x7f\xc9\xa0\xb1\x60\x3a\x30\x24\xe1\xd8\x0e\xc8\xc4\x32\xd9\xd7\x3c\xe9\xc0\x2a\x61\xd7\xee\x63\x9f\x71\xff\x43\x8b\x71\x7f\xe7\x4a\x27\xe8\x73\x9d\xa9\x3f\xeb\x54\xfb\x59\x67\x09\xf1\x6d\x47\x6a\x96\xb5\x9c\x17\xcd\xd2\x76\xe3\xcc\x99\x37\xe5\xa7\x54\x25\xf8\xbe\x8e\xe7\xfe\xe4\xa0\xaf\x03\xe2\x79\x12\x60\xee\x04\xbe\x65\x48\xda\x64\x2d\xe3\x85\x36\x69\xcb\x70\x66\x46\x32\xd4\xdb\x44\xd2\x71\x38\x92\x92\x5b\xa4\x92\x45\x9e\xbf\x39\x1c\xbf\x04\xc8\x53\xa5\x48\xc6\x76\xb7\x7c\xa7\xf6\x6e\x94\xdd\xdb\x92\xbb\x91\x5d\x2a\x79\x43\x45\x96\xe0\x4f\x1e\x4a\x72\x1c\x90\x99\x34\x0f\x23\x9f\x24\x33\x7f\x18\xeb\x3f\xc6\x61\x3f\x9a\x73\x09\xf6\x45\x4e\xad\x3d\x2a\x53\xa6\x85\x1e\xa0\xa4\xd8\x03\x92\x1d\x28\xac\x52\xae\x1b\x9b\x96\x61\x6b\x20\x1e\x2b\xac\x4c\xb0\x1d\x1d\xd9\x2a\x95\xe0\x73\x43\x4b\x0a\x2e\x81\x24\x6b\xc8\x1f\x03\xb2\x1e\x17\x91\xc9\x9e\xa1\x71\x3d\x68\xd7\x72\x51\x92\x1a\x92\xa2\x5c\xd3\xd3\x67\x87\x7e\x87\xa9\x2e\x73\xec\xfd\xde\xe8\x0b\xd3\x02\x94\x69\x61\x6e\x01\xca\xa4\x30\x9d\x99\xa5\x4b\x1b\xe8\x40\xa2\xdb\x97\x4f\x67\xb7\x4b\x48\xbf\x08\xc7\xa4\x6d\xbc\x7f\xe3\xdc\x4f\x93\x37\x0c\xe6\xdb\xce\xb9\xdb\x6d\x2d\x10\x49\x91\x77\x40\xd0\x50\x69\x87\x4e\x94\xee\x5e\xff\x34\x27\xb7\x63\x70\x54\x6e\xbb\x09\x73\xba\x51\x18\x4f\xb5\x56\x18\x0e\xe4\xfe\x65\x67\x19\xae\x5d\xf0\x86\x8f\xfd\x91\xba\x72\x6b\x78\x6c\x90\xd1\x7c\x6f\xb3\x0a\x21\xb8\x0a\xcc\x06\xa6\x2b\xb6\x74\x50\x3a\x7b\xee\x90\xff\x05\x00\x00\xff\xff\xd4\x1c\xb4\x97\x83\x0c\x00\x00")

func templatesServerPropertiesBytes() ([]byte, error) {
	return bindataRead(
		_templatesServerProperties,
		"templates/server.properties",
	)
}

func templatesServerProperties() (*asset, error) {
	bytes, err := templatesServerPropertiesBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/server.properties", size: 3203, mode: os.FileMode(420), modTime: time.Unix(1608309903, 0)}
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
	"templates/run": templatesRun,
	"templates/server.properties": templatesServerProperties,
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
	"templates": &bintree{nil, map[string]*bintree{
		"run": &bintree{templatesRun, map[string]*bintree{}},
		"server.properties": &bintree{templatesServerProperties, map[string]*bintree{}},
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

