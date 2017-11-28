package bytes

import (
	"bytes"
)

/**
 * 拼接byte数组
 */
func BytesCombine(pBytes ...[]byte) []byte {
	// 将要拼接的数字凑成一个二维数组,通过join进行拼接
	len := len(pBytes)
	s := make([][]byte, len)
	for index := 0; index < len; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")

	return bytes.Join(s, sep)
}

/**
 * []byte转string
 * string(byteVar[:]) 由于golang中并不是像C一样遇到'\0'就自动结尾，所以需要手工转换一下
 */
func BytesToString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}
