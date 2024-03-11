package lasagna

const (
	defaultAveragePreparationTime = 2
	noodleQtyPerLayer             = 50
	sauceQtyPerLayer              = 0.2
	defaultPortions               = 2
)

// PreparationTime estimates the time it takes to prepare lasagna, given
// the layers and the average preparation time for a layer.
func PreparationTime(layers []string, avgPreparationTime int) int {
	if avgPreparationTime == 0 {
		avgPreparationTime = defaultAveragePreparationTime
	}
	return len(layers) * avgPreparationTime
}

// Quantities computes the amounts of noodles and sauce needed.
func Quantities(layers []string) (noodles int, sauce float64) {
	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			noodles += noodleQtyPerLayer
		case "sauce":
			sauce += sauceQtyPerLayer
		}
	}
	return
}

// AddSecretIngredient adds the secret ingredient from one recipe to the other
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe scales the proportions of the recipe for a larger number of portions
func ScaleRecipe(quantities []float64, portions int) (newQuantities []float64) {
	for i := 0; i < len(quantities); i++ {
		newQuantities = append(newQuantities, quantities[i]*float64(portions)/defaultPortions)
	}
	return
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
