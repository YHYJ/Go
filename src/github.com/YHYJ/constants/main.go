// 常量不能用':='语法声明

package main

import (
	"fmt"
)

const Pi float64 = 3.1415926535897932384626

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Pi =", Pi)
}
