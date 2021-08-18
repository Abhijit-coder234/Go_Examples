package main

import "fmt"

type Car interface {
	print() 
	Name()        
}

type hyundai struct {
	name  string
	price float64
	make  string
}

type maruti struct {
	name string
	price float64
	make string
}

func (h hyundai) print() {
	fmt.Println(h)
}

func (h hyundai) Name() {
	fmt.Print(h.name)
}

func (m maruti) print() {
	fmt.Println(m)
}

func (m maruti) Name() {
	fmt.Print(m.name)
}

func main() {
	h1 := hyundai{name: "tucson", price:300, make: "abba"}
	m1 := maruti{name: "m800", price:500, make: "baaba"}
	h1.print()
	m1.print()
}

