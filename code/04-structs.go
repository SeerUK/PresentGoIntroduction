package main

import "fmt"

// START OMIT
type Greeter struct {
	Name string
}

// Capitalisation of names is important...
func (g *Greeter) Greet() string {
	return fmt.Sprintf("Hello, %s!", g.Name)
}

func main() {
	greeter := Greeter{
		Name: "Byng",
	}

	fmt.Println(greeter.Greet())
}

// END OMIT
