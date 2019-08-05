package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("need to enter a number")
		return
	}

	guess, _ := strconv.Atoi(os.Args[1]) // guess number from CLI and convert to int
	if guess <= 0 {                      // check if the number is a +ve number
		fmt.Println("Please enter a number greater than zero")
		return
	}

	rand.Seed(time.Now().UnixNano()) // seed number

	// generates random number
	if n := rand.Intn(guess + 1); n == guess { // if random number == the entered in the CLI
		fmt.Println("YOU WIN")
		return
	}

	words := strings.Fields("LOSER TRY-AGAIN FAILED") // random messages if you lose
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println()

	fmt.Println(words[0])
}
