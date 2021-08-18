package main 

import "fmt"

func main() {
	c := make(chan int)
    //sending
	go func() {
		c <- 1
	} ()  //lambda function
    //sniff
	val := <-c     
    fmt.Println(val)
    //sending
	go func() {
		c <- 2
	} ()
    //sniff
	val = <-c
	fmt.Println(val)
}