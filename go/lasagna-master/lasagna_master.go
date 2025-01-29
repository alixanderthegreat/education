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
	var noodles, sauce int
	for _, layer := range layers {
		switch {
		case layer == "noodles":
			noodles++
		case layer == "sauce":
			sauce++
		}
	}
	return noodles * 50, // 50g for each layer of noodles
		float64(sauce) * 0.2 // 0.2ltr for each layer of sauce
}

// AddSecretIngredient take uncommon item from preferences and modifies prearainged
func AddSecretIngredient(preferences, prearainged []string) {
}

// ScaleRecipe scales quantities to the number of portions
func ScaleRecipe(quantities []float64, portions int) []float64 {
	return nil
}
