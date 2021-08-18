package main

import "fmt"

func main() {
	c := make(chan int, 3) //buffer channel of length 3

	go func() {
		c <- 0
		c <- 1
		c <- 2
		c <- 3
		c <- 4
		close(c)
	} ()

	for i := range c {
		fmt.Println(i)
	}
}