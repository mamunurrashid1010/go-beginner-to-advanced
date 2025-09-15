package main

import "fmt"

/*
   Program Description:
   --------------------
   This Go program demonstrates the concept of pointers.
   A pointer is a variable that stores the memory address of another variable.

   Why use pointers?
   - To directly access and modify the memory location of a variable.
   - Useful when passing large data structures to functions (avoid copying).
   - Helps in efficient memory management.

   Key Concepts:
   1. Declare a pointer
   2. Assign the address of a variable to a pointer
   3. Dereference a pointer to access the value
   4. Update a variable’s value through a pointer
   5. Pass pointer to functions
*/

// Function that takes a pointer and modifies the value
func updateValue(num *int) {
	*num = *num + 10 // Dereference and update the value
}

func main() {
	// Example 1: Normal variable
	x := 5
	fmt.Println("Initial value of x:", x)

	// Example 2: Pointer to x
	var p *int = &x // p stores the memory address of x
	fmt.Println("Address of x:", p)
	fmt.Println("Value of x using pointer:", *p) // Dereference pointer

	// Example 3: Modify value using pointer
	*p = 20
	fmt.Println("\nAfter modifying via pointer:")
	fmt.Println("Value of x:", x)
	fmt.Println("Value via pointer:", *p)

	// Example 4: Pass pointer to function
	fmt.Println("\nBefore calling updateValue, x =", x)
	updateValue(&x) // Pass address of x
	fmt.Println("After calling updateValue, x =", x)
}
