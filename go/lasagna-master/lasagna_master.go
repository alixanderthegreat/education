package lasagna

// PreparationTime returns the product len of layers and time per layer;
//
// default: when time parameter arguement is 0, time = 2
func PreparationTime(layers []string, time int) int {
	return 0
}

// Quantities returns quantities of noodles & sauce needed
func Quantities(layers []string) (int, float64) {
	return 0, 0.0
}

// AddSecretIngredient take uncommon item from preferences and modifies prearainged
func AddSecretIngredient(preferences, prearainged []string) {
}

// ScaleRecipe scales quantities to the number of portions
func ScaleRecipe(quantities []float64, portions int) []float64 {
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
