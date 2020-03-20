package rand

import "testing"

func TestRandomNumber(t *testing.T) {

	tt := []struct {
		max    int
		min    int
		number int
	}{
		{3, 0, 0},
	}

	for _, tc := range tt {
		number := RandomNumber(tc.max, tc.min)
		if tc.number != number {
			t.Errorf("%v should be in range [%v, %v]", tc.number, tc.min, tc.max)
		}
	}
}
