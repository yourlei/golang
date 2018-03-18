package main

import (
	"fmt"
	"chapter1/shape"
)

func main() {
	react := shape.React{Width: 10, Height: 8}
	reactArea := react.Area()

	fmt.Printf("the react area is %d", reactArea)
}