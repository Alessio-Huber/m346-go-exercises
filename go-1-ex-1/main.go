package main

import "fmt"

func main() {
	// Variablen deklarieren und initialisieren
	firstName := "Max"
	lastName := "Mustermann"
	dayOfBirth := 1
	monthOfBirth := 1
	yearOfBirth := 2000
	numberOfSiblings := 2
	heightInMeters := 1.75
	zodiacSign := '\u264E' // Beispiel für Waage

	// Ausgabe der Informationen
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
