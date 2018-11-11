package main

import "fmt"

func main() {
	messageChannel := make(chan string)

	go func() {
		messageChannel <- "42"
	}()

	fmt.Print(<-messageChannel)
}
