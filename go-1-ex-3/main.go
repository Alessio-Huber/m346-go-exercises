package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var eyes = rand.Intn(6) + 1 // Simulate a standard 6-sided die
	var when = time.Now()

	// Open eyes.txt for writing
	eyesFile, err := os.Create("eyes.txt")
	if err != nil {
		fmt.Println("Error creating eyes.txt:", err)
		return
	}
	defer func() {
		if err := eyesFile.Close(); err != nil {
			fmt.Println("Error closing eyes.txt:", err)
		}
	}()

	// Open dice.log for writing
	logFile, err := os.Create("dice.log")
	if err != nil {
		fmt.Println("Error creating dice.log:", err)
		return
	}
	defer func() {
		if err := logFile.Close(); err != nil {
			fmt.Println("Error closing dice.log:", err)
		}
	}()

	// Use fmt.Fprintln to write to eyes.txt
	fmt.Fprintln(eyesFile, "The dice shows", eyes, "eyes")

	// Use fmt.Fprintln to write to dice.log
	fmt.Fprintln(logFile, "The dice was rolled at", when)
}
