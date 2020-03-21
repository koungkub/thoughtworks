package hand

import (
	"testing"
)

func TestValidatePredictorFailure(t *testing.T) {

	tt := []struct {
		pattern string
	}{
		{""},
		{"C"},
		{"O"},
		{"OO"},
		{"CO"},
		{"CC"},
		{"CC5"},
		{"OO6"},
		{"@#$"},
		{"chicken"},
		{"s3cret"},
	}

	for _, tc := range tt {
		err := ValidatePredictor(tc.pattern)
		if err == nil {
			t.Errorf("%v should error", tc.pattern)
		}
	}
}

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
			t.Errorf("%v should not error; error message %v", tc.pattern, err.Error())
		}
	}
}
