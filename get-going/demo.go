package main

import (
	"fmt"
	"strings"
)

func main() {
	m := 1
	n := 2
	arr := []int{1, 2, 3, 4}
	new := "hello"
	fmt.Println(new)
	new = "bye abhijit"
	fmt.Println(new)
	fmt.Println(m + n)
	fmt.Println("Hello world!")
    fmt.Println(strings.Count(new, "bye"))
	fmt.Println(strings.Compare(new, "abhijit"))
	fmt.Println(arr)
	arr = append(arr,3)
	fmt.Println(arr)
}
