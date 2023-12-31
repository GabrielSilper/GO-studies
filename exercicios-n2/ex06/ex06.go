package main

import (
	"fmt"
)

func main() {
	const (
		_ = 2023 + iota
		x1
		x2
		x3
		x4
	)

	fmt.Println(x1, x2, x3, x4)
}
