package main

import "fmt"

func main() {

	var numbers = make([]int, 5, 5)

	printSlice(numbers)

}
func printSlice(x []int) {

	fmt.Printf("len=%d cap=%d slice=%d", len(x), cap(x), x)

}
