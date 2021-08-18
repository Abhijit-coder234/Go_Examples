package main

import "fmt"

func main() {
	// arr := []int{1, 2, 3, 4}
	// for i, val := range arr {
	// 	fmt.Println(i)
	// 	fmt.Println(val)
	// }

	mymap := make(map[string]interface{})
	mymap["abhijit"] = 22
	mymap["day"] = "sunday"

	for n, k := range mymap {
		fmt.Println(n)
		fmt.Println(k)
	}
}
