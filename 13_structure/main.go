package main

import "fmt"

/*
   Program Description:
   --------------------
   This Go program demonstrates the usage of "struct" (structure).
   A struct in Go is a user-defined data type that groups together
   different fields (variables) under one name.

   Why use struct?
   - To represent real-world entities (like Student, Employee, Product).
   - Each struct can have multiple fields of different data types.
   - It allows organizing related data in one place.

   Key Operations:
   1. Define a struct
   2. Create struct variables
   3. Initialize struct values
   4. Access struct fields
   5. Update struct fields
   6. Use struct with functions
*/

// Define a struct for Student
type Student struct {
	ID     int
	Name   string
	Age    int
	Grade  string
	Active bool
}

func printStudentDetails(s Student) {
	fmt.Printf("ID: %d, Name: %s, Age: %d, Grade: %s, Active: %t\n", s.ID, s.Name, s.Age, s.Grade, s.Active)
}
func main() {
	// Example 1: Declare a struct variable and assign value
	var student1 Student
	student1.ID = 1
	student1.Name = "Hasan"
	student1.Age = 24
	student1.Grade = "A"
	student1.Active = true

	fmt.Println("Student 1 details:")
	printStudentDetails(student1)

	// Example 2: Initialize struct using shorthand
	student2 := Student{
		ID:     2,
		Name:   "Kamel",
		Age:    15,
		Grade:  " A+",
		Active: true,
	}
	fmt.Println("\nStudent 2 details:")
	printStudentDetails(student2)

	// Example 3: Partial initialization (remaining fields get zero values)
	student3 := Student{ID: 3, Name: "Arif", Age: 19}
	fmt.Println("\nStudent 3 details (partial init):")
	printStudentDetails(student3)

	// Example 4: Update struct field
	student3.Grade = "A+"
	student3.Active = true
	fmt.Println("\nStudent 3 details (after update):")
	printStudentDetails(student3)
}
