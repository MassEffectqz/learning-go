package luhn

import "unicode"

func Valid(id string) bool {
	var digits []int
	for _, r := range id {
		if r == ' ' {
			continue
		}
		if !unicode.IsDigit(r) {
			return false
		}
		digits = append(digits, int(r-'0'))
	}
	if len(digits) < 2 {
		return false
	}
	sum := 0
	doubled := false
	for i := len(digits) - 1; i >= 0; i-- {
		d := digits[i]
		if doubled {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		doubled = !doubled
	}
	return sum%10 == 0
}
