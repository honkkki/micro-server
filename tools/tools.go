package tools

import (
	"os"
	"path/filepath"
)

func GetRootDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var infer func(d string) string
	infer = func(d string) string {
		// 这里要确保项目根目录下存在 template 目录
		if exists(d + "/cmd") {
			return d
		}

		return infer(filepath.Dir(d))
	}

	return infer(cwd)
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
