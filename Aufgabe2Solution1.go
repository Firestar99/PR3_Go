package main

import (
	"fmt"
)

func doSomething(done chan bool) {
	done <- true
}

func main() {
	done := make(chan bool)
	go doSomething(done)
	<-done
	fmt.Println("done")
}
