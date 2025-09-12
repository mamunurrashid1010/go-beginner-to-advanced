package main

import "fmt"

/*
   Program Description:
   --------------------
   This Go program demonstrates the usage of "map".
   A map is a built-in data type in Go that stores key-value pairs.
   - Keys must be unique and comparable (e.g., string, int).
   - Values can be any type.

   Key Operations:
   1. Create a map
   2. Insert values
   3. Access values
   4. Update values
   5. Delete values
   6. Check if a key exists
   7. Iterate over a map
*/

func main() {
	// Example 1: Creating and initializing a map
	studentGrades := map[string]int{
		"Hasan": 90,
		"Noman": 80,
		"Arif":  75,
	}

	fmt.Println("Initial Map:", studentGrades)

	// Example 2: Adding a new key-value pair
	studentGrades["Kamal"] = 85
	fmt.Println("After adding Kamal:", studentGrades)

	// Example 3: Accessing a value by key
	fmt.Println("Hasan's Grade:", studentGrades["Hasan"])

	// Example 4: Updating a value
	studentGrades["Arif"] = 77
	fmt.Println("After updating Arif's grade:", studentGrades)

	// Example 5: Deleting a key
	delete(studentGrades, "Arif")
	fmt.Println("After deleting Arif:", studentGrades)

	// Example 6: Checking if a key exists
	grade, exists := studentGrades["Kamal"]
	if exists {
		fmt.Println("Kamal's grade is:", grade)
	} else {
		fmt.Println("Kamal not found in map")
	}

	// Example 7: Iterating over a map
	fmt.Println("\nIterating over map:")
	for name, grade := range studentGrades {
		fmt.Printf("Student: %s, Grade: %d\n", name, grade)
	}
}
