//练习 4.10： 修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl-Note/ch4/github"
)

//!+
type Value struct {
	number int
	login  string
	title  string
	time   time.Time
}

func main() {
	// 创建三个切片，分别表示不到一个月的、不到一年的、超过一年。
	t1 := make([]*Value, 0)
	t2 := make([]*Value, 0)
	t3 := make([]*Value, 0)
	// 获取当前时间
	now := time.Now()
	// 进行查询得到result
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	// 得到result的总条数
	fmt.Printf("%d issues:\n", result.TotalCount)
	// 对result的每个Item进行遍历
	for _, item := range result.Items {
		// 用当前时间减去创建的时间再除以(24 * 30),以月为单位
		data := now.Sub(item.CreatedAt).Hours() / (24 * 30)
		if data <= 1 {
			// 进行构建结构体并填充数据
			v := &Value{
				item.Number,
				item.User.Login,
				item.Title,
				item.CreatedAt,
			}
			t1 = append(t1, v)
		} else if data <= 12 {
			v := &Value{
				item.Number,
				item.User.Login,
				item.Title,
				item.CreatedAt,
			}
			t2 = append(t2, v)
		} else if data > 12 {
			v := &Value{
				item.Number,
				item.User.Login,
				item.Title,
				item.CreatedAt,
			}
			t3 = append(t3, v)
		}
		//fmt.Printf("#%-5d %9.9s %.55s,%s\n",
		//	item.Number, item.User.Login, item.Title, item.CreatedAt.String())
	}
	// 分别打印输出不到一个月的、不到一年的、超过一年。
	fmt.Println("--------------- 不到一个月的Issues --------------")
	for _, v := range t1 {
		// 这所以用指针 这里就是证明如果不用的话 v就会赋值一遍然后被取值。
		fmt.Println("时间：", v.time.String(), "数量: ", v.number, "login: ", v.login, "title: ", v.title)
	}
	fmt.Println("--------------- 一年之内的Issue --------------")
	for _, v := range t2 {
		fmt.Println("时间：", v.time, "数量: ", v.number, "login: ", v.login, "title: ", v.title)
	}
	fmt.Println("--------------- 超过一年的Issues --------------")
	for _, v := range t3 {
		fmt.Println("时间：", v.time, "数量: ", v.number, "login: ", v.login, "title: ", v.title)
	}
}
