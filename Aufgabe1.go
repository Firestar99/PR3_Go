/*
1.)
Das nachfolgende Programm verursacht einen Deadlock.
Erklären sie warum es zu diesem Fehler kommt und korrigieren sie das Programm indem sie den Wert 42 in einer weiteren Go-Routine dem Channel übergeben.
*/
package main

func main() {
	c := make(chan int)
	c <- 42
	val := <-c
	println(val)
}
