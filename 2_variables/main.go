package main

import "fmt"

func main() {
	// 1. Declaring single variable with type
	var title1 string = "Title 1" // string
	fmt.Println("Title 1 :", title1)

	var age int = 28 // integer
	fmt.Println("Age :", age)

	var temperature float32 = 28.56 // float32
	fmt.Println("Temperature :", temperature)

	var isActive bool = true // bool
	fmt.Println("Active :", isActive)

	// 2. Declaring single variable without type (type inferred)
	var title2 = "Title 2" // string
	fmt.Println("Title 2 :", title2)

	var roll = 21 // integer
	fmt.Println("Roll :", roll)

	// 3. Short variable declaration (inside function only)
	country := "Bangladesh"
	fmt.Println("Country :", country)

	// 4. Multiple variable declaration in one line
	var x, y, z int = 1, 2, 3
	fmt.Println("x, y, z:", x, y, z)

	// 5. Multiple variable declaration with type inference
	a, b, c := "Go", true, 3.14
	fmt.Println("a, b, c:", a, b, c)

	// 6. Grouped variable declaration
	var (
		framework  = "Gin"
		version    = 1.22
		isReleased = true
	)
	fmt.Println("Framework:", framework, "Version:", version, "Released:", isReleased)

	// 7. Zero values (default initialization)
	var zeroString string
	var zeroInt int
	var zeroBool bool
	var zeroFloat float64
	fmt.Println("Zero Values :", zeroString, zeroInt, zeroBool, zeroFloat)

}
