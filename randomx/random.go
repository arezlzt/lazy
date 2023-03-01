package randomx

import (
	"crypto/rand"
	"github.com/arezlzt/lazy/arrayx"
	"github.com/arezlzt/lazy/structx"
	"golang.org/x/exp/constraints"
	"math"
	"math/big"
	rand2 "math/rand"
	"strconv"
	"strings"
	"time"
)

type Provability[T comparable] struct {
	Data        T   // 数据
	Provability int // 随机概率数值
}

func RandomProvability[T comparable](randList []Provability[T]) T {
	if 100 != arrayx.Sum(structx.StructColumn[Provability[T], int](randList, "provability")) {
		panic("the provability sum not equal to 100")
	}
	number, _ := rand.Int(rand.Reader, big.NewInt(100))
	r := int(number.Uint64())
	var result T
	left := 0
	for _, val := range randList {
		right := left + val.Provability
		if r >= left && r < right {
			result = val.Data
			break
		}
		left = right
	}
	return result
}

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	letterRunes := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	r := rand2.New(rand2.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	return string(b)
}

// RealRandFloat64
// 生成范围随机float64
func RealRandFloat64(min, max float64) float64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	minStr := strconv.FormatFloat(min, 'f', -1, 64)
	decimalLength := 0
	if strings.Index(minStr, ".") > -1 {
		decimalLength = strings.Index(minStr, ".")
	}
	multipleNum := len(minStr) - (decimalLength + 1)
	multiple := math.Pow10(multipleNum)
	minMultiple := min * multiple
	maxMultiple := max * multiple
	randVal := RealRandInteger[int64](int64(minMultiple), int64(maxMultiple))
	result := float64(randVal) / multiple
	return result
}

// RealRandInteger
// 真随机数
func RealRandInteger[T constraints.Integer](min, max T) T {
	if min >= max || min < 0 || max < 0 {
		return max
	}
	number, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	return T(number.Int64()) + min
}
