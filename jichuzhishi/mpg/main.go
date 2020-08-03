package main

func start(c chan int) {
	c <- 100
}

func main() {
	c := make(chan int)
	go start(c)
	<-c
}