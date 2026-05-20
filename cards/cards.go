package cards

func FavoriteCards() []int {
	return []int{2,6,9}
}

func GetItem(slice []int, index int) int {
	if index < 0 || index >= len(slice) { 
		return -1
	}
	return slice[index]
}

func SetItem(slice []int, index, value int) []int { 
	if index < 0 || index >= len(slice) { 
		slice = append(slice, value)
		return slice
	}
	slice[index] = value 
	return slice
}	

func PrependItems(slice []int, values ...int) []int {
	slice = append(values, slice...)
	return slice
}

func RemoveItem(slice []int, index int) []int {
	if index < 0 || index > len(slice) {
		return slice
	}
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}
