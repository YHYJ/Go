/* 零值 */
// 没有声明初始值的变量默认会被赋予它们的零值

package main

import (
	"fmt"
)

func main() {
	var (
		i int
		f float64
		b bool
		s string
		bt byte
		r rune
		c complex64
	)
	fmt.Printf("%v %v %v %v %v %v %v\n", i, f, b, s, bt, r, c)
}
