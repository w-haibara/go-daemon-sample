package util

import (
	"os"
	"path/filepath"
)

var (
	UnixDomainPath = filepath.Join(os.TempDir(), "go-deamon-sample.sock")
)
