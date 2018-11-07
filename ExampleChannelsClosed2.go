package main

import "fmt"

func main() {
	messageChannel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			messageChannel <- i
		}
		close(messageChannel)
	}()

	var total int
	for {
		number, isOpen := <-messageChannel
		fmt.Print(number)
		total += number

		if !isOpen {
			break
		}
	}

	fmt.Println()
	fmt.Println(total)
	fmt.Println("Done!")
}
