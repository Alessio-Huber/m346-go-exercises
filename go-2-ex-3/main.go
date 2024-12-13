package main

import "fmt"

func main() {
	// Create a map called "modules"
	modules := make(map[int]string)

	// Adding some initial values to the map
	modules[104] = "Introduction to Programming"
	modules[117] = "Data Structures"
	modules[346] = "Algorithms"

	// Print the initial values
	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// Delete one module (e.g., module 117)
	delete(modules, 117)

	// Add one new module (e.g., module 200)
	modules[200] = "Web Development"

	// Replace one module (e.g., replace module 346 with a new value)
	modules[346] = "Advanced Algorithms"

	// Print the final state of the map
	fmt.Println(modules)
}
