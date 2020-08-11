package main

import (
	"fmt"
	"log"
	"os"

	"gopl-Note/ch4/github"
)

func main() {
	// 从输入中得到参数并进行询问
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	// 输出总共的计数
	fmt.Printf("%d issues:\n", result.TotalCount)
	// 对每个Item进行打印输出
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

// go build
// issues repo:golang/go is:open json decoder
