package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Message struct {
	str string
	wait chan bool
}

func main1() {
	joe := boring("Joe")
	ann := boring("ann")
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(<-joe)
	// 	fmt.Println(<-ann)
	// }

	c := fanIn(joe, ann)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("You are boring, Im leaving")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ;i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { for {c <- <- input1} }()
	go func() { for {c <- <- input2} }()
	return c
}

func newFanIn(input1, input2 <-chan string) <- chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <- input1: c <- s
			case s := <- input2: c <- s
			}
		}
	}()
	return c
}

func main2() {
	c := boring("Joe")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <- c:
			fmt.Println(s)
		case <- timeout:
			fmt.Println("You talked too much")
			return
		}
	}
}

func main() {
	const n = 100000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}

func f(left, right chan int) {
	left <- 1 + <- right
}
