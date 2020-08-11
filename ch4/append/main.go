package main

import "fmt"

// appendInt : 专门用于处理[]int类型的slice的插入处理
func appendInt(x []int, y int) []int {

	var z []int
	zlen := len(x) + 1
	// 如果当前切片长度+1 小于等于切片的容量，还可以继续进行插入
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		// 当前切片长度+1 已经大于切片容量了，需要扩容
		zcap := zlen
		// 进行二倍扩容
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		// 根据zlen创建一个新的切片，长度为zlen，容量为zcap
		z = make([]int, zlen, zcap)
		// 最后进行复制
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		// 使用cap()函数可以直接输出容量
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
