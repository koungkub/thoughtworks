package rand

import "math/rand"

const (
	// Max : max number can be expected
	Max = 5

	// Min : min number can be expected
	Min = 0
)

// RandomNumber : random number in range [min, max]
func RandomNumber(max, min int) int {
	return min + rand.Intn(max-min)
}
