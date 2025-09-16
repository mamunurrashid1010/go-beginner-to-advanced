package main

import (
	"fmt"
	"strings"
)

/*
Description:
------------
This program demonstrates **string manipulation in Go**
using the built-in `strings` package.

The `strings` package provides many useful functions
for working with strings, such as:
- Changing case (ToUpper, ToLower, Title)
- Trimming spaces
- Checking prefixes/suffixes
- Splitting and Joining strings
- Replacing substrings
- Finding substrings (Contains, Index)
*/
func main() {
	// Sample string
	text := "   Hello, GoLang World!   "

	// 1. Trim spaces
	trimmed := strings.TrimSpace(text)
	fmt.Printf("Original: %q\nTrimmed: %q\n\n", text, trimmed)

	// 2. Change case
	fmt.Println("Uppercase:", strings.ToUpper(trimmed))
	fmt.Println("Lowercase:", strings.ToLower(trimmed))
	fmt.Println("Title Case:", strings.Title(trimmed)) // Capitalize each word
	fmt.Println()

	// 3. Check prefix and suffix
	fmt.Println("Starts with 'Hello'? ->", strings.HasPrefix(trimmed, "Hello"))
	fmt.Println("Ends with 'World!'? ->", strings.HasSuffix(trimmed, "World!"))
	fmt.Println()

	// 4. Replace substrings
	replaced := strings.ReplaceAll(trimmed, "GoLang", "")
	fmt.Println("Replaced Text:", replaced)
	fmt.Println()

	// 5. Split and Join
	words := strings.Split(trimmed, " ")
	fmt.Println("Split into words:", words)

	// 6. Find substring
	fmt.Println("Contains 'World'? ->", strings.Contains(trimmed, "World"))
	fmt.Println("Index of 'GoLang' ->", strings.Index(trimmed, "GoLang"))
	fmt.Println("Count of 'l' ->", strings.Count(trimmed, "l"))

}
