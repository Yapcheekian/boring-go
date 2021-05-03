package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
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
