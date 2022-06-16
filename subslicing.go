package main

import "fmt"

func main() {
	// numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// printSlice(numbers)

	// fmt.Printf("numbers==%v\n", numbers)

	// fmt.Printf("numbers[1:4]==%v\n", numbers[1:4])

	// fmt.Printf("numbers[:3]==%v\n", numbers[:3])

	// fmt.Printf("numbers[4:]==%v\n", numbers[4:])

	// fmt.Printf("numbers[7:]==%v\n", numbers[7:])

	// number1 := make([]int, 0, 5)
	// number2 := numbers[:2]
	// number3 := numbers[3:7]
	// printSlice(number3)

	// printSlice(number2)
	// printSlice(number1)

	var numbers []int
	printSlice(numbers)
	numbers = append(numbers, 0)
	printSlice(numbers)
	numbers = append(numbers, 1)
	printSlice(numbers)
	numbers = append(numbers, 2, 3, 4, 5)
	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
