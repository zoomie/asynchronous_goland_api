package main

import "fmt"

type Redis struct {
	foo string
	point_num *int
}

func (r *Redis) ChangeValue(value *int) {
	r.point_num = value
}

func main() {
	// channel := make(chan string)
	num1 := 10
	num2 := 20
	redis := Redis{"memory", &num1}
	fmt.Println(redis)
	redis.ChangeValue(&num2)
	fmt.Println(redis)
	fmt.Println(*redis.point_num)
	// redis := Redis{"test", }
	// redis.Store(4)
	// fmt.Println(redis)
}