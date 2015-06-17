package thumbnail

import (
	"fmt"
	"io/ioutil"
)

// Thumbnail temp path
var ThumbnailTempDirPath string

// MakeDir is return tempDir for thumbnails
func makeDir() (name string, err error) {
	name, err = ioutil.TempDir("", "thumbnail")
	return
}

func init() {
	fmt.Println("run thumbnail package init...")
	name, err := makeDir()
	if err != nil {
		ThumbnailTempDirPath = ""
		return
	}
	ThumbnailTempDirPath = name
}
