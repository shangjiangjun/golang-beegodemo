package models

import (
	"crypto/md5"
	"fmt"
	"time"
)

/**
# 公共函数
*/

// 时间戳转日期
func UnixToDate(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	// 时间模板 2006 1 2 3 4 5
	template := "2006-01-02 15:04:05"
	return t.Format(template)
}

func Hello(str string) string {
	return str + "world"
}

// md5 加密
func Md5(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x \n", md5.Sum(data))
}
