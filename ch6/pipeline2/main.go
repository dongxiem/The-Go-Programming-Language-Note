package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		// 进行naturals 通道的关闭
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		// 进行squares 通道的关闭
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}
