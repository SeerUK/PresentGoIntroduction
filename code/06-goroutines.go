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

	go delayedReturn("World", 2000, ch)
	go delayedReturn("Hello...", 200, ch)

	fmt.Println("Waiting for return...")

	fmt.Print(<-ch)
	fmt.Println(<-ch)
}

// END OMIT
