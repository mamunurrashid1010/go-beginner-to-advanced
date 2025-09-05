package main

import "fmt"

func main() {

	/*
		Topics:
		fmt.Print()		: Just prints text, no newline unless you add \n.
		fmt.Println()	: Prints text and automatically adds a newline.
		fmt.Printf()	: Prints text with formatting options (%d, %s, %f).
	*/

	// fmt.Print - prints text exactly as it is (no new line automatically)
	fmt.Print("Hello")
	fmt.Print(" World!") // This continues on the same line
	fmt.Print("\n")      // You need to add "\n" manually if you want a new line

	// fmt.Println - prints text and adds a new line at the end automatically
	fmt.Println("Hello")
	fmt.Println("World!") // Each will be printed on a new line

	// fmt.Printf - allows formatting with verbs (like placeholders)
	// %s = string, %d = integer, %f = float, %t = boolean, etc.
	title := "Test title"
	mark := 85
	gpa := 4.5

	fmt.Printf("Title %s and Mark %d \n", title, mark)
	fmt.Printf("GPA is %.1f \n", gpa) // .1f = 1 decimal place
}
