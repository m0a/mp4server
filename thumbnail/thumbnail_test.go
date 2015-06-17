package thumbnail

import (
	"testing"
)

func TestTempDirMake(t *testing.T) {
	name := ThumbnailTempDirPath
	t.Logf("name=[%s]\n", name)
}
