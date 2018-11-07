package main

import "fmt"

func main() {
	messageChannel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			messageChannel <- i
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Print(<-messageChannel)
	}

	fmt.Println()
	fmt.Println("Done!")
}
