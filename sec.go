package main

import "fmt"

type Redis struct {
	conn string
}

func (r *Redis) Put(input string, c chan int) {
	r.conn = input
	c <- 0
}

func main()  {
	channel := make(chan int)
	r := Redis{"test string"}
	fmt.Println(r)
	go r.Put("new string", channel)
	<-channel
	fmt.Println(r)
	// fmt.Println("test")
}