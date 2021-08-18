package main

import "fmt"

func main() {
	mymap := make(map[string]interface{})
	mymap["abhijit"] = 22
	mymap["aishu"] = "22"
	fmt.Println(mymap)
}
