package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"time"
	"wcgblog/service"
)

// 设置md5
func Md5(soure string) string {
	hash := md5.New()
	hash.Write([]byte(soure))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetCurrentTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(loc)
}

// 格式化时间
func DateFormat(date time.Time, layout string) string {
	return date.Format(layout)
}

// 截取字符串
func Truncate(s string, n int) string {
	runes := []rune(s)
	if len(runes) > n {
		return string(runes[:n])
	}
	return s
}

// 截取字符串
func Substring(source string, start, end int) string {
	rs := []rune(source)
	length := len(rs)
	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	return string(rs[start:end])
}

// 判断数字是否是偶数
func IsEven(number int) bool {
	return number%2 == 0
}

// 判断数字是否是奇数
func IsOdd(number int) bool {
	return !IsEven(number)
}

// 求和
func Add(a1, a2 int) int {
	return a1 + a2
}

// 相减
func Minus(a1, a2 int) int {
	return a1 - a2
}

func ListTag() (tagstr string) {
	tags, err := service.LisrTag()
	if err != nil {
		return
	}
	tagNames := make([]string, 0)
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}
	tagstr = strings.Join(tagNames, ",")
	return
}
