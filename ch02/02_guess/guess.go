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
	seeder := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seeder)
	target := r.Intn(100) + 1

	fmt.Println("I have chosen a number between 1 and 100: ")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)
	success := false
	for i := 0; i < 10; i++ {
		fmt.Println("You have", 10-i, "Guesses left")
		fmt.Println("Make your guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops. Your guess was LOW!")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH!")
		} else {
			success = true
			fmt.Println("Good job! You guessed it")
			break
		}

	}
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was: ", target)
	}

}
