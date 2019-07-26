package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		speed int64 = time.Now().UnixNano()
		n     int   = 100
	)
	rand.Seed(speed)                                   // 设置一个随机种子，Init才会每次生成不一样的数字
	fmt.Println("My favorite number is", rand.Intn(n)) // 生成[0, n)的伪随机数
}
