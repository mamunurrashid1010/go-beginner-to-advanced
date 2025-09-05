package main

import (
	"bufio"   // For buffered input
	"fmt"     // For printing output
	"os"      // For accessing standard input
	"strconv" // For converting string input to numbers
	"strings" // For trimming newline characters
)

func main() {
	// Create a new reader to read input from the console (stdin)
	reader := bufio.NewReader(os.Stdin)

	// -------------------------------
	// Example 1: Read a string input
	// -------------------------------
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n') // Read input until user presses Enter
	name = strings.TrimSpace(name)     // Remove newline/extra spaces
	fmt.Println("Hello,", name)        // Print greeting with the entered name

	// -------------------------------
	// Example 2: Read an integer input
	// -------------------------------
	fmt.Print("Enter your age: ")
	ageInput, _ := reader.ReadString('\n') // Read age as string
	ageInput = strings.TrimSpace(ageInput) // Clean up input
	age, err := strconv.Atoi(ageInput)     // Convert string to integer
	if err != nil {
		fmt.Println("Invalid age input!")
		return
	}
	fmt.Println("You are", age, "years old.")

	// -------------------------------
	// Example 3: Read a float number
	// -------------------------------
	fmt.Print("Enter your height (in meters): ")
	heightInput, _ := reader.ReadString('\n') // Read height as string
	heightInput = strings.TrimSpace(heightInput)
	height, err := strconv.ParseFloat(heightInput, 64) // Convert to float
	if err != nil {
		fmt.Println("Invalid height input!")
		return
	}
	fmt.Printf("Your height is %.2f meters.\n", height) // Print with formatting
}
