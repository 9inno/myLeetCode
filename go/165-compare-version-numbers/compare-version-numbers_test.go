package leetcode

import "testing"

func TestCompareVersion(t *testing.T) {
	tests := []struct{
		string1 string
		string2 string
		result int
	}{
		{"0.1", "1.1", -1},
		{"7.5.2.4", "7.5.3", -1},
		{"1.01", "1.0001", 0},
		{"1.2.1", "1.2", 1},
	}

	for _ ,test :=range tests {
		if actual := CompareVersion(test.string1, test.string2); actual != test.result {
			t.Errorf("got %d , expected %d", actual, test.result)
		}
	}
}