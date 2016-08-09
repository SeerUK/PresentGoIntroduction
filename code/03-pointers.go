package main

import "fmt"

// START OMIT
func main() {
	n1, n2 := 42, 84

	p1 := &n1           // Create a pointer to n1
	fmt.Println(1, p1)  // Print the pointer
	fmt.Println(2, *p1) // Print the underlying value (dereference)

	*p1 = 21           // Set n1 through the pointer
	fmt.Println(3, n1) // Print the new value of n1

	p2 := n2
	p2 = 42
	fmt.Println(4, p2) // What will these be?
	fmt.Println(5, n2) // ...
}

// END OMIT
