package main

import (
	"fmt"
	"go-studies/interest"
)

func main() {
	n := interest.YearsBeforeDesiredBalance(200.75, 214.88)
	fmt.Println(n)
}
