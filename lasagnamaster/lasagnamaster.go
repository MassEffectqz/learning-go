package lasagnamaster

func PreparationTime(layers []string, time int) int {
	if time > 0 {
		return len(layers) * time
	}
	return len(layers) * 2
}
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, i := range layers {
		if i == "sauce" {
			sauce += 0.2
		}
		if i == "noodles" {
			noodles += 50
		}
	}
	return noodles, sauce
}
func AddSecretIngredient(friendList, myList []string) {
    secret := friendList[len(friendList)-1]
    myList[len(myList)-1] = secret
}
func ScaleRecipe(quantities []float64, num int) []float64 {
	scaled := make([]float64, len(quantities))
	for i := range quantities {
		scaled[i] = quantities[i] * float64(num) / 2.0
	}
	return scaled
}
