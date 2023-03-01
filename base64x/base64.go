package base64

import (
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// Encode base64编码
func Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

// Decode base64解码
func Decode(input string) string {
	decodeString, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return ""
	}
	return string(decodeString)
}

// ImageUrlToBase64 图片Url转base64
func ImageUrlToBase64(url string) (string, error) {
	if url == "" {
		return "", errors.New("url is empty")
	}
	data, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(data.Body)
	result, _ := ioutil.ReadAll(data.Body)
	base64Data := base64.StdEncoding.EncodeToString(result)
	return base64Data, nil
}
