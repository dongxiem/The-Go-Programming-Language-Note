package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// basename : 使用库函数来移除目录和'.‘之后的后缀
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	// 找到使用'/'的最后的位置
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	// 将'/'之后的所有字符串重新组成一个切片
	s = s[slash+1:]
	// 找到'.'出现的最后位置
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		// 进行截取，不包含该'.'
		s = s[:dot]
	}
	return s
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
}
