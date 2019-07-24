/* 逻辑运算符 */

package main

import (
	"fmt"
)

func main() {
	var x bool = true
	var y bool = false

	fmt.Println("x =", x)
	fmt.Println("y =", y)

	fmt.Println("x && y =", x && y) // and
	fmt.Println("x || y =", x || y) // or
	fmt.Println("!x =", !x)         // not
}
