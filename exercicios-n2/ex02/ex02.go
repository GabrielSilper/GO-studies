package main

import "fmt"

func main() {
	x := 45 < 56
	y := 50 == (45 + 5)
	z := 33 >= 45

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
