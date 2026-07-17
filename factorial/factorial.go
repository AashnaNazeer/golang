package main

import "fmt"

func main() {
	var n int
	fact := 1

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fact *= i
	}

	fmt.Println("Factorial =", fact)
}
