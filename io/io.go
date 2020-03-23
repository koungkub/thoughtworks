package io

import (
	"bufio"
	"os"
	"strings"
)

func Welcome() string {
	return "Welcome to the game !!"
}

func Turn(predictor bool) string {

	if predictor {
		return "You are the predictor, what is your input ?"
	}
	return "AI is the predictor, what is your input ?"
}

func Scan() string {

	input := bufio.NewReader(os.Stdin)

	txt, _ := input.ReadString('\n')
	txt = strings.Replace(txt, "\n", "", -1)

	return txt
}
