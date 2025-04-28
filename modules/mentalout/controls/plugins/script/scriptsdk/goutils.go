package scriptsdk

import (
	"crypto/md5"
	"github.com/gogf/gf/v2/encoding/gcharset"
	"github.com/google/uuid"
	"golang.org/x/text/encoding/unicode"
)

// 生成 `uuid` V4 字符串
func UUIDv4() string {
	return uuid.New().String()
}

// 中文编码转换
func GBKToUTF8(src string) (dst string) {
	var err error
	dst, err = gcharset.Convert("UTF-8", "GBK", src)
	if err != nil {
		dst = src
	}
	return
}

func UTF8ToGBK(src string) (dst string) {
	var err error
	dst, err = gcharset.Convert("GBK", "UTF-8", src)
	if err != nil {
		dst = src
	}
	return
}
func Md5(data []byte) (hash []byte) {
	tmp := md5.Sum(data)
	hash = tmp[:]
	return
}

func Unicode(src []byte, littleEndian bool, ignoreBOM bool) (dst []byte) {
	var endian unicode.Endianness
	if littleEndian {
		endian = unicode.LittleEndian
	} else {
		endian = unicode.BigEndian
	}
	var bom unicode.BOMPolicy
	if ignoreBOM {
		bom = unicode.IgnoreBOM
	} else {
		bom = unicode.UseBOM
	}
	dst, _ = unicode.UTF16(endian, bom).NewEncoder().Bytes(src)
	return
}
