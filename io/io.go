package io

import (
	"bufio"
	"fmt"
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

func Scan(pattern string) string {

	input := bufio.NewReader(os.Stdin)

	fmt.Print(pattern)
	txt, _ := input.ReadString('\n')
	txt = strings.Replace(txt, "\n", "", -1)

	return strings.ToUpper(txt)
}
