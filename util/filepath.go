package util

import (
	"log"
	"os"
	"path/filepath"
)

// 在当前路径及其子目录下寻找指定文件 f；如果出错，则返回空字符串。
func FindFile(f string) (p string, err error) {
	// 获取当前路径
	curPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	// 在当前路径下寻找 f
	err = filepath.Walk(curPath, func(path string, info os.FileInfo, err error) error {
		if filepath.Base(path) == f {
			p = path
			err = nil
		}
		return err
	})
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return
}
