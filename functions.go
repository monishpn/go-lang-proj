package main

import (
	"fmt"
)

func addition(a, b int64) (c int64) {
	c = a + b
	return
}

func main() {
	var a, b int64

	fmt.Print("Enter the value for A")
	fmt.Scanf("%d", &a)
	fmt.Print("Enter the value for B")
	fmt.Scanln(&b)
	c := addition(a, b)
	fmt.Printf("The sum is: %d\n", c)
}
