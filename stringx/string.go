package stringx

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"time"
	"unicode"
)

// buffer 内嵌bytes.Buffer，支持连写
type buffer struct {
	*bytes.Buffer
}

// CamelToCase 驼峰式写法转为下划线写法
func CamelToCase(name string) string {
	buff := &buffer{Buffer: new(bytes.Buffer)}
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				_, _ = buff.WriteString("_")
			}
			_, _ = buff.WriteRune(unicode.ToLower(r))
		} else {
			_, _ = buff.WriteRune(r)
		}
	}
	return buff.String()
}

// CaseToCamel 下划线写法转为驼峰写法
func CaseToCamel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = cases.Title(language.Und).String(name)
	return strings.Replace(name, " ", "", -1)
}

// Uuid 由 32 个十六进制数字组成，以 6 个组显示，由连字符 - 分隔
func Uuid() string {
	unix32bits := uint32(time.Now().UTC().Unix())
	buff := make([]byte, 12)
	numRead, err := rand.Read(buff)
	if numRead != len(buff) || err != nil {
		return ""
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x-%x", unix32bits, buff[0:2], buff[2:4], buff[4:6], buff[6:8], buff[8:])
}

// RightTrim 移除字符串右侧的空白字符或其他预定义字符。
func RightTrim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimRightFunc(str, unicode.IsSpace)
	}
	return strings.TrimRight(str, characterMask[0])
}

// LeftTrim 移除字符串左侧的空白字符或其他预定义字符。
func LeftTrim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimLeftFunc(str, unicode.IsSpace)
	}
	return strings.TrimLeft(str, characterMask[0])
}

// Replace 字符串替换
func Replace(str string, from []string, to string) string {
	for _, val := range from {
		str = strings.Replace(str, val, to, -1)
	}
	return str
}
func PadLeft(input string, padLength int, padString string) string {
	output := padString
	for padLength > len(output) {
		output += output
	}
	if len(input) >= padLength {
		return input
	}
	return output[:padLength-len(input)] + input
}
func PadRight(input string, padLength int, padString string) string {
	output := padString
	for padLength > len(output) {
		output += output
	}
	if len(input) >= padLength {
		return input
	}
	return input + output[:padLength-len(input)]
}
