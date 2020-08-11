// 练习 4.4: 编写一个 rotate 函数，通过一次循环完成旋转。
package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(rotate(a, 2))
}

// rotate : 通过一次循环完成旋转
// 从新创建一个数组，新数组下标为原数组下标加上偏移值
// 如果超出最大长度则从最左边开始
func rotate(s []int, r int) []int {
	// 获取原数组长度，并根据该长度创建一个新的数组
	lens := len(s)
	arr := make([]int, lens)
	// 对s进行遍历，k指的是数组的索引位置
	for k := range s {
		// 添加k个位置
		index := r + k
		// 如果索引位置大于或等于该数组长度，则减去数组长度
		if index >= lens {
			index -= lens
		}
		// 落位
		arr[k] = s[index]
	}
	// 已经完成
	return arr
}
