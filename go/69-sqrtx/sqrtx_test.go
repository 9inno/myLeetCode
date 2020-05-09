package leetcode

import (
	"math"
	"math/rand"
	"testing"
)

func Test(t *testing.T) {

	for i := 0 ; i < 100; i ++ {
		nums := rand.Intn(math.MaxInt32)
		actual := MySqrt(nums)
		actual1 := MySqrt1(nums)

		if actual != actual1 {
			t.Errorf("MySqrt got %d , MySqrt1 got %d , with num %d", actual, actual1, nums)
		}
	}
}
