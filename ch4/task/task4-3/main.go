// 练习 4.3： 重写reverse函数，使用数组指针代替slice。
package main

import "fmt"

// reverse1 : 没有传递指针进来
// 这么写的时候  s 已经是复制值了，所以失败。
func reverse1(s [5]int) {
	// 这时候再去取地址
	st := &s
	// 首尾交换进行逆转
	for i, j := 0, len(*st)-1; i < j; i, j = i+1, j-1 {
		(*st)[i], (*st)[j] = (*st)[j], (*st)[i]
		//fmt.Println(*st)
	}
}

// reverse2 : 传递指针进来，直接这么写就成功了！
func reverse2(s *[5]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		// 直接在指针上面操作即可
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
func main() {
	t := [5]int{1, 443, 223, 2, 0}
	// 测试reverse1
	reverse1(t)
	fmt.Println(t)
	// 测试reverse2
	reverse2(&t)
	fmt.Println(t)
}
