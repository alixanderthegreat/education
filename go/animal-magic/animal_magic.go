package chance

import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return 1 + rand.Intn(20)
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	for {
		// https://pkg.go.dev/math/rand#NormFloat64
		var desiredStdDev, desiredMean, r float64 = 3.6, 4.2, rand.Float64()
		r = r * desiredStdDev * desiredMean
		if r > 0 && r < 12.0 {
			return r
		}
	}
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	panic("Please implement the ShuffleAnimals function")
}
