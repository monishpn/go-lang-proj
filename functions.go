package main

import (
	"fmt"
)

func addition(a, b int64) int64 {
	return a + b
}

func main() {
	c := addition(4, 5)
	fmt.Printf("The sum is: %d\n", c)
}
