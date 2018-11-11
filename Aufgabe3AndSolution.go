/*
3.)
Schreiben sie ein Programm das die Summe der Quadrate der einzelnen Ziffern der Zahl 12345678 parallel berechnet und diese ausgibt.

*/

//Solution
package main

import (
	"fmt"
)

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		go func() {
			digit := number % 10
			sum += digit * digit
			number /= 10
		}()
	}
	squareop <- sum
}

func main() {
	number := 123456789
	sqrch := make(chan int)
	go calcSquares(number, sqrch)
	squares := <-sqrch
	fmt.Println("Final output", squares)
}
