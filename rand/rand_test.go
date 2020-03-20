package rand

import "testing"

func TestRandomNumber(t *testing.T) {

	tt := []struct {
		max int
		min int
	}{
		{3, 0},
		{5, 2},
	}

	for _, tc := range tt {
		number := RandomNumber(tc.max, tc.min)
		if number < tc.min || number >= tc.max {
			t.Errorf("%v should be in range [%v, %v]", number, tc.min, tc.max)
		}
	}
}
