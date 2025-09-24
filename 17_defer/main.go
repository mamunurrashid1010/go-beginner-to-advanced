package main

import (
	"fmt"
)

/*
	Program Description:
	--------------------
	This program demonstrates the use of the `defer` statement in Go.

	# What is `defer`?
	- The `defer` keyword is used to delay the execution of a function until
	  the surrounding function (where `defer` is written) returns.
	- Deferred calls are executed in **Last In, First Out (LIFO)** order.
	- It is commonly used for cleanup tasks like closing files,
	  unlocking resources, or releasing connections.

	This example shows how multiple `defer` statements are executed in reverse order,
	and how defer ensures cleanup happens even if the function exits early.
*/

func main() {
	fmt.Println("Start of main function")

	// First defer statement (executed last)
	defer fmt.Println("First deferred: Closing resources")

	// Second defer statement (executed second)
	defer fmt.Println("Second deferred: Logging out user")

	// Third defer statement (executed first among deferred calls)
	defer fmt.Println("Third deferred: Ending session")

	// Normal execution
	fmt.Println("Doing some work in the main function...")

	// Demonstrate early return (deferred functions will still run)
	if true {
		fmt.Println("Returning early from main...")
		return
	}

	// This line will never execute because of the return above
	fmt.Println("End of main function")
}
