package main

func write(ch chan int, i int) {
	ch <- i
}

func main() {
	c := make(chan int)
	go write(c, 42)
	val := <-c
	println(val)
}
