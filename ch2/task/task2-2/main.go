// 练习 2.2： 写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数
// 如果缺省的话则是从标准输入读取参数
// 然后做类似Celsius和Fahrenheit的单位转换，长度单位可以对应英尺和米，重量单位可以对应磅和公斤等。
package main

import (
	"flag"
	"fmt"
	"strconv"
)

var (
	weight, help *string
	long         *string
	t            string
)

// init ：初始化方法
func init() {
	// !!: 这里不能使用 *flag.String 这种方式，原因是 这个时候还没有解析直接获取值的话是没有任何内容的。
	// 长度重量的默认分别为："m"、"kg"!!!
	long = flag.String("long", "m", "长")
	weight = flag.String("weight", "kg", "重量")
	help = flag.String("help", "", "帮助")
	flag.Parse()
}

func main() {
	fmt.Println("将输入的数值转换为标准单位")
	// 先对长度进行转换
	Transfor(*long)
	// 再对help进行解析
	Transfor(*help)
	// 最后对重量进行转换
	Transfor(*weight)
}

// Transfor : 进行转换的方法
func Transfor(str string) {
	// 输入的数值
	s := flag.Arg(0)
	// 转换为int类型
	st, _ := strconv.Atoi(s)
	// 再转换为float64类型
	stt := float64(st)
	// 判断是要进行哪种转换，默认长度为"m",重量为"kg"
	switch str {
	case "m":
		fmt.Println(stt)
	case "cm":
		fmt.Println(float64(stt / 100))
	case "mm":
		fmt.Println(float64(stt / 1000))
	case "kg":
		fmt.Println(stt)
	case "g":
		fmt.Println(float64(stt / 1000))
	case "t":
		fmt.Println(float64(stt * 1000))
	}
}
