package zu

import (
	"os"
	"os/exec"
)

// FileExists ...
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Brotli ...
func Brotli(path string) (err error) {
	return exec.Command(`/usr/bin/brotli`, `-Zfk`, path).Run()
}

// Gzip ...
func Gzip(path string) (err error) {
	return exec.Command(`/usr/bin/gzip`, `--best`, `-k`, path).Run()
}
