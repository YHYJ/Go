/* 结构体(struct) */
// 一个结构体就是一组字段(field)

package main

import (
	"fmt"
)

type Vertex struct {
	A, B int
	X    uint
	Y    float32
	Z    string
}

var (
	v = Vertex{
		A: 1, B: 2,
		X: 190, Y: 22,
		Z: "i am struct",
	} // 创建一个Vertex类型的结构体
	v1 = Vertex{A: 1, B: 2} // X、Y和Z被隐式赋予零值
	v2 = Vertex{}
	pv = &Vertex{A: 1, B: 2} // 创建一个 *Vertex 类型的结构体指针
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
	// 没有直接输出内存地址，而是使用了取地址符&, 但实质上输出的还是内存地址
	fmt.Printf("Typr: %T, Value: %v\n", V, V)

	fmt.Println((*V).X) // 和普通指针一样的获取内存空间里的内容的应用方法
	fmt.Println(V.X)    // 但因为这样写太罗嗦，所以go允许隐式间接引用

	/* ---------------------------------- */

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Printf("Typr: %T, Value: %v\n", pv, pv)
}
