package main

import "fmt"

/*
	Slice in Go:
	-------------
	- A slice is a lightweight, flexible view of an array.
	- Unlike arrays, slices do not have a fixed size.
	- Slices can grow or shrink dynamically.
	- A slice has three properties:
		1. Pointer: reference to the underlying array.
		2. Length: number of elements in the slice.
		3. Capacity: maximum size up to which slice can grow.

	This program demonstrates:
	- Creating a slice
	- Appending elements
	- Slicing (sub-slice)
	- Getting length and capacity
*/
func main() {
	// 1. Creating a slice directly with values
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Original slice:", numbers)

	// 2. Append new elements to a slice (it increases size dynamically)
	numbers = append(numbers, 60, 70)
	fmt.Println("After appending:", numbers)

	// 3. Creating a sub-slice (from index 1 to 3 -> excludes index 3)
	subSlice := numbers[1:3]
	fmt.Println("Sub-slice (from index 1 to 3):", subSlice)

	// 4. Checking length and capacity of slice
	fmt.Println("Length of slice:", len(numbers))   // number of elements
	fmt.Println("Capacity of slice:", cap(numbers)) // total available space

	// 5. Using make() to create a slice with defined length and capacity
	newSlice := make([]int, 3, 5) // length=3, capacity=5
	newSlice[0] = 100
	newSlice[1] = 200
	newSlice[2] = 300
	fmt.Println("Slice created with make():", newSlice)
	fmt.Println("Length:", len(newSlice), "Capacity:", cap(newSlice))
}
