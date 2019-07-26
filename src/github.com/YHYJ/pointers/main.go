/* 指针指向一个值的内存地址 */
// *T 是指向T类型的指针，零值是nil
// & 是取地址符，会生成一个指向内存地址的指针，可以为指针变量赋值
// * 是取值符，会获取指针指向的内存地址的内容

package main

import (
	"fmt"
)

func main() {
	cache := 5.0
	var x *float64 = &cache // 指针变量x指向的是一个内存地址
	fmt.Printf("Type: %T, Value: %v\n", x, x)
	fmt.Printf("Type: %T, Value: %v\n", *x, *x) // 通过指针x获取cache的值，既所谓的‘间接引用/重定向’
}
