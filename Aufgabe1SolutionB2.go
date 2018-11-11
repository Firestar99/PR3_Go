package main

func write42(ch chan int) {
	ch <- 42
}

func main() {
	c := make(chan int)
	go write42(c)
	val := <-c
	println(val)
}
