package main

import (
	"fmt"
)

// func <func_name>(<params> <param type>) <return type> { ... }
// 可以一次指定多个一样类型的变量
func add(x, y float64) {
	num := x + y
	fmt.Println("x + y =", num)
}

func main() {
	add(3, 3.16)
}
