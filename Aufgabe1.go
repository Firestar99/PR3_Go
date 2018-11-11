/*
1.)
Das nachfolgende Programm verursacht einen Deadlock.
a.)Erklären sie warum es zu diesem Fehler kommt.
b.)Korrigieren sie das Programm indem sie den Wert 42 in einer weiteren Go-Routine dem Channel übergeben.
c.)Erläutern sie den Nutzen von Channels bei der Arbeit mit Go-Routinen.
*/
package main

func main() {

	c := make(chan int)
	c <- 42
	val := <-c
	println(val)
}
