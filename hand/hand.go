package hand

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const (

	// Open : is represent 'O'
	Open = `O`
)

// Compare : compare hand and check predictor is predic success or failure
func Compare(predictor, notPredictor string) error {

	ans, err := strconv.Atoi(string(predictor[2]))
	if err != nil {
		return err
	}

	openHand := CountOpen(predictor, notPredictor)
	if openHand != ans {
		return errors.New("predictor is invalid")
	}

	return nil
}

// CountOpen : counter open ('O') in hand
func CountOpen(patterns ...string) int {

	var counter int
	for _, pattern := range patterns {
		counter += strings.Count(pattern, Open)
	}

	return counter
}

// Validate : validate pattern base on isPredictor flag
func Validate(isPredictor bool, pattern string) error {

	if isPredictor {
		if err := validatePredictor(pattern); err != nil {
			return err
		}
		return nil
	}

	if err := validateNotPredictor(pattern); err != nil {
		return err
	}
	return nil
}

func validateNotPredictor(pattern string) error {

	validate := regexp.MustCompile(`^[CO][CO]$`)

	if match := validate.MatchString(pattern); !match {
		return errors.New("invalid input")
	}

	return nil
}

func validatePredictor(pattern string) error {

	validate := regexp.MustCompile(`^[CO][CO][0-4]$`)

	if match := validate.MatchString(pattern); !match {
		return errors.New("invalid input")
	}

	return nil
}
