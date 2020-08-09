package main

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(600)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountRewrite3(600)
	}
}

// go test -bench=. -cpu=4
// goos: windows
// goarch: amd64
// BenchmarkPopCount-4     1000000000               0.486 ns/op
// BenchmarkPopCount2-4    157997226                8.10 ns/op
// PASS
// ok      _/D_/Garmen/GoHub/gopl-Note/ch2/task/task2-5    2.847s
