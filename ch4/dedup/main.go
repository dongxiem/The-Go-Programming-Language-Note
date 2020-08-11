package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// string/bool 形式的set
	seen := make(map[string]bool)

	input := bufio.NewScanner(os.Stdin)
	// 读取输入数据
	for input.Scan() {
		line := input.Text()
		// 如果输入数据没有在set中，则加入set并打印输出
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
