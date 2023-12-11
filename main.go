package main

import (
	"fmt"
	"math/rand"
	"time"
)

var number int

func LoopLearn() {
	for i := 0; i <= 1000; i++ {
		if i%100 == 0 {
			rand.Seed(time.Now().UnixNano())
			number = rand.Intn(100) // Set number to a random value and limit it for readability
			fmt.Printf("Index %d, New Number: %d\n", i, number)
		}
	}
	// Calculate the square of the number
	fmt.Println("The square of the number is:", number*number)
}

func Add(a, b int) int {
	return a + b
}

func main() {
	LoopLearn()
	sum := Add(5, 7)
	fmt.Printf("Sum of 5 and 7 is: %d\n", sum)
}
