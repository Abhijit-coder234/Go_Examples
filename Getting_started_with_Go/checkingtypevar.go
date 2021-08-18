package main

import "fmt"

func main() {
	var a interface{}
	a = 1
	// fmt.Println("Enter the info to a: ")
	// fmt.Scanln(&a)
	switch a.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	case bool:
		fmt.Println("Bool")
	default:
		fmt.Println("not defined")
	}
}