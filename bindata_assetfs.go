package main

import (
	"github.com/elazarl/go-bindata-assetfs"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _assets_filelist_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\xc9\x28\xc9\xcd\xb1\xe3\xe2\xb4\xc9\x48\x4d\x4c\x01\xd1\xfa\x30\x46\x52\x7e\x4a\x25\x90\xe6\xac\xae\xce\x49\xcd\x53\xd0\xab\xad\x55\x78\xda\xd0\xfd\xb8\x71\xdd\xe3\xe6\x55\x8f\x9b\x76\x3e\x6e\xee\xb0\x49\x2a\x82\xc8\x17\x25\xe6\xa5\xa7\x2a\xa8\x64\xe6\xa5\xa4\x56\xe8\x28\xa8\xa4\x65\xe6\xa4\x2a\x58\xd9\x02\x75\x00\x65\x39\x6d\x12\x15\x32\x8a\x52\xd3\x6c\xd5\xf5\xab\xab\xc1\x52\x7a\x01\x89\x25\x19\x01\x40\xa1\xcc\x8a\xda\x5a\xb8\x60\x50\x6a\x62\x0e\x48\xa2\xb6\x56\xdd\x0e\x26\xe6\x06\x24\xfc\x12\x73\x53\x6b\x6b\x6d\xf4\x13\xed\xe0\xd6\xa5\xe6\xa5\x80\x8c\xb6\xd1\x87\xb8\x10\xe8\x62\xb0\x17\x00\x01\x00\x00\xff\xff\x88\xe3\x0b\x26\xca\x00\x00\x00")

func assets_filelist_html_bytes() ([]byte, error) {
	return bindata_read(
		_assets_filelist_html,
		"assets/filelist.html",
	)
}

func assets_filelist_html() (*asset, error) {
	bytes, err := assets_filelist_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "assets/filelist.html", size: 202, mode: os.FileMode(420), modTime: time.Unix(1434526809, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _assets_mp4play_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x92\x41\x6f\xdb\x3e\x0c\xc5\xcf\xc9\xa7\x20\x74\xc9\xe5\x1f\xab\xf9\xa3\x1d\x86\xc2\xf6\x71\x18\x86\xed\xb4\xde\x03\xda\x62\x23\x35\xb2\x24\x88\x74\xb2\x2c\xc8\x77\x9f\x6c\xa7\x05\x76\xd9\x29\x88\x49\x3e\xfe\x1e\x9f\x6a\x2b\x83\x6f\xd7\xab\xda\x12\x9a\xe9\xd7\xbb\x70\x04\x9b\xe9\xb5\x51\x5a\x9f\xde\xb8\xfa\x4d\xa1\x37\xa1\x0a\x24\xfa\xb1\xda\xfd\xaf\x4f\xce\x50\xdc\x96\x42\xcf\xac\x20\x93\x6f\x14\xcb\xc5\x13\x5b\x22\x51\x93\x04\xf7\xd9\x25\x01\xce\xfd\x3f\x35\xaa\x37\x56\x6d\xad\x97\xee\x69\x4e\xbf\x33\x74\xd1\x5c\xda\xf5\x04\xb5\x6b\xaf\xd7\xea\x8b\xf3\x14\x70\xa0\xdb\xad\xb4\xec\xa6\x86\x79\x1e\x9c\x69\x14\xfd\xc2\x21\x79\xda\xcf\x5f\xf6\x3b\x05\xbd\x47\xe6\x46\xbd\x53\x42\x59\xbf\x35\xf4\x8a\xa3\x97\x2d\x1f\x5d\x50\xeb\xd5\xaa\x8f\x41\x72\xf4\x0c\xa9\xe0\x47\x2c\x32\x38\x4a\x54\x70\x76\x46\x6c\xa3\x3e\x3f\x3c\x28\xb0\xe4\x0e\x56\x1a\xf5\xa9\xfc\x29\x23\x2b\x83\x82\x5b\x26\x19\x53\xb3\xb9\x7e\xac\x8d\x49\x5c\x0c\xea\x59\xf2\x48\xb7\x4d\x3b\x35\xd6\x1c\xc7\xdc\xd3\xdd\x3e\xf6\x3d\x31\xeb\xbf\x5c\x28\x90\x4b\xa2\x66\x33\x33\xea\x21\x3d\x6e\x40\x2f\xa3\xe9\x03\xbf\x50\x87\x89\x5f\xb5\x2f\x11\x4e\x8e\xce\x20\xd6\x15\x37\xb3\xf1\xb2\x19\x99\xa0\xc8\x75\x9e\xe0\x1b\x9e\xf0\xe7\x7c\xc4\xff\x00\x83\x81\xe2\x8e\x4b\x5b\x86\x31\x1d\x32\x1a\x17\x0e\x20\x11\x10\xce\xd4\x41\x97\xe3\x99\x4b\x49\x2c\x0a\xd4\x78\x0f\xda\x8a\xa4\x67\xbd\xa4\x32\x05\x1b\x07\x3d\x3d\x8b\xa7\xed\x72\x45\x1e\x53\x8a\x59\x74\xe1\xc6\x7c\xa0\x72\x94\x7d\xe7\x31\x1c\x55\x7b\xaf\x30\x7c\x7d\xf9\xf1\xfd\x69\xa1\xab\x35\x96\x54\xd3\x62\x68\xd1\x9c\xc3\x5d\x42\xad\xf5\xf2\xe0\xfe\x04\x00\x00\xff\xff\x20\x7e\x7f\x18\x78\x02\x00\x00")

func assets_mp4play_html_bytes() ([]byte, error) {
	return bindata_read(
		_assets_mp4play_html,
		"assets/mp4play.html",
	)
}

func assets_mp4play_html() (*asset, error) {
	bytes, err := assets_mp4play_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "assets/mp4play.html", size: 632, mode: os.FileMode(420), modTime: time.Unix(1434522410, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"assets/filelist.html": assets_filelist_html,
	"assets/mp4play.html": assets_mp4play_html,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"assets": &_bintree_t{nil, map[string]*_bintree_t{
		"filelist.html": &_bintree_t{assets_filelist_html, map[string]*_bintree_t{
		}},
		"mp4play.html": &_bintree_t{assets_mp4play_html, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}


func assetFS() *assetfs.AssetFS {
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: k}
	}
	panic("unreachable")
}
