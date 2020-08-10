package main

import (
	"testing"
)

func BenchmarkEathOther(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EathOther("12345", "54321")
	}
}

func TestEathOther(t *testing.T) {
	if k, v := EathOther("12345", "54321"); k {
		t.Log("ok", v)
	} else {
		t.Fatal("error", v)
	}
}

// D:\Garmen\GoHub\gopl-Note\ch3\task\task3-12>go test -bench=. -cpu=4
// goos: windows
// goarch: amd64
// BenchmarkEathOther-4    22207539                63.2 ns/op
// PASS
// ok      _/D_/Garmen/GoHub/gopl-Note/ch3/task/task3-12   1.692s
