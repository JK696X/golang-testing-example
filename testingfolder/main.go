package main

import "fmt"

func main() {
	fmt.Printf("This is a test print for golang.\n ")
	result := Calculate(2)
	fmt.Printf("Here is the result: %v\n", result)
}

func Calculate(x int) (result int) {
	result = x + 2
	return result
}
