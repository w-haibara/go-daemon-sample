package util

import (
	"os"
	"path/filepath"
)

var (
	UnixDomainPath = filepath.Join(os.TempDir(), "go-daemon-sample.sock")
)
