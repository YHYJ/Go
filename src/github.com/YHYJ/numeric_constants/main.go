package main

import (
	"fmt"
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func main() {
	fmt.Println(Small)
	fmt.Println(Small * 0.1)
	fmt.Println(Big * 0.1)
}
