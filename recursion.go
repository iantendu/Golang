package main

import "fmt"

func factorial(i int) int {

	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

func main() {

	for i := 0; i < 16; i++ {
		fmt.Println("Factorial of %d is %d", i, factorial(i))
	}

}
