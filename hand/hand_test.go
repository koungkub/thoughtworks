package hand

import (
	"testing"
)

func TestCountOpen(t *testing.T) {

	tt := []struct {
		predictor    string
		notPredictor string
		totalOpen    int
	}{
		{"OO2", "OO", 4},
		{"CC1", "CC", 0},
		{"CO", "CC", 1},
		{"OC", "OC", 2},
	}

	for _, tc := range tt {
		total := CountOpen(tc.predictor, tc.notPredictor)
		if total != tc.totalOpen {
			t.Errorf("%v and %v should be %v", tc.predictor, tc.notPredictor, tc.totalOpen)
		}
	}
}

func TestValidate(t *testing.T) {

	tt := []struct {
		pattern     string
		isPredictor bool
	}{
		{"CO2", true},
		{"CO1", true},
		{"CO", false},
		{"", false},
		{"CC", false},
		{"CAC", false},
	}

	for _, tc := range tt {
		err := Validate(tc.isPredictor, tc.pattern)
		if err != nil {
			if tc.isPredictor {
				t.Errorf("%v should error; predictor status is %v", tc.pattern, tc.isPredictor)
			}
		}
	}
}

func TestValidateNotPredictorFailure(t *testing.T) {

	tt := []struct {
		pattern string
	}{
		{"OO1"},
		{"CO2"},
		{"OC3"},
		{"CC3"},
		{""},
		{"HEN"},
		{"12"},
	}

	for _, tc := range tt {
		err := ValidateNotPredictor(tc.pattern)
		if err == nil {
			t.Errorf("%v should error", tc.pattern)
		}
	}
}

func TestValidateNotPredictorSuccess(t *testing.T) {

	tt := []struct {
		pattern string
	}{
		{"OO"},
		{"CO"},
		{"OC"},
		{"CC"},
	}

	for _, tc := range tt {
		err := ValidateNotPredictor(tc.pattern)
		if err != nil {
			t.Errorf("%v should error; errror message %v", tc.pattern, err.Error())
		}
	}
}

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
