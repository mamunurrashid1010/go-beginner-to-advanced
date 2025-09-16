package main

import (
	"fmt"
	"strconv"
)

/*
Description:
------------
This program demonstrates **data conversion (type casting)** in Go.
Go is a statically typed language, which means you must explicitly
convert (cast) one data type to another when needed.

- Basic conversions: int to float, float to int
- Converting numeric values to string
- Converting string values to numeric (using strconv)
*/
func main() {
	// Example 1: int to float
	var numInt int = 20
	var numFloat float64 = float64(numInt) // explicit conversion
	fmt.Printf("Interger: %d, Converted to Float: %.2f\n", numInt, numFloat)

	// Example 2: float to int
	var priceFloat float64 = 55.50
	var priceInt int = int(priceFloat)
	fmt.Printf("Float: %.2f, Converted to Int: %d\n", priceFloat, priceInt)

	// Example 3: int to string
	var total int = 360
	var totalString string = strconv.Itoa(total) // Itoa = Integer to ASCII (string)
	fmt.Printf("Integer: %d, Converted to String: %q\n", total, totalString)

	// Example 4: string to int
	var strNumber string = "123"
	strInt, err := strconv.Atoi(strNumber) // Atoi = ASCII (string) to Integer
	if err != nil {
		fmt.Println("Conversion error:", err)
	} else {
		fmt.Printf("String: %q, Converted to Int: %d\n", strNumber, strInt)
	}

	// Example 5: string to float
	var strFloat string = "45.67"
	floatNum, err := strconv.ParseFloat(strFloat, 64) // ParseFloat converts string to float64
	if err != nil {
		fmt.Println("Conversion error:", err)
	} else {
		fmt.Printf("String: %q, Converted to Float: %.2f\n", strFloat, floatNum)
	}
}
