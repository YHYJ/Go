package main

import (
	"fmt"
)

// 一个函数可以有多个返回值，在声明的时候用()括起来
func swap(x, y string) (string, string) {
	return x, y
}

func main() {
	a, b := swap("X", "Y")
	fmt.Printf("a = %s, b = %s\n", a, b)
}
