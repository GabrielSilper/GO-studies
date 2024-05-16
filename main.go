package main

import "fmt"

func main() {
	var p *int

	a := 45

	p = &a

	fmt.Println(*p)
}
