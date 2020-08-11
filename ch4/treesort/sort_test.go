package treesort_test

import (
	"gopl-Note/ch4/treesort"
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	// 生成随机数据
	for i := range data {
		data[i] = rand.Int() % 50
	}
	// 根据规则进行排序
	treesort.Sort(data)
	// 判断是否按照排序规则
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

// D:\Garmen\GoHub\gopl-Note\ch4\treesort>go test -bench=. -cpu=4
// PASS
// ok      _/D_/Garmen/GoHub/gopl-Note/ch4/treesort        0.244s
//
// D:\Garmen\GoHub\gopl-Note\ch4\treesort>go test
// PASS
// ok      _/D_/Garmen/GoHub/gopl-Note/ch4/treesort        0.216s
