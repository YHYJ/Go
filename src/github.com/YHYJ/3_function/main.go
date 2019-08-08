/* 函数 */
// 函数可以有0~N个参数

package main

import (
	"fmt"
)

// function add一次接受两个float64类型的参数
func add(x, y float64) {
	num := x + y
	fmt.Println("x + y =", num)
}

func main() {
	add(3, 3.16)
}
