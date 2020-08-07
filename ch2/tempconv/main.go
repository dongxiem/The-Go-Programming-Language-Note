package main

import "fmt"

// Celsius 和 Fahrenheit虽然有着相同的底层类型float64，但是它们是不同的数据类型
// 因此它们不可以被相互比较或混在一个表达式运算
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

// 将 Celsius 类型转换为 Fahrenheit 类型
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// 将 Fahrenheit 类型转换为 Fahrenheit 类型
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// 该方法返回该类型对象c带着°C温度单位的字符串
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func main() {
	// 1.转换测试
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	// fmt.Printf("%g\n", boilingF-FreezingC)       // 编译错误："compile error: type mismatch"

	// 2.对比测试
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // "true"
	fmt.Println(f >= 0) // "true"
	// fmt.Println(c == f)          // 编译错误：“compile error: type mismatch”
	fmt.Println(c == Celsius(f)) // "true"!

	// 3.string测试
	tmp := FToC(212.0)
	fmt.Println(tmp.String()) // "100°C"
	fmt.Printf("%v\n", tmp)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", tmp)   // "100°C"
	fmt.Println(tmp)          // "100°C"
	fmt.Printf("%g\n", tmp)   // "100"; does not call String
	fmt.Println(float64(tmp)) // "100"; does not call String
}
