package main

import (
	"fmt"
	"time"
)

// START OMIT
func delayedReturn(in string, delay time.Duration, ch chan string) {
	time.Sleep(delay * time.Millisecond)

	ch <- in
}

func main() {
	ch := make(chan string)

	defer close(ch)

	go delayedReturn("I'm thinking...", 2000, ch)
	go delayedReturn("Hold on...", 200, ch)
	go delayedReturn("Hmmm...", 400, ch)

	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Println(<-ch)
}

// END OMIT
