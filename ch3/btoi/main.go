package main

import "fmt"

// btoi : 将布尔表达式转为int类型
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	fmt.Println(btoi(true))
}
