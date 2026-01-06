package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	secretNumber := rand.Intn(100) + 1

	var guess int
	attempts := 0

	fmt.Println("🎯 Welcome to the Number Guessing Game!")
	fmt.Println("I have selected a number between 1 and 100.")

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)
		attempts++

		if guess < secretNumber {
			fmt.Println("Too low! Try again.")
		} else if guess > secretNumber {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("🎉 Correct! You guessed the number in %d attempts.\n", attempts)
			break
		}
	}
}
