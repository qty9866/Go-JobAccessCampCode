package main

func postorder(root *treeNode, targetNum int) bool {
	if root == nil {
		return false
	}
	totalCompare++
	if root.data > targetNum {
		if Inorder(root.left, targetNum) {
			return true
		}
	}
	if root.data < targetNum {
		if Inorder(root.right, targetNum) {
			return true
		}
	}
	if root.data == targetNum {
		return true
	}
	return false
}
