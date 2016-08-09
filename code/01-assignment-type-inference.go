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
	bar := "Hello, World!"
	baz := bar
	qux := createGreeting("Gophers")

	fmt.Println(foo)
	fmt.Println(bar)
	fmt.Println(baz)
	fmt.Println(qux)
}

func createGreeting(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// END OMIT
