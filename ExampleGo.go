package main

import (
	"fmt"
	"time"
)

func main() {
	/*go*/ printHellos()

	for i := 0; i < 5; i++ {
		fmt.Println("Done!")
		time.Sleep(1000 * time.Millisecond)
	}
}

func printHellos() {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello!")
		time.Sleep(1000 * time.Millisecond)
	}
}
