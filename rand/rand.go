package rand

import "math/rand"

const (
	// Max : max number can be expected
	Max = 5

	// Min : min number can be expected
	Min = 0
)

// RandomHand : random hand as 'O' is open, 'C' is close
func RandomHand(seed int) string {
	return []string{"O", "C"}[seed]
}

// RandomNumber : random number in range [min, max]
func RandomNumber(max, min int) int {
	return min + rand.Intn(max-min)
}
