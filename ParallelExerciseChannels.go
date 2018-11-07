package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(countDeviders(123456780 + i))
	}
}

func countDeviders(number int) string {
	var count = 0
	for i := 1; i < number; i++ {
		if number%i == 0 {
			count++
		}
	}
	return strconv.Itoa(number) + ": " + strconv.Itoa(count)
}
