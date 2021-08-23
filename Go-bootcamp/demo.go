package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string
	dir, file = path.Split("bad/new/Go")
	fmt.Println("file:", file)
	fmt.Println("die:", dir)
}