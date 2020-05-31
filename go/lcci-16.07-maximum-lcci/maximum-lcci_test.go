package leetocde

import "testing"

func TestMaximum(t *testing.T) {
	tests := []struct{
		a , b ,result int
	}{
		{2, 7 , 7},
		{2147483647, -2147483648, 2147483647 },
	}

	for _ , test := range tests {
		if actual := Maximum(test.a, test.b) ; actual != test.result {
			t.Errorf("got %d , expected %d", actual, test.result)
		}
	}
}
