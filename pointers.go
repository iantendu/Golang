package main

import "fmt"

func main() {
	var a int = 20
	var ip *int
	ip = &a

	fmt.Printf("The address of a variable: %x\n", &a)

	fmt.Printf("The address stored in ip variable: %x\n", ip)

	fmt.Printf("Value of *ip variable: %d\n", *ip)

}
