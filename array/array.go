package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50}
	sum := 0

	for _, value := range arr {
		sum += value
	}

	fmt.Println("Array:", arr)
	fmt.Println("Sum =", sum)
}
