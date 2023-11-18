package main

import "fmt"

// Define a struct named 'Person'
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// Create an instance of the 'Person' struct with a composite literal.
	// It is of value type, meaning that when you assign person1 to another
	// variable or pass it to a function, a copy of the struct is made.
	person1 := Person{
		FirstName: "John",
		Age:       30,
	}

	// declared as a pointer to a Person struct.
	// It is of reference type, meaning that when you assign person2 to another
	// variable or pass it to a function, a copy of the pointer is made, but the
	// underlying struct is not copied.
	person2 := &Person{
		FirstName: "Jane",
		Age:       25,
	}

	// Access and print struct fields
	fmt.Println("First Name:", person1.FirstName)
	fmt.Println("Age:", person1.Age)

	// Access and print struct fields through a pointer
	fmt.Println("Person 2 - First Name:", (*person2).FirstName) // Dereference the pointer
	// Alternatively, Go allows shorthand notation for accessing fields through a pointer
	fmt.Println("Person 2 - Last Name:", person2.LastName) // Shorthand notation

	// Modify a Person (value type)
	modifyPersonValue(person1)
	// Age is not modified
	fmt.Println("Person 1 - Age:", person1.Age)

}

// Function to modify a Person (value type)
func modifyPersonValue(p Person) {
	p.Age = 40
}
