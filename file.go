package zu

import (
	"os"
	"os/exec"
	"path"
)

// 临时文件相关配置 ...
var (
	TempDir         = os.TempDir()
	TempFilePattern = `tmp-go-*`
)

// FileExists ...
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Brotli ...
func Brotli(path string) (err error) {
	return exec.Command(`/usr/bin/brotli`, `--best`, `--keep`, `--force`, path).Run()
}

// Gzip ...
func Gzip(path string) (err error) {
	return exec.Command(`/usr/bin/gzip`, `--best`, `--keep`, `--force`, path).Run()
}

// WriteFile ...
func WriteFile[T []byte | string](file string, content T) (err error) {

	if !IsDir(path.Dir(file)) {
		// 提前判断目录有些降低效率，但可以避免没有建目录时生成过多临时文件
		return os.ErrPermission
	}

	tmpName, err := WriteTemp(content)
	if err != nil {
		return
	}

	err = os.Chmod(tmpName, 0644)
	if err != nil {
		return
	}

	// 这个函数存在的意义在这一步，改名是个原子操作，避免程序异常导致文件写不完整
	err = os.Rename(tmpName, file)
	if err != nil {
		os.Remove(tmpName)
		return
	}
	return
}

// WriteTemp ...
func WriteTemp[T []byte | string](content T) (filename string, err error) {

	f, err := os.CreateTemp(TempDir, TempFilePattern)
	if err != nil {
		return
	}

	switch v := any(content).(type) {
	case string:
		_, err = f.WriteString(v)
	case []byte:
		_, err = f.Write(v)
	}
	f.Close()
	if err != nil {
		return
	}
	return f.Name(), nil
}

// IsDir 确认目录存在
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
