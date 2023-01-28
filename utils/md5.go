package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// Md5Crypt 给字符串生成md5
func Md5Crypt(str string, salt ...interface{}) (CryptStr string) {
	if n := len(salt); n > 0 {
		slice := make([]string, n+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
