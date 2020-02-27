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

// 在当前路径及其子目录下寻找指定目录
func FindDir(dir string) (d string, err error) {
    // 获取当前路径
	curDir, err := os.Getwd()
	CheckErr(err)
	// 寻找 dir 路径
	err = filepath.Walk(curDir, func(p string, i os.FileInfo, err error) error {
		if i.IsDir() && filepath.Base(p) == dir {
			d = p
			err = nil
		}
		return err
	})
	CheckErr(err)
	return
}
