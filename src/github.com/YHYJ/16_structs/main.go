/* File: main.go */
/* Auther: YJ */
/* Email: yj1516268@outlook.com */
/* Created Time: 2020-01-16 15:17:15 */

// Description: 数据类型`struct`
// 一个结构体是一组字段(field)的集合

package main

import (
	"fmt"
)

// Vertex 是一个结构体
type Vertex struct {
	A, B int
	X    uint
	Y    float32
	Z    string
}

// 创建结构体`Vertex`的实例
var (
	v = Vertex{
		A: 1, B: 2,
		X: 190, Y: 22,
		Z: "i am struct",
	} // 创建一个Vertex类型的结构体
	v1 = Vertex{A: 1, B: 2}  // X、Y和Z被隐式赋予零值
	v2 = Vertex{X: 1, A: 12} // 字段名顺序无关
	v3 = Vertex{}
)

func main() {
	fmt.Println(v)
	// 使用'struct.field'访问结构体字段
	fmt.Println(v.X)
	fmt.Println(v.Y)
	fmt.Println(v.Z)

	/* ---------------------------------- */

	// 使用结构体指针访问结构体字段
	V := &v
	// 注意结构体指针不同于普通指针
	// 为了允许隐式间接引用，这里V的输出内容做了调整
	// 没有直接输出内存地址，而是输出了取地址符&, 但实质上输出的还是内存地址
	fmt.Printf("V -- Typr: %T, Value: %v\n", V, V)
	V.X = 110
	fmt.Println((*V).X, V.X) // 和普通指针一样的获取内存空间里的数据的方法，但因为这样写太啰嗦，所以go允许隐式间接引用
	fmt.Printf("V -- Typr: %T, Value: %v\n", *V, *V)

	/* ---------------------------------- */

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
}
