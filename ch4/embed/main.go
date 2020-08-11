package main

import "fmt"

type Point struct{ X, Y int }

type Circle struct {
	// 匿名成员Point
	Point
	Radius int
}

type Wheel struct {
	// 匿名成员Circle
	Circle
	Spokes int
}

func main() {
	var w Wheel
	//!+
	// 结构体字面值必须遵循形状类型声明时的结构
	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}
	// Printf函数中%v参数包含的#副词，它表示用和Go语言类似的语法打印值。对于结构体类型来说，将包含每个成员的名字。
	fmt.Printf("%#v\n", w)
	// Output:
	// Wheel{Circle:Circle{Point:Point{X:8, Y:8}, Radius:5}, Spokes:20}

	// 我们可以用简短形式访问匿名成员嵌套的成员
	w.X = 42

	fmt.Printf("%#v\n", w)
	// Output:
	// Wheel{Circle:Circle{Point:Point{X:42, Y:8}, Radius:5}, Spokes:20}
	//!-
}
