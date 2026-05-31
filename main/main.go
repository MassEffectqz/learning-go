package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
	"strings"
)

const (
	ADMIN = iota
	USER
	MANAGER
	VIEWER
)

func main() {
	fmt.Println(strings.Repeat("=", 15))
	// PrintToMain()
	fmt.Println(strings.Repeat("=", 15))
	fmt.Println(ADMIN, USER, MANAGER, VIEWER)
	fmt.Println(strings.Repeat("=", 15))
	random := rand.IntN(100-50) + 50
	fmt.Println(random)
	fmt.Println(strings.Repeat("=", 15))
	a := []int{1, 2, 3, 4, 5, 6}
	a = slices.DeleteFunc(a, func(x int) bool {
		return x%2 == 0
	})
	for _, v := range a {
		fmt.Println(v)
	}
	word := "Приветики"
	runeWords := []rune(word)
	for i := range runeWords { 
		fmt.Printf("%c\n", runeWords[i])
	}
	text := bufio.NewScanner(os.Stdin)

	fmt.Println(text)
	var c int = 5
	var d *int 
	d = &c
	fmt.Println(*d)
	
}
