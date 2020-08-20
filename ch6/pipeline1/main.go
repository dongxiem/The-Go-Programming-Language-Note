package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	// 启动一个Counter goroutine
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// 启动一个Squarer goroutine
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// 在主协程main中进行打印Printer
	for {
		fmt.Println(<-squares)
	}
}
