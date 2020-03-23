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

func ToggleState(state bool) bool {
	return !state
}

// Compare : compare hand and check predictor is predic success or failure
func Compare(patterns ...string) error {

	var openHand, predicNumber int
	for _, pattern := range patterns {
		openHand += countOpen(pattern)
		if err := validatePredictor(pattern); err == nil {
			predicNumber, _ = strconv.Atoi(string(pattern[2]))
		}
	}

	if openHand != predicNumber {
		return errors.New("prediction is invalid")
	}

	return nil
}

func countOpen(patterns ...string) int {

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

	if match, _ := regexp.MatchString(`^[CO][CO]$`, pattern); match {
		return nil
	}

	if match, _ := regexp.MatchString(`^[CO][CO].$`, pattern); match {
		return errors.New("Bad input: no prediction expected, you are not the predictor")
	}

	return errors.New("Bad input: correct input should be the form __ (ex: CC, CO) first two letters indicate [O]pen or [C]losed state for each hand")
}

func validatePredictor(pattern string) error {

	if match, _ := regexp.MatchString(`^[CO][CO][0-4]$`, pattern); match {
		return nil
	}

	if match, _ := regexp.MatchString(`^[CO][CO].$`, pattern); match {
		return errors.New("Bad input: prediction should be in the range of 0-4")
	}

	if match, _ := regexp.MatchString(`^..[0-4]$`, pattern); match {
		return errors.New("Bad input: first two charactor should be [O]pen or [C]losed")
	}

	return errors.New("Bad input: correct input should be the form ___ (ex: CC1, CO3) first two letters indicate [O]pen or [C]losed state for each hand followed by the prediction (0-4)")
}
