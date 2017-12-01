package crypto

import (
	"io"
	"fmt"
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"github.com/qd-um/golang-crypto/blowfish"
	"github.com/klauspost/compress/zlib"
	gbytes "github.com/datochan/gcom/bytes"
	"os"
)

/**
 * MD5加密
 * param []byte data: 待加密的数据
 * return string
 */
func EncryptMd5Hex(data []byte) string{
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制

	return md5str1
}

/**
 * MD5加密(文件)
 * param string fileName: 待加密的文件
 * return string
 */
func EncryptMd5Sum(fileName string) (string, error){
	inFile, err := os.OpenFile(fileName, os.O_RDONLY | os.O_RDWR, 0666)
	if err != nil { return "", err }

	defer inFile.Close()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, inFile); err != nil {
		return "", err
	}

	strMd5 := fmt.Sprintf("%x", md5hash.Sum(nil)) //将[]byte转成16进制

	return strMd5, nil
}

func Blowfish(content []byte) []byte {
	cipher, err := blowfish.NewCipher([]byte("SECURE20031107_TDXAB"))
	if err != nil {
		return nil
	}
	offset := 0
	result := new(bytes.Buffer)
	totalLength := len(content)
	dataEncrypted := make([]byte, 8)

	for offset+8 <= totalLength {
		var tmpBigInt [2]uint32  // 用于大小端尾转换
		cipher.Encrypt(dataEncrypted, content[offset:offset+8])
		tmpBuffer := bytes.NewBuffer(dataEncrypted)
		binary.Read(tmpBuffer, binary.LittleEndian, &tmpBigInt)
		binary.Write(result, binary.BigEndian, tmpBigInt)
		offset += 8
	}

	if totalLength % 8 != 0 {
		var tmpBigInt [2]uint32
		delta := totalLength % 8
		//deltaContent := make([]byte, 8)
		deltaContent := gbytes.BytesCombine(content[offset:offset+delta], make([]byte, 8-delta))

		cipher.Encrypt(dataEncrypted, deltaContent)
		tmpBuffer := bytes.NewBuffer(dataEncrypted)
		binary.Read(tmpBuffer, binary.LittleEndian, &tmpBigInt)
		binary.Write(result, binary.BigEndian, tmpBigInt)
	}

	return result.Bytes()
}


/**
 * zlib压缩
 */
func ZLibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	defer w.Close()

	w.Write(src)
	return in.Bytes()
}

/**
 * zlib解压缩
 */
func ZLibUnCompress(compressSrc []byte) []byte {
	var out bytes.Buffer
	b := bytes.NewReader(compressSrc)
	r, _ := zlib.NewReader(b)
	defer r.Close()

	io.Copy(&out, r)

	return out.Bytes()
}


