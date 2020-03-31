// Code generated by go-bindata.
// sources:
// templates/html/inc_foot.html
// templates/html/inc_head.html
// templates/html/index.html
// templates/html/pip.html
// DO NOT EDIT!

package templates

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

var _templatesHtmlInc_footHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\xcc\x4b\x8e\x4f\xcb\xcf\x2f\x51\xaa\xad\xe5\xe2\xe4\xb4\xd1\x4f\xc9\x2c\xb3\xe3\xe2\xb4\xd1\x4f\xca\x4f\xa9\xb4\xe3\xb2\xd1\xcf\x28\xc9\xcd\xb1\xe3\xe4\xaa\xae\x4e\xcd\x4b\xa9\xad\xe5\x02\x04\x00\x00\xff\xff\x39\xe9\x08\xde\x39\x00\x00\x00")

func templatesHtmlInc_footHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlInc_footHtml,
		"templates/html/inc_foot.html",
	)
}

func templatesHtmlInc_footHtml() (*asset, error) {
	bytes, err := templatesHtmlInc_footHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/inc_foot.html", size: 57, mode: os.FileMode(420), modTime: time.Unix(1585680005, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlInc_headHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\xb1\x6e\x83\x30\x10\x86\x67\x78\x0a\xf7\xf6\xc0\x0b\x18\x96\xb6\x73\x3b\x74\xe9\x54\x5d\xcd\x11\x2e\x35\xb6\xe5\x73\xd3\x22\xc4\xbb\x57\x84\x44\xa2\x51\xa6\x88\x05\xac\xf3\xef\xef\x7e\xfd\x67\x8f\x63\x43\x2d\x3b\x52\xc0\xce\x7c\x74\x84\x0d\x4c\x53\xae\x1f\x9e\x5e\x1e\xdf\xde\x5f\x9f\x55\x97\x7a\x5b\xe7\x7a\xfe\x29\x8b\x6e\x5f\x01\x39\xa8\xf3\x4c\xcf\xd2\x3a\x57\x4a\xa9\x6c\xfe\xe8\xc4\xc9\x52\xbd\xf7\xbb\x9f\xce\x8b\x77\x2d\x47\x49\x3b\x09\x98\x18\xad\x2e\x97\xdd\x7c\x91\x8a\x89\x1c\x92\x4a\x43\xa0\x0a\x12\xfd\xa6\xf2\x80\x47\x5c\xaa\xa0\x24\x9a\x0a\x56\x95\x52\x2c\x87\x30\xf4\x18\x0a\x13\xbd\x48\x87\x1c\xa5\x38\x08\xd4\xba\x5c\x14\x77\x72\xad\x37\x68\x5b\x1f\x71\x4f\x45\xcf\xee\x3f\xf2\x44\xbc\x8f\xbb\x0a\xa0\x70\x94\xb6\xb0\xba\x46\x7e\x47\xbe\x61\x75\x23\xbf\xe7\x81\x15\x18\x78\x6b\xdf\x17\x74\x8f\xe1\x6a\x7c\xd9\x76\x70\x76\x7c\x15\xf8\x2a\x17\xcb\xee\x4b\x45\xb2\x15\x48\x1a\x2c\x49\x47\x94\x60\xdd\xce\x88\x80\xea\x22\xb5\x15\xcc\xeb\x9b\x1d\x4e\x9a\xf2\x72\x3f\x74\xb9\xbc\x83\x4c\x7f\xfa\x66\xb8\xa4\xd4\xf0\x51\x19\x8b\x22\x15\x18\xef\x12\xb2\xa3\x08\xe7\x23\xe3\x48\xae\x99\xa6\xfc\x2f\x00\x00\xff\xff\xce\xd3\xbb\xc6\x79\x03\x00\x00")

func templatesHtmlInc_headHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlInc_headHtml,
		"templates/html/inc_head.html",
	)
}

func templatesHtmlInc_headHtml() (*asset, error) {
	bytes, err := templatesHtmlInc_headHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/inc_head.html", size: 889, mode: os.FileMode(420), modTime: time.Unix(1585680355, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x56\x48\x49\x4d\xcb\xcc\x4b\x55\x50\xca\xcc\x4b\x49\xad\x50\x52\xa8\xad\xe5\xaa\xae\x56\x28\x49\xcd\x2d\xc8\x49\x2c\x01\x0b\x27\xc7\x67\xa4\x26\xa6\x28\x29\xe8\x81\xe4\x14\x14\x14\x14\x6c\x52\x32\xcb\x14\x92\x73\x12\x8b\x8b\x6d\x95\x8a\xf2\xcb\x95\xec\xb8\x20\xc2\xfa\x29\x99\x65\x76\x5c\x98\xda\xd3\xf2\xf3\x4b\xa0\xda\xab\xab\x15\x52\xf3\x52\x40\x2c\x40\x00\x00\x00\xff\xff\x28\xc9\x55\xba\x7a\x00\x00\x00")

func templatesHtmlIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlIndexHtml,
		"templates/html/index.html",
	)
}

func templatesHtmlIndexHtml() (*asset, error) {
	bytes, err := templatesHtmlIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/index.html", size: 122, mode: os.FileMode(420), modTime: time.Unix(1585680041, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHtmlPipHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x52\x4b\x72\xdb\x30\x0c\x5d\xdb\xa7\xc0\x60\x5d\x5b\x17\x90\x74\x95\x0e\x2d\x42\x11\x52\x0a\xe4\x10\xa0\x53\x8f\x26\x77\xef\x50\x92\x5b\x25\x69\x77\xdd\x3d\x3e\x7c\x1e\x1f\x80\x65\x01\x4f\x23\x0b\x01\xa6\xc8\x62\x2c\x29\x86\xc7\x4b\x14\x84\xf7\xf7\xf3\xb2\x80\xd1\x9c\x82\x33\x02\x64\x19\xbe\x4f\xe4\x3c\xc2\xb5\xc6\x00\x00\x5a\xcf\x77\x18\x82\x53\xed\x30\xc7\x37\xec\xcf\xa7\xf3\xe9\xb4\xb2\xec\x3b\x9c\x5d\xc2\xbe\x6d\x3c\xdf\xfb\x23\xad\xe6\xac\x28\xf6\xad\x26\x27\x2b\x33\xc4\x98\x7d\x65\x9a\x4a\xfd\xa5\xe4\xb5\xcc\xe9\x62\xf1\x32\xc6\x3c\x57\x95\x13\x40\x5b\xf1\x06\x01\x5a\x96\x54\x0c\xec\x91\xa8\x43\xa3\x9f\x86\x20\x6e\xa6\x3f\x85\xc1\x59\xa8\xa6\x8e\xcd\x9e\x5c\x0a\x6e\xa0\x29\x06\x4f\x79\x0b\x82\x45\x08\xce\xbe\xad\xd1\xbb\x0b\x85\x3a\x44\x68\x7e\x8b\xdd\x8a\x59\x94\xa7\xf1\x9b\x09\xdc\x4c\x2e\x29\xf3\xec\xf2\x03\xf7\x5f\x68\xb9\xcd\x6c\x1f\x15\xb7\x42\xec\xeb\xbb\x6d\xb6\xd7\xee\xa6\x79\xda\xf9\x6a\x3e\x93\x96\x60\xfa\xf4\x5d\xc2\x36\x33\x27\x9e\xbd\x33\xd2\x4b\x60\xb5\x3a\xbc\x12\x3e\xa6\xb0\x18\x65\xa5\xc1\x58\x5e\x3e\x25\xed\x2a\xdb\x16\x77\xdc\x06\x96\x1f\x90\x29\xd4\x15\x3d\x02\xe9\x44\x64\x78\x18\x6a\x33\xa8\x22\x4c\x99\xc6\x0e\x2b\x6e\xde\xa6\xa8\x51\x46\xce\x6a\x57\x4d\xce\xd8\x85\x6b\xe2\x74\x5d\xf3\x9a\xfe\xdc\xea\x90\x39\x1d\xf7\xd2\xbc\xba\xbb\xdb\x58\x04\xcd\x43\x87\x07\xe6\x9f\xfd\x5e\xb7\xdb\x58\x93\xfe\x5f\x57\x16\xb6\x4f\xad\xbf\x5e\xfc\x18\xa3\xed\x17\xbf\x2c\x40\xe2\x2b\xfa\x15\x00\x00\xff\xff\x18\x14\x6b\xd5\x36\x03\x00\x00")

func templatesHtmlPipHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlPipHtml,
		"templates/html/pip.html",
	)
}

func templatesHtmlPipHtml() (*asset, error) {
	bytes, err := templatesHtmlPipHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/pip.html", size: 822, mode: os.FileMode(420), modTime: time.Unix(1585680010, 0)}
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
	"templates/html/inc_foot.html": templatesHtmlInc_footHtml,
	"templates/html/inc_head.html": templatesHtmlInc_headHtml,
	"templates/html/index.html": templatesHtmlIndexHtml,
	"templates/html/pip.html": templatesHtmlPipHtml,
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
		"html": &bintree{nil, map[string]*bintree{
			"inc_foot.html": &bintree{templatesHtmlInc_footHtml, map[string]*bintree{}},
			"inc_head.html": &bintree{templatesHtmlInc_headHtml, map[string]*bintree{}},
			"index.html": &bintree{templatesHtmlIndexHtml, map[string]*bintree{}},
			"pip.html": &bintree{templatesHtmlPipHtml, map[string]*bintree{}},
		}},
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

