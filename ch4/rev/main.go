package main

import "fmt"

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	// 反转一次
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"

	// 旋转数组
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2]) // 先反转前部分
	reverse(s[2:]) // 再反转后部分
	reverse(s)     // 最后反转整个数组
	fmt.Println(s) // "[2 3 4 5 0 1]"

}
