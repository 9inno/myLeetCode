package leetcode

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct{
		s []byte
		result [] byte
	} {
		{[]byte ("hello"), [] byte ("olleh")},
		{[]byte ("Hannah"), [] byte ("hannaH")},
	}

	for _ , test := range tests {
		if  ReverseString(test.s) ; !reflect.DeepEqual(test.s, test.result) {
			t.Errorf("got %v, excpeted %v", test.s, test.result)
		}
	}
}

func TestReverseString1(t *testing.T) {
	tests := []struct{
		s []byte
		result [] byte
	} {
		{[]byte ("hello"), [] byte ("olleh")},
		{[]byte ("Hannah"), [] byte ("hannaH")},
	}

	for _ , test := range tests {
		if  ReverseString1(test.s) ; !reflect.DeepEqual(test.s, test.result) {
			t.Errorf("got %v, expected %v", test.s, test.result)
		}
	}
}