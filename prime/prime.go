package main

import "fmt"

func main() {
	var n int
	isPrime := true

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	if n <= 1 {
		isPrime = false
	} else {
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
	}

	if isPrime {
		fmt.Println("Prime Number")
	} else {
		fmt.Println("Not a Prime Number")
	}
}
