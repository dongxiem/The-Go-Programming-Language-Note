package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// 进行计数
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// 进行乘方计算
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// 使用for-range进行通道取值打印
	// 这个for-range模式会感应到通道的关闭
	for x := range squares {
		fmt.Println(x)
	}
}
