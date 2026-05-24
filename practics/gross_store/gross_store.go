package gross

func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}
func NewBill() map[string]int {
	return make(map[string]int)
}
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
	if !exists {
		return false
	}
	bill[item] += value
	return true
}
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	current, exists := bill[item]
	if !exists {
		return false
	}
	unitValue, unitExist := units[unit]
	if !unitExist {
		return false
	}
	newQty := current - unitValue
	if newQty <= 0 {
		delete(bill, item)
	} else {
		bill[item] = newQty
	}
	return true
}
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, exists := bill[item]
	return qty, exists
}
