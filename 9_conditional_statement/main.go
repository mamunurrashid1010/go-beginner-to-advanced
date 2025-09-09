package main

import "fmt"

/*
	Conditional Statements in Go:
	------------------------------
	Go supports conditional logic using `if`, `else if`, and `else`.

	- if: Executes a block of code if the condition is true.
	- else if: Adds more conditions if the previous `if` was false.
	- else: Executes a block of code when all conditions fail.

	Note:
	- Go does not require parentheses around conditions.
	- The curly braces `{}` are mandatory.
*/
func main() {
	// Example variable
	score := 75

	// 1. Simple if statement
	if score >= 80 {
		fmt.Println("Grade: A")
	}

	// 2. if - else statement
	if score >= 50 {
		fmt.Println("Result: Passed")
	} else {
		fmt.Println("Result: Failed")
	}

	// 3. if - else if - else statement
	if score >= 90 {
		fmt.Println("Excellent performance")
	} else if score >= 70 {
		fmt.Println("Good performance") // this will run since score = 75
	} else if score >= 50 {
		fmt.Println("Average performance")
	} else {
		fmt.Println("Poor performance")
	}

	// 4. if with short statement (declare variable inside if)
	if marks := 40; marks >= 50 {
		fmt.Println("Passed with marks:", marks)
	} else {
		fmt.Println("Failed with marks:", marks)
	}
}
