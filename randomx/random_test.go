package randomx

import (
	"fmt"
	"testing"
)

func TestRandProvability(t *testing.T) {
	list := make([]Provability[string], 0)
	list = append(list, Provability[string]{
		Provability: 10,
		Data:        "iphone",
	})
	list = append(list, Provability[string]{
		Provability: 20,
		Data:        "ipad",
	})
	list = append(list, Provability[string]{
		Provability: 70,
		Data:        "other",
	})
	for i := 0; i < 100; i++ {
		data := RandomProvability[string](list)
		fmt.Println(data)
	}

}
func Test_RandStringRunes(t *testing.T) {
	t.Log(RandStringRunes(100))
}
func Test_RealRandFloat64(t *testing.T) {
	t.Log(RealRandFloat64(0.101, 1.01))
}
func Test_RealRandInteger(t *testing.T) {
	t.Log(RealRandInteger[int16](10, 20))
}
