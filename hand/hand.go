package hand

import (
	"errors"
	"regexp"
)

// Validate : validate pattern base on isPredictor flag
func Validate(isPredictor bool, pattern string) error {

	if isPredictor {
		if err := ValidatePredictor(pattern); err != nil {
			return err
		}
		return nil
	}

	if err := ValidateNotPredictor(pattern); err != nil {
		return err
	}
	return nil
}

// ValidateNotPredictor : validate not predictor input
func ValidateNotPredictor(pattern string) error {

	validate := regexp.MustCompile(`^[CO][CO]$`)

	if match := validate.MatchString(pattern); !match {
		return errors.New("invalid input")
	}

	return nil
}

// ValidatePredictor : validate predictor input
func ValidatePredictor(pattern string) error {

	validate := regexp.MustCompile(`^[CO][CO][0-4]$`)

	if match := validate.MatchString(pattern); !match {
		return errors.New("invalid input")
	}

	return nil
}
