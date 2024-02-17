package main

import (
	"fmt"
	"go-studies/techpalace"
)

func main() {
	message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`

	fmt.Println(techpalace.CleanupMessage(message))
}
