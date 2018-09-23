package utils

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"errors"
)

/**
 * 判断指定文件是否存在
 */
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
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

/**
 * 获取程序所在的文件夹
 */
func CurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New("error: Can't find \"/\" or \"\\\"")
	}
	return string(path[0 : i+1]), nil
}
