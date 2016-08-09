package main

import "fmt"

// START OMIT
type Greeter interface {
	Greet(string) string
}

type HelloGreeter struct{}

func (g HelloGreeter) Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	var greeter Greeter

	greeter = HelloGreeter{}

	fmt.Println(greeter.Greet("Byng"))
}

// END OMIT
