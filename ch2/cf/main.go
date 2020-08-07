package main

import (
	"fmt"
	"os"
	"strconv"

	tempconv "gopl-Note/ch2/tempconvplus"
)

func main() {
	// 循环处理
	for _, arg := range os.Args[1:] {
		// 进行解析
		t, err := strconv.ParseFloat(arg, 64)
		// 判断是否错误
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		// 进行转换
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		// 输出
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
