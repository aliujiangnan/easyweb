package util

import (
	"fmt"
    "crypto/md5"
)

var authPriKey = "~@!#$!@$@^&$*#"

func Sign(str string) string {
    data := []byte(str + authPriKey)
    has := md5.Sum(data)
    md5str := fmt.Sprintf("%x", has)
    return md5str
}