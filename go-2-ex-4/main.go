package main

import "fmt"

// Declare a type for Student (with first and last name)
type Student struct {
	FirstName string
	LastName  string
}

// Declare a type for Class (consisting of multiple students)
type Class struct {
	ClassName string
	Students  []Student
}

// Declare a map of modules being attended by multiple classes
type Module struct {
	ModuleName string
	Classes    []Class
}

func main() {
	// Create some students
	student1 := Student{FirstName: "Alice", LastName: "Smith"}
	student2 := Student{FirstName: "Bob", LastName: "Johnson"}
	student3 := Student{FirstName: "Charlie", LastName: "Brown"}

	// Create a class with students
	class1 := Class{
		ClassName: "Computer Science 101",
		Students:  []Student{student1, student2},
	}

	class2 := Class{
		ClassName: "Mathematics 101",
		Students:  []Student{student3},
	}

	// Create a map of modules attended by multiple classes
	modules := map[string]Module{
		"CS101": {
			ModuleName: "Introduction to Computer Science",
			Classes:    []Class{class1},
		},
		"MA101": {
			ModuleName: "Basic Mathematics",
			Classes:    []Class{class2},
		},
	}

	// Output everything using fmt.Println()
	for moduleCode, module := range modules {
		fmt.Println("Module Code:", moduleCode)
		fmt.Println("Module Name:", module.ModuleName)
		for _, class := range module.Classes {
			fmt.Println("  Class Name:", class.ClassName)
			for _, student := range class.Students {
				fmt.Printf("    Student: %s %s\n", student.FirstName, student.LastName)
			}
		}
	}
}
