package aesx

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	key = "bGcGfWb3Kg2s4gcG"
	iv  = "aebksHkG4jAEk2Ag"
)

func TestEncrypt(t *testing.T) {
	expect := "DIVKKtNuw62mZx4GO3mdHQ=="
	actual, err := Encrypt("hello world", []byte(key), iv)
	if err != nil {
		assert.Error(t, err)
	} else {
		assert.Equal(t, expect, actual)
	}

}
func TestDecrypt(t *testing.T) {
	expect := "hello world"
	actual, err := Decrypt("DIVKKtNuw62mZx4GO3mdHQ==", []byte(key), iv)
	if err != nil {
		assert.Error(t, err)
	} else {
		assert.Equal(t, expect, actual)
	}
	time.Now()
}
