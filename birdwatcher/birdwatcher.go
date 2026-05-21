package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
	var value int
	for i := 0; i < len(birdsPerDay); i++ {
		value += birdsPerDay[i]
	}
	return value
}
func BirdsInWeek(birdsPerDay []int, week int) int {
	startIndex := (week - 1) * 7
	endIndex := startIndex + 7
	weeklyBirds := birdsPerDay[startIndex:endIndex]
	total := 0
	for _, count := range weeklyBirds {
		total += count
	}
	return total
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}
	return birdsPerDay
}
