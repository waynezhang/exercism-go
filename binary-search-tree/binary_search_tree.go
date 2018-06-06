package binarysearchtree

// SearchTreeData doc here
type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// Bst doc here
func Bst(data int) *SearchTreeData {
	return &SearchTreeData{data: data}
}

// Insert doc here
func (tree *SearchTreeData) Insert(data int) {
	if data <= tree.data {
		if tree.left == nil {
			tree.left = &SearchTreeData{data: data}
		} else {
			tree.left.Insert(data)
		}
	} else {
		if tree.right == nil {
			tree.right = &SearchTreeData{data: data}
		} else {
			tree.right.Insert(data)
		}
	}
}

// MapString doc here
func (tree *SearchTreeData) MapString(mapFunc func(int) string) []string {
	arr := make([]string, 0)
	if tree.left != nil {
		arr = append(arr, tree.left.MapString(mapFunc)...)
	}
	arr = append(arr, mapFunc(tree.data))
	if tree.right != nil {
		arr = append(arr, tree.right.MapString(mapFunc)...)
	}
	return arr
}

// MapInt doc here
func (tree *SearchTreeData) MapInt(mapFunc func(int) int) []int {
	arr := make([]int, 0)
	if tree.left != nil {
		arr = append(arr, tree.left.MapInt(mapFunc)...)
	}
	arr = append(arr, mapFunc(tree.data))
	if tree.right != nil {
		arr = append(arr, tree.right.MapInt(mapFunc)...)
	}
	return arr
}
