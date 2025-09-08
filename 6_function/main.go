package main

import "fmt"

/*
	A simple function without parameters and return value
	This function just prints a greeting message
*/
func greet() {
	fmt.Println("Hello! Welcome to the Go Functions Example")
}

/*
	A function with parameters but no return value
	It takes a name as input and prints a personalized greeting
*/
func greetUser(name string) {
	fmt.Printf("Hello, %s! Glad to have you here.\n", name)
}

/*
	A function with parameters and a return value
	It adds two integers and returns the result
*/
func addNumbers(a int, b int) int {
	return a + b
}

/*
	A function with multiple return values
	It divides two numbers and returns both the quotient and remainder
*/
func divideNumbers(divided int, divisor int) (int, int) {
	quotient := divided / divisor
	remainder := divided % divisor
	return quotient, remainder
}
func main() {
	// Calling a function without parameters
	greet()

	// Calling a function with parameters
	greetUser("Mamun")

	// Using a function with return value
	sum := addNumbers(10, 20)
	fmt.Println("The sum of 10 and 20 is:", sum)

	// Using a function with multiple return values
	quotient, remainder := divideNumbers(22, 5)
	fmt.Printf("22 divided by 5 gives quotient: %d and remainder: %d\n", quotient, remainder)
}
