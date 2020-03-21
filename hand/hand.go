package hand

import (
	"errors"
	"regexp"
)

const (
	// Open : represent Open as 'O'
	Open = `O`

	// Close : represent Close as 'C'
	Close = `C`
)

// ValidatePredictor : validate predictor input
func ValidatePredictor(pattern string) error {

	validate := regexp.MustCompile(`^[CO][CO][0-4]$`)

	if match := validate.MatchString(pattern); !match {
		return errors.New("invalid input")
	}

	return nil
}
