package main

import (
	"fmt"
	"myapp/birdwatcher"
	"myapp/cards"
	"strings"
)

func main() {
	a := "круто"
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println(a)
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println(cards.FavoriteCards())
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println(cards.GetItem([]int{1, 2, 3, 5}, 4))
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println(cards.SetItem([]int{1, 2, 3, 4, 5}, 2, 4))
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println(cards.PrependItems([]int{3, 2, 6, 4, 8}))
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println(cards.RemoveItem([]int{3, 2, 6, 4, 8}, 11))
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println(birdwatcher.TotalBirdCount([]int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}))
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println(birdwatcher.BirdsInWeek([]int{1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7}, 2))
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println(birdwatcher.FixBirdCountLog([]int{0,1,2,3,4,5}))
}
