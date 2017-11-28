package utils

import (
	"fmt"
	"time"
	"encoding/hex"
	"crypto/rand"
	mrand "math/rand"
	"github.com/Tang-RoseChild/mahonia"
)

/**
 * 字符串转换编码
 * param string src: 待转换的字符串
 * param string strCode: 源编码
 * param string tagCode: 目标编码
 */
func ConvertTo(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)

	return string(cdata)
}

// 从 slice 中查找指定的字符串
// 返回相应的索引值，找不到返回-1
func FindInStringSlice(str string, s []string) int {
	for i, e := range s {
		if e == str {
			return i
		}
	}
	return -1
}

// 生成UUID字符串
func NewUUID() string {
	u := [16]byte{}
	rand.Read(u[:])
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (4 << 4)
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

// 生成MAC地址
func RandomMacAddress() string {
	mrand.Seed(time.Now().UnixNano())
	mac := [6]byte{0x80, 0x3f,
					byte(mrand.Intn(0x7F)), byte(mrand.Intn(0x7F)),
					byte(mrand.Intn(0x7F)), byte(mrand.Intn(0x7F))}

	return hex.EncodeToString(mac[:])
}
