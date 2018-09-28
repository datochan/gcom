package utils

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
	"path/filepath"
)

/**
 * 获取当前应用程序所在目录
 */
func GetExeDir() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir + "/"
}

/**
 * 判断指定文件是否存在
 */
func IsFileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

/**
 * 判断指定目录是否存在
 */
func IsDirExist(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
	return true
}

/**
 获取指定目录中的文件列表
 :param file_path: 指定的路径
 :return: []string, error
 :notice
 只遍历文件忽略文件夹
 */
func FileListInPath(filePath string) ([]string, error) {
	var resultList []string
	dirList, err := ioutil.ReadDir(filePath)
	if err != nil {
		fmt.Println("read dir error")
		return nil, err
	}

	for _, v := range dirList {
		fileName := v.Name()
		fileMode := v.Mode()
		if 0 == strings.Index(fileName, ".") || fileMode.IsDir() {
			continue
		}

		resultList = append(resultList, fileName)
	}

	return resultList, nil
}
