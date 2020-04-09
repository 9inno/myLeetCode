package leetcode

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct{
		n int
		result []string
	}{
		{3, []string{"((()))", "(()())", "(())()","()(())","()()()"},},
	}

	for _ , test := range tests {
		if actual := GenerateParenthesis(test.n); !reflect.DeepEqual(actual, test.result) {
			t.Errorf( "got %v , expected %v", actual, test.result)
		}
	}
}