package requestx

import "net/url"

// UrlEncode 编码
func UrlEncode(s string) string {
	escape := url.QueryEscape(s)
	return escape
}

// UrlDecode 解码
func UrlDecode(s string) string {
	unescape, _ := url.QueryUnescape(s)
	return unescape
}

// ParseQuery 获取URL参数
func ParseQuery(s string) map[string][]string {
	u, err := url.Parse(s)
	if err != nil {
		return nil
	}
	urlParam := u.RawQuery
	m, _ := url.ParseQuery(urlParam)
	return m
}
