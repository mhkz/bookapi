package file

import (
	"os"
	"path"
	"mime/multipart"
	"io/ioutil"
)

// 获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckExist
func CheckExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// CheckPermission
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir
func IsNotExistMkDir(src string) error {
	if exist := CheckExist(src); exist == false {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MkDir
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Open
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}