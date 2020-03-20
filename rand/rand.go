package rand

import (
	"fmt"
	"math/rand"
)

const (
	// Max : max number can be expected
	Max = 5

	// Min : min number can be expected
	Min = 0
)

// RandomHands : random hand if predictor true return `%s%s%d` pattern
// if predictor is false return `%s%s` pattern
func RandomHands(isPredictor bool) string {

	if isPredictor {
		return fmt.Sprintf("%v%v%v", RandomHand(RandomNumber(Max, Min)), RandomHand(RandomNumber(Max, Min)), RandomNumber(Max, Min))
	}

	return fmt.Sprintf("%v%v", RandomHand(RandomNumber(Max, Min)), RandomHand(RandomNumber(Max, Min)))
}

// RandomHand : random hand as 'O' is open, 'C' is close
func RandomHand(seed int) string {
	return []string{"O", "C", "O", "C"}[seed]
}

// RandomNumber : random number in range [min, max]
func RandomNumber(max, min int) int {
	return min + rand.Intn(max-min)
}
