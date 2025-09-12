package main

import "fmt"

/*
   Program Description:
   --------------------
   This Go program demonstrates the usage of the "switch" statement.
   A switch statement is a cleaner alternative to multiple "if-else-if"
   conditions. It evaluates an expression and executes the matching case.
   We also show how to use "default" when no case matches.
*/

func main() {
	// Example variable to test switch
	day := 3

	// Switch statement to check the value of day
	switch day {
	case 1:
		fmt.Println("Today is Monday")
	case 2:
		fmt.Println("Today is Tuesday")
	case 3:
		fmt.Println("Today is Wednesday")
	case 4:
		fmt.Println("Today is Thursday")
	case 5:
		fmt.Println("Today is Friday")
	case 6:
		fmt.Println("Today is Saturday")
	case 7:
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Invalid day number")
	}

	// Switch without an expression (acts like if-else chain)
	num := -5
	switch {
	case num > 0:
		fmt.Println("Number is Positive")
	case num < 0:
		fmt.Println("Number is Negative")
	default:
		fmt.Println("Number is Zero")
	}

}
