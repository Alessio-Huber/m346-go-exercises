package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(6) + 1 // Simulate a standard 6-sided die
	var when = time.Now()

	// Open eyes.txt for writing
	eyesFile, err := os.Create("eyes.txt")
	if err != nil {
		fmt.Println("Error creating eyes.txt:", err)
		return
	}
	defer eyesFile.Close() // Ensure the file is closed after writing

	// Open dice.log for writing
	logFile, err := os.Create("dice.log")
	if err != nil {
		fmt.Println("Error creating dice.log:", err)
		return
	}
	defer logFile.Close() // Ensure the file is closed after writing

	// Use fmt.Fprintln to write to eyes.txt
	fmt.Fprintln(eyesFile, "The dice shows", eyes, "eyes")

	// Use fmt.Fprintln to write to dice.log
	fmt.Fprintln(logFile, "The dice was rolled at", when)
}
