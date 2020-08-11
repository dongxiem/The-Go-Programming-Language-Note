// 练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
package main

import "fmt"

func main() {
	str := []string{"a", "a", "bbb", "bbb", "d", "b", "c", "c"}
	fmt.Println(RemoveDuplicates(str))
}

// RemoveDuplicates : 原地完成消除[]string中相邻重复的字符串
func RemoveDuplicates(str []string) []string {
	// 对字符串数组进行遍历，注意是两个两个的对比，所以是i < len(str)-1
	// 其实就是双指针进行
	for i := 0; i < len(str)-1; i++ {
		// 如果两个字符串相等
		if str[i] == str[i+1] {
			// 则进行copy覆盖
			copy(str[i:], str[i+1:])
			// 同时将最后一个元素剪去
			str = str[:len(str)-1]
			// i再减1，继续当前位置上的判断！
			i--
		}
	}
	return str
}
