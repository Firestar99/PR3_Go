package main

//b.)
func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	val := <-c
	println(val)
}
