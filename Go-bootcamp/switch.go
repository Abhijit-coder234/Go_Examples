package main

import (
	"fmt"
)

func main() {
	var new int

	switch new {
	case new > 0:
		fmt.Println("positive")
        fallthrough
	case new < 0:
		fmt.Println("negative")
		fallthrough
	default:
		fmt.Println("zero")
	}
}