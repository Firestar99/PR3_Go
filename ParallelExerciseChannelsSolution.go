package main

import (
	"fmt"
	"strconv"
)

func main() {
	prints := make(chan string)
	for i := 0; i < 5; i++ {
		n := i
		go func() {
			prints <- countDividers(123456780 + n)
		}()
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-prints)
	}
}

func countDividers(number int) string {
	var count = 0
	for i := 1; i < number; i++ {
		if number%i == 0 {
			count++
		}
	}
	return strconv.Itoa(number) + ": " + strconv.Itoa(count)
}
