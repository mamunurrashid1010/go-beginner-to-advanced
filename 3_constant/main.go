package main

import "fmt"

/*
	In Go, a constant is a value that cannot be changed once defined.
	Constants are useful when you have values that stay the same throughout the program.
*/

const Pi = 3.1416 // Constant for the value of Pi

// group multiple constants
const (
	Language = "Go"
	Version  = 1.25
	Year     = 2025
)

func main() {
	// Printing single constant
	fmt.Println("Value of Pi:", Pi)

	// Printing grouped constants
	fmt.Println("Programming Language:", Language)
	fmt.Println("Current Version:", Version)
	fmt.Println("Year:", Year)

	// Constants can also be used in expressions
	const Radius = 5
	const Area = Pi * Radius * Radius // area of a circle

	fmt.Println("Radius:", Radius)
	fmt.Println("Area of Circle:", Area)

}
