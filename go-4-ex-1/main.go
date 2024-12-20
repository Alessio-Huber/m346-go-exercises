package main

import "fmt"

// computeGrade berechnet die Note basierend auf der erreichten Punktzahl und der maximalen Punktzahl.
func computeGrade(gotPoints float64, maxPoints float64) float64 {
	if maxPoints == 0 {
		return 6.0 // Wenn die maximale Punktzahl 0 ist, geben wir die schlechteste Note zurück.
	}
	grade := 5/maxPoints*gotPoints + 1
	return grade
}

func main() {
	fmt.Println(computeGrade(17.5, 28.0))
	fmt.Println(computeGrade(25.0, 30.0))
	fmt.Println(computeGrade(0.0, 20.0))
}
