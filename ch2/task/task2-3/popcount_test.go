package main

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(1000)
	}
}

func BenchmarkPopCountCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountCycle(1000)
	}
}

// 测试运行如下：
// go test -cpu=4 -bench=.
// goos: windows
// goarch: amd64
// BenchmarkPopCount-4             1000000000               0.460 ns/op
// BenchmarkPopCountCycle-4        52138549                24.9 ns/op
// PASS
// ok      _/D_/Garmen/GoHub/gopl-Note/ch2/task/task2-3    2.062s
