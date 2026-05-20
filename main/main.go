package main

import (
	"fmt"
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
	fmt.Println(cards.GetItem([]int{1,2,3,5}, 4))
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println(cards.SetItem([]int{1,2,3,4,5}, 2, 4))
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println(cards.PrependItems([]int{3, 2, 6, 4, 8}))
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println(cards.RemoveItem([]int{3,2,6,4,8}, 11))
	fmt.Println(strings.Repeat("=", 15))
}