package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Atoi returns an error value
	// So, you should always check it
    var n int 
	if a := os.Args; len(a) != 2 {
		fmt.Println("Enter a number, try again.")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf("Cannot convert %v.\n", a[1])   
	} else {
		fmt.Printf("%s *2 is %d.\n", a[1], n*2)
	}
	fmt.Println(n) //shadow of n 
}