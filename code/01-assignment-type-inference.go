package main

import (
	"fmt"
)

// START OMIT
func main() {
	// Method 1
	var foo string

	foo = "Hello, World!"

	// Method 2
	bar := foo

	fmt.Println(foo)
	fmt.Println(bar)
}

// END OMIT
