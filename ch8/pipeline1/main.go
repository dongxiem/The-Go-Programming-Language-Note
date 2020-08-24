package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// 进行计数
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// 进行乘方计算
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// 进行通道取值打印
	for {
		fmt.Println(<-squares)
	}
}
