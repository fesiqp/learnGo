package main

import (
	"fmt"
)

func main() {
	x := 0

	// It's an anonymous functions, attributed to a variable.
	// It's one of the ways to declare a function inside another function
	// Assigning a function to variable is called function expression, and
	// the function must be anonymous
	increment := func() int {
		x++
		return x

	}
	fmt.Println("First call:", increment())
	fmt.Println("Second call", increment())
}
