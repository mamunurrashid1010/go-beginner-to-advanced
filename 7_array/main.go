package main

import "fmt"

func main() {
	// -------------------------------
	// 1. Declare an array with fixed size
	// Arrays in Go have a fixed length and cannot be resized
	// Here, we declare an integer array of size 5
	// -------------------------------
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println("Array elements (numbers):", numbers)

	// -------------------------------
	// 2. Declare and initialize an array in one line
	// -------------------------------
	colors := [3]string{"Red", "Green", "Blue"}
	fmt.Println("Array elements (colors):", colors)

	// -------------------------------
	// 3. Let Go infer the array length using "..."
	// -------------------------------
	primes := [...]int{2, 3, 5, 7, 11}
	fmt.Println("Prime numbers:", primes)

	// -------------------------------
	// 4. Accessing elements by index
	// -------------------------------
	fmt.Println("First prime number:", primes[0])
	fmt.Println("Last prime number:", primes[len(primes)-1])

	// -------------------------------
	// 5. Iterating through an array using for loop
	// -------------------------------
	fmt.Println("Iterating over primes using for loop:")
	for i := 0; i < len(primes); i++ {
		fmt.Printf("Index %d: %d\n", i, primes[i])
	}
}
