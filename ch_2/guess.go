package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix() // get the current date and time, as an integar
	rand.Seed(seconds)           // seed the random number generator
	target := rand.Intn(100) + 1 // generate an integer between 1 and 100
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin) // create a bufio.Reader, which lets us read keyboard input
	success := false                    //setup to print a failure message by default
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")

		fmt.Print("Make a guess: ")           //ask for a number
		input, err := reader.ReadString('\n') // read what a user types up, until they press Enter
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input) // convert the input string to an integar
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Opps. Your guess was HIGH")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
}
