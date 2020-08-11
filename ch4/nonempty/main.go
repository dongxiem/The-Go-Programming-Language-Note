package main

import "fmt"

// nonempty : 在原有slice内存空间之上返回不包含空字符串的列表
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			// 输入的slice和输出的slice共享一个底层数组。
			// 这可以避免分配另一个数组，不过原来的数据将可能会被覆盖
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// nonempty2 : 使用append函数实现
func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	// 测试nonempty
	data1 := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data1)) // `["one" "three"]`
	fmt.Printf("%q\n", data1)           // `["one" "three" "three"]`

	// 测试nonempty2
	data2 := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty2(data2)) // `["one" "three"]`
	fmt.Printf("%q\n", data2)            // `["one" "three" "three"]`
}
