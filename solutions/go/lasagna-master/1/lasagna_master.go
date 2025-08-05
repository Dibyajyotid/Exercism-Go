package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) int {
    if avgPrepTime == 0 {
        return int(len(layers) * 2)
    }

    return int(len(layers) * avgPrepTime)
}

// TODO: define the 'Quantities()' function
func Quantities( layers []string) (noodles int, sauce float64) {
    for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, ownList []string) {
    if len(friendList) > 0 && len(ownList) > 0 {
        ownList[len(ownList) - 1] = friendList[len(friendList) - 1]
    }
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int) []float64 {
    scaledQuantities := []float64{}
    factor := float64(portion) / 2.0

    for _, q := range quantities {
        scaledQuantities = append(scaledQuantities, q * factor)
    }

    return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
