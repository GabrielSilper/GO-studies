package main

import "fmt"

func main() {
	rect := Rect{Width: 5, Height: 2}
	rect.SetHeight(3)
	fmt.Println(rect.ShowSize())
}

type Rect struct {
	Width, Height int
}

func (rect Rect) ShowSize() int {
	return rect.Height * rect.Width
}

func (rect *Rect) SetHeight(height int) {
	rect.Height = height
}
