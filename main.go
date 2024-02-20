package main

import (
	"fmt"
	cards "go-studies/card_tricks"
)

func main() {
	slice := []int{3, 2, 6, 4, 8}
	newCards := cards.RemoveItem(slice, 2)
	fmt.Println(newCards)
}
