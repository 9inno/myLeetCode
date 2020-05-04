package leetcode

import (
	"math/rand"
	"testing"
)

func Test(t *testing.T) {

	for i := 0 ; i < 10; i ++ {
		number := rand.Intn(100)
		var tmpSlice  []int
		for j := 0 ; j < number; j ++ {
			tmpSlice = append(tmpSlice, rand.Intn(50))
			result := Jump(tmpSlice)
			result1 := Jump1(tmpSlice)
		  	if result != result1{
				t.Errorf("jump got %d , jump1 got %d, with slice %v ", result, result1, tmpSlice)
			}
		}
	}
}