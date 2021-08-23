package main

import (
	"fmt"
	"os"
	"strconv"
	"math/rand"
	"time"
)

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println(usage, maxTurns)
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number")
		return
	}

	if guess < 0 {
		fmt.Println("Pick a positive number")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(100)
		fmt.Println(n)

		if n == guess {
			fmt.Println("You have won")
			return
		}
	}

	fmt.Println("You have lost")

}
