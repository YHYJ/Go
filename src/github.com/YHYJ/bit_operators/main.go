/* 位运算符 */

package main

import (
	"fmt"
)

func main() {
	var (
		x uint = 60
		y uint = 13
	)

	fmt.Println("x & y =", x&y)   //按位与
	fmt.Println("x | y =", x|y)   //按位或
	fmt.Println("x ^ y =", x^y)   //按位异或
	fmt.Println("x << y =", x<<y) //按位左移
	fmt.Println("x >> y =", x>>y) //按位右移
}
