package main

import (
	"fmt"
	"time"
)

type treeNode struct {
	data  int
	root  *treeNode
	left  *treeNode
	right *treeNode
}

// 简单构建一个树
func buildTree() *treeNode {
	n1 := &treeNode{data: 51}
	n2 := &treeNode{data: 35}
	n3 := &treeNode{data: 65}

	n1.left = n2
	n2.root = n1
	n1.right = n3
	n3.root = n1
	return n1
}

// 插入数据
func insertNode(root *treeNode, newNode *treeNode) *treeNode {
	if root == nil {
		return newNode
	}
	if newNode.data == root.data {
		return root
	}
	if newNode.data < root.data {
		if root.left == nil {
			//	找到位置开始插入
			root.left = newNode
			newNode.root = root
		} else {
			insertNode(root.left, newNode)
		}
	} else {
		if root.right == nil {
			root.right = newNode
			newNode.root = root
		} else {
			insertNode(root.right, newNode)
		}
	}
	return root
}

// 帮助理解如何删除叶子节点
func deleteNodeLeaf(leaf *treeNode, value int) *treeNode {
	refRoot := leaf
	if refRoot.data == value && refRoot.left == nil && refRoot.right == nil {
		refRoot = refRoot.root
		tmpNode := leaf
		if refRoot.left == tmpNode {
			// 删除左边叶子
			refRoot.left = nil
			tmpNode.root = nil
			return refRoot
		} else {
			// 删除右边叶子
			refRoot.right = nil
			tmpNode.root = nil
			return refRoot
		}
	}
	return refRoot
}

// 找到左边的后继
func findNextGenFromLeft(root *treeNode) *treeNode {
	if root == nil {
		return nil
	}
	tmpNode := root
	for {
		if tmpNode.right != nil {
			tmpNode = tmpNode.right
		} else {
			break
		}
	}
	return tmpNode
}

// 找到右边的后继
func findNextGenFromRight(root *treeNode) *treeNode {
	if root == nil {
		return nil
	}
	tmpNode := root
	for {
		if tmpNode.left != nil {
			tmpNode = tmpNode.left
		} else {
			break
		}
	}
	return tmpNode
}

// 删除节点
func deleteNode(root *treeNode, value int) *treeNode {
	if value < root.data {
		deleteNode(root.left, value)
	} else if value > root.data {
		deleteNode(root.right, value)
	} else {
		//现在root指向要删除的节点
		leftGen := findNextGenFromLeft(root.left)
		rightGen := findNextGenFromRight(root.right)
		if leftGen == nil && rightGen == nil {
			// 现在要删除的是叶子节点
			top := root.root
			down := root
			if top.left == down {
				//表示是左子树
				top.left = nil
				down.root = nil
				return nil
			} else {
				//表示是右子树
				top.right = nil
				down.root = nil
				return nil
			}
		} else if leftGen != nil {
			root.data = leftGen.data
			deleteNode(leftGen, leftGen.data)
		} else if rightGen != nil {
			root.data = rightGen.data
			deleteNode(rightGen, rightGen.data)
		}

	}
	return root
}

var totalCompare = 0

/*
preorder 前序遍历
总比较次数 3850000000
总共耗时 9.0590426s

inorder 中序遍历
总比较次数 3850000000
总共耗时 9.6180665s

postorder(不比较) 后续遍历
总比较次数 3850000000
总共耗时 9.6360882s

postorder(比较) 后续遍历
总比较次数 2478000000
总共耗时 6.3568514s
*/
func main() {
	var root *treeNode
	start := time.Now()
	for _, value := range sampleData {
		root = insertNode(root, &treeNode{data: int(value)})
	}
	//var arr = []int{4, 2, 6, 3, 1, 5, 7}
	//for _, value := range arr {
	//	root = insertNode(root, &treeNode{data: value})
	//}
	for i := 0; i < 2000000; i++ {
		postorder(root, 501)
		postorder(root, 888)
		postorder(root, 900)
		postorder(root, 3)
	}
	fmt.Println("总比较次数", totalCompare)
	finish := time.Now()
	fmt.Println("总共耗时", finish.Sub(start))
}
