package main

// εεΊιε
func preOrder(root *treeNode, targetNum int) bool {
	if root == nil {
		return false
	}
	totalCompare++
	if root.data == targetNum {
		return true
	}
	if preOrder(root.left, targetNum) {
		return true
	}
	if preOrder(root.right, targetNum) {
		return true
	}
	return false
}
