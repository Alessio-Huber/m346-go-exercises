package main

import (
	"fmt"
	"math"
)

func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	hypotenuse1 := computeHypotenuse(3, 4)
	fmt.Printf("Hypotenuse for a=3 and b=4: %.2f\n", hypotenuse1)

	hypotenuse2 := computeHypotenuse(5, 12)
	fmt.Printf("Hypotenuse for a=5 and b=12: %.2f\n", hypotenuse2)

	hypotenuse3 := computeHypotenuse(8, 15)
	fmt.Printf("Hypotenuse for a=8 and b=15: %.2f\n", hypotenuse3)

	hypotenuse4 := computeHypotenuse(7, 24)
	fmt.Printf("Hypotenuse for a=7 and b=24: %.2f\n", hypotenuse4)
}
