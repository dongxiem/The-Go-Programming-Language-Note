// 练习 4.1： 编写一个函数，计算两个SHA256哈希码中不同bit的数目。（参考2.6.2节的PopCount函数。)
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	sha1 := sha256.Sum256([]byte("123"))
	sha2 := sha256.Sum256([]byte("456"))
	fmt.Println(sha1, sha2)
	n := SHA256Count(sha1, sha2)
	fmt.Println("不同的数目是:", n)

}

// SHA256Count : 计算两个SHA256哈希码中不同bit的数目
func SHA256Count(shaOne, shaTwo [32]byte) int {
	var difference int
	for i := 0; i < len(shaOne); i++ {
		if shaTwo[i] != shaOne[i] {
			difference++
		}
	}
	return difference
}
