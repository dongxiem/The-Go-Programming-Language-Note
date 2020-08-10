package main

import (
	"bufio"
	"fmt"
	"os"
)

// basename : 移除目录和'.‘之后的后缀
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	// 从输入中读取字符串
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// 进行处理之后输出
		fmt.Println(basename(input.Text()))
	}
	// NOTE: ignoring potential errors from input.Err()
}
