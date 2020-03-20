package hand

import (
	"testing"
)

func TestValidatePredictorSuccess(t *testing.T) {

	tt := []struct {
		pattern string
	}{
		{"OO1"},
		{"CO2"},
		{"OC3"},
		{"CC4"},
		{"CO4"},
	}

	for _, tc := range tt {
		err := ValidatePredictor(tc.pattern)
		if err != nil {
			t.Errorf("%v should be error; error message %v", tc.pattern, err.Error())
		}
	}
}
