/* 数组 */
// 数组的长度是其类型的一部分，因此数组不能改变大小
// [n]T 表示拥有 n 个 T 类型值的数组，定义方式：
// 1. var array_name [array_length]array_element_type
// 2. var_name := [array_length]array_element_type{array_element, ...}
// 3. ...

package main

import (
	"fmt"
)

func array() {
	var (
		a [2]string // 将变量a声明为拥有2个string的数组
	)

	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println(primes)
}

func slices() {
	arrays := [7]string{"1", "11", "2", "3", "2"} // 赋予的值零值
	var s []string = arrays[1:4]                  // array的切片操作
	fmt.Println(arrays)
	fmt.Println(s)
}

func main() {
	array()
	slices()
}
