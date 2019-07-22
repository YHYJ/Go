package main

import (
	"fmt"
)

func main() {
	var (
		x uint = 60
		y uint = 13
	)

	fmt.Println("x & y =", x&y)
	fmt.Println("x | y =", x|y)
	fmt.Println("x ^ y =", x^y)
	fmt.Println("x << y =", x<<y)
	fmt.Println("x >> y =", x>>y)
}
