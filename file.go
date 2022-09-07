package toolkit

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// ExistFile 判断文件是否存在
func ExistFile(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

// ExistDir 判断目录是否存在
func ExistDir(name string) (bool, error) {
	fileInfo, err := os.Stat(name)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

// ReadDir 读取目录的内容，分别返回目录，文件的全路径集合
func ReadDir(path string) ([]string, []string, error) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, nil, errors.Wrap(err, "call ReadDir error")
	}
	files := make([]string, 0, len(fileInfos))
	dirs := make([]string, 0, len(fileInfos))

	for _, fileInfo := range fileInfos {
		fullPath := filepath.Join(path, fileInfo.Name())
		if fileInfo.IsDir() {
			dirs = append(dirs, fullPath)
		} else {
			files = append(files, fullPath)
		}
	}
	return dirs, files, nil
}

// FileSizeByPath 通过路径计算文件的大小（Byte单位)
func FileSizeByPath(path string) (int64, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return -1, err
	}
	return fileInfo.Size(), nil
}

// FileSize 通过打开的文件计算大小（Byte单位）
func FileSize(file *os.File) (int64, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return -1, err
	}
	return fileInfo.Size(), nil
}

// Mkdirs 根据路径创建目录
func Mkdirs(path string, perms ...os.FileMode) error {
	var perm os.FileMode = 0755
	if len(perms) > 0 {
		perm = perms[0]
	}
	return os.MkdirAll(path, perm)
}

// SplitFile 获取文件的目录和文件名
func SplitFile(path string) (string, string, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return "", "", errors.Wrap(err, "call SplitFile fail, abs filepath error")
	}
	dir, filename := filepath.Split(path)
	return dir, filename, nil
}
