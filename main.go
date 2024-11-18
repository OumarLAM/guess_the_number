package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))
	correctNumber := rand.Intn(101) // Random number between 1 and 100

	fmt.Println("Welcome to the Guess the Number Game!")
	fmt.Println("I have chosen a number between 1 and 100. Can you guess it?")
	fmt.Println("The faster you guess, the better!")

	// Start the timer
	startTime := time.Now()

	for {
		// Prompt the player for input
		fmt.Print("Enter your guess: ")
		var input string
		fmt.Scanln(&input)

		// Trim spaces and validate input
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 100.")
			continue
		}		

		// Check the guess
		if guess < correctNumber {
			fmt.Println("Too low! Try again.")
		} else if guess > correctNumber {
			fmt.Println("Too high! Try again.")
		} else {
			// Stop the timer
			elapsedTime := time.Since(startTime)
			// Announce the winner
			fmt.Printf("🎉 Congratulations! You guessed the correct number: %d\n", correctNumber)
			fmt.Printf("It took you %.2f seconds.\n", elapsedTime.Seconds())
			break
		}
	}
}
