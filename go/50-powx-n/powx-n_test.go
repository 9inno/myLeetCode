package leetcode

import (
	"math"
	"math/rand"
	"testing"
)

func TestMyPow(t *testing.T) {

	for i := 0 ; i < 100; i ++ {
		x := rand.Float64()
		n := rand.Int()
		result := MyPow(x, n)
		actual := math.Pow(x, float64(n))
		if result != actual {
			t.Errorf("x %f, n %d , got %f , expected %f", x, n, result, actual)
		}
	}
}
