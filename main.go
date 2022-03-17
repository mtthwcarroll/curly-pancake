package main

// Imports
import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
)

// Main function
func main() {
	// Setting up a regex for validation use
	expr, _ := regexp.Compile("[0-9]")

	fmt.Println("Welcome to my number guessing game. I am thinking of a number" +
		"\nbetween 0 and 100. Make a guess, and I will tell you if my number is lower\n" +
		"or higher than your guess!")

	// The number the use will try and guess. Have to add 1 because the lower limit
	// is 0 and that can't be changed
	goal := rand.Intn(98) + 1

	// Variable for storing user guess
	var guess string
	// Declare and assign i for the for loop
	i := 0
	// Essentially a while loop that breaks when we increment i
	for i < 1 {
		// Read user input
		fmt.Scanln(&guess)
		// Validate input
		if expr.MatchString(guess) && len(guess) < 3 && len(guess) > 0 {
			guessInt, _ := strconv.Atoi(guess)
			if guessInt < goal {
				fmt.Println("My number is higher!")
			} else if guessInt > goal {
				fmt.Println("My number is lower!")
			} else {
				fmt.Println("You guessed it, good job! Goodbye!")
				i++
			}
		} else {
			fmt.Println("Please enter a number between 0 and 100")
		}
	}
}
