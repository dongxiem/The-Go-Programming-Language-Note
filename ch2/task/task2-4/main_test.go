package main

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(600)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountRewrite2(600)
	}
}

// go test -cpu=4 -bench=.
// goos: windows
// goarch: amd64
// BenchmarkPopCount-4     1000000000               0.457 ns/op
// BenchmarkPopCount2-4    12113343                92.9 ns/op
// PASS
// ok      _/D_/Garmen/GoHub/gopl-Note/ch2/task/task2-4    1.973s
