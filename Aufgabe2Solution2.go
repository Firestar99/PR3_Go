package main

import (
	"fmt"
)

func main() {
	done := make(chan bool)
	go func() {
		done <- true
	}()
	<-done
	fmt.Println("done")
}
