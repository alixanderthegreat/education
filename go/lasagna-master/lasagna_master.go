package lasagna

// PreparationTime returns the product len of layers and time per layer;
//
// default: when time parameter arguement is 0, time = 2
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// Quantities returns quantities of grams of noodles & liters of sauce needed
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
