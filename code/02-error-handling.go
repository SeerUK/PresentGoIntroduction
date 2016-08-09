package main

import (
	"errors"
	"fmt"
)

// START OMIT
func main() {
	// notVeryUsefulFunction returns an error if you ask for one.
	err := notVeryUsefulFunction(true)
	if err != nil {
		// Perhaps we should try not making an error...
		fmt.Println("Oops, lets just pretend that never happened...")
	}

	err = notVeryUsefulFunction(false)
	if err != nil {
		panic(err)
	}
}

// END OMIT

func notVeryUsefulFunction(shouldError bool) error {
	if shouldError {
		return errors.New("Well, you asked for it...")
	}

	return nil
}
