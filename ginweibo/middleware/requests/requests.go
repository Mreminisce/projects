package requests

import (
	"ginweibo/utils/view"
	"regexp"
	"strconv"
	"strings"
)

type (
	ValidatorFunc   = func() (msg string)        // 验证器函数
	ValidatorMap    = map[string][]ValidatorFunc // 验证器数组 map
	ValidatorMsgArr = map[string][]string        // 错误信息数组
)

func RunValidators(m ValidatorMap, msgMap ValidatorMsgArr) (errors []string) {
	for k, validators := range m {
		customMsgArr := msgMap[k] // 自定义错误信息数组
		customMsgArrLen := len(customMsgArr)
		for i, fn := range validators {
			msg := fn()
			if msg != "" {
				if i < customMsgArrLen && customMsgArr[i] != "" {
					msg = customMsgArr[i] // 采用自定义的错误信息输出
				} else {
					names := strings.Split(k, "|") // 采用默认的错误信息输出
					data := make(map[string]string)
					for ti, tv := range names {
						data["$key"+strconv.Itoa(ti+1)+"$"] = tv
					}
					msg = view.ParseEasyTemplate(msg, data)
				}
				errors = append(errors, msg)
				break // 进行下一个字段的验证
			}
		}
	}
	return errors
}

// value 必须存在
func RequiredValidator(value string) ValidatorFunc {
	return func() (msg string) {
		if value == "" {
			return "$key1$ 必须存在"
		}
		return ""
	}
}

func MixLengthValidator(value string, minStrLen int) ValidatorFunc {
	return func() (msg string) {
		l := len(value)
		if l < minStrLen {
			return "$key1$ 必须大于 " + strconv.Itoa(minStrLen)
		}
		return ""
	}
}

func MaxLengthValidator(value string, maxStrLen int) ValidatorFunc {
	return func() (msg string) {
		l := len(value)
		if l > maxStrLen {
			return "$key1$ 必须小于 " + strconv.Itoa(maxStrLen)
		}
		return ""
	}
}

func EqualValidator(v1 string, v2 string) ValidatorFunc {
	return func() (msg string) {
		if v1 != v2 {
			return "$key1$ 必须等于 $key2$"
		}
		return ""
	}
}

func EmailValidator(value string) ValidatorFunc {
	return func() (msg string) {
		pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` // 匹配电子邮箱
		reg := regexp.MustCompile(pattern)
		status := reg.MatchString(value)
		if !status {
			return "$key1$ 邮箱格式错误"
		}
		return ""
	}
}
