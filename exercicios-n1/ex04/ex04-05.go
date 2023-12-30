package main

import "fmt"

type meuInt int

var x meuInt
var y int

func main() {
	fmt.Printf("valor de x=%v do tipo %T\n", x, x)

	x = 42
	fmt.Printf("valor de x=%v do tipo %T\n", x, x)

	y = int(x)
	fmt.Printf("valor de y=%v do tipo %T\n", y, y)
}
