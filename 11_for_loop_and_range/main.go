package main

import "fmt"

/*
   Program Description:
   --------------------
   This Go program demonstrates:
   1. The basic "for" loop.
   2. The "for-range" loop which is commonly used to iterate over arrays, slices, maps, and strings.

   Key Points:
   - "for" is the only looping construct in Go.
   - "range" returns index and value when iterating over collections.
*/
func main() {
	// Example 1: Basic for loop (print numbers 1 to 5)
	fmt.Println("Basic for loop:")
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
	}

	// Example 2: For loop as a while loop (print numbers until condition is false)
	fmt.Println("\nFor loop acting like while loop:")
	x := 1
	for x <= 3 { // no init and post statement, just condition
		fmt.Println("x is:", x)
		x++
	}

	// Example 3: Using range with a slice
	fmt.Println("\nUsing range with a slice:")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Example 4: Using range with a string
	fmt.Println("\nUsing range with a string:")
	word := "GoLang"
	for index, char := range word {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

}
