package main

import "errors"

func main() {

}

func notVeryUsefulFunction(shouldError bool) error {
	if shouldError {
		return errors.New("Well, you asked for it...")
	}

	return nil
}
