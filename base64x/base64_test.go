package base64

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecode(t *testing.T) {
	expect := "hello world"
	actual := Decode("aGVsbG8gd29ybGQ=")
	assert.Equal(t, expect, actual)
}
func TestEncode(t *testing.T) {
	expect := "aGVsbG8gd29ybGQ="
	actual := Encode("hello world")
	assert.Equal(t, expect, actual)
}
func TestImageUrlToBase64(t *testing.T) {
	t.Log(ImageUrlToBase64("https://lmg.jj20.com/up/allimg/1113/052420110515/200524110515-2-1200.jpg"))
}
