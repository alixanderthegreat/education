package chance

import (
	"fmt"
	"math/rand"
)

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return 1 + rand.Intn(20)
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	// define iterator at a higher scope
	var i int
	for {
		i++ // iterate over the iterator
		// produce a deterministically weighted random float
		var ( // https://pkg.go.dev/math/rand#NormFloat64
			random_float              = rand.NormFloat64()
			desired_standard_devation = 3.6
			desired_mean              = 4.2
			energy                    = random_float * desired_standard_devation * desired_mean
		)
		if energy > 0 && energy < 12.0 {
			fmt.Println("iteration: ", i, " energy:", energy)
			return energy
		}

	}
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	panic("Please implement the ShuffleAnimals function")
}
