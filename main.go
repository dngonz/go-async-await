package main

import (
	"log"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		log.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func asyncFunction(s int32) <-chan int32 {
	r := make(chan int32)
	go func() {
		defer close(r)
		time.Sleep(time.Second * 2)
		r <- s
	}()
	return r
}

func main() {
	// say
	go say("world")
	say("hello")

	// sum
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	log.Println(x, y, x+y)

	// asyncFunction
	firstChannel, secondChannel := asyncFunction(2), asyncFunction(3)
	first, second := <-firstChannel, <-secondChannel
	log.Println(first, second) // 2 & 3
}
