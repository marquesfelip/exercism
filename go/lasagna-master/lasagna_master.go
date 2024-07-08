package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return len(layers) * 2
	}
	return len(layers) * time
}

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList, ownList []string) {
	ownList[len(ownList) -  1] = friendsList[len(friendsList) - 1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	var amounts []float64

	for i := range quantities {
		amounts = append(amounts, (quantities[i] / 2) * float64(portions))
	}

	return amounts
}
