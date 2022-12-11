package main

func Inorder(root *treeNode, targetNum int) bool {
	if root == nil {
		return false
	}
	totalCompare++
	if Inorder(root.left, targetNum) {
		return true
	}
	if root.data == targetNum {
		return true
	}
	if Inorder(root.right, targetNum) {
		return true
	}
	return false
}
