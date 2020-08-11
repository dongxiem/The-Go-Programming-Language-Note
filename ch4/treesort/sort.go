package treesort

type tree struct {
	value       int
	left, right *tree
}

// Sort : 进行排序
func Sort(values []int) {
	var root *tree
	// 对这些数据逐个添加排序
	for _, v := range values {
		root = add(root, v)
	}
	// 最后将已经排序完成的数据返回
	appendValues(values[:0], root)
}

// appendValues ： 将t的元素按顺序添加到tree，并返回已经排序的切片
func appendValues(values []int, t *tree) []int {
	if t != nil {
		// 先循环遍历左子树左子树
		values = appendValues(values, t.left)
		// 再将当前节点的数值返回，并添加到values中
		values = append(values, t.value)
		// 最后循环遍历右子树
		values = appendValues(values, t.right)
	}
	return values
}

// add ：将value添加到tree中
func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	// 需要添加的value小于当前value则向左子树，大于等于则向右子树
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
