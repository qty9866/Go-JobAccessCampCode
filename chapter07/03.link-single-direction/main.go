package main

import (
	"fmt"
)

// 定义单链表节点结构体
type linkedNode struct {
	data int
	next *linkedNode
}

func main() {
	n1 := &linkedNode{
		data: 1,
		next: nil,
	}
	n2 := &linkedNode{
		data: 2,
		next: nil,
	}
	n3 := &linkedNode{
		data: 3,
		next: nil,
	}
	n4 := &linkedNode{
		data: 4,
		next: nil,
	}
	n6 := &linkedNode{
		data: 6,
		next: nil,
	}
	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n6

	n5 := &linkedNode{
		data: 5,
		next: nil,
	}

	n7 := &linkedNode{
		data: 7,
		next: nil,
	}
	fmt.Println("插入5:-----------")
	// 添加新链表
	insertNode(n1, n5)
	insertNode(n1, n7)
	n1 = deleteNode(n1, 1)

	RangeLinkedList(n1)
}

// insertNode 插入节点
func insertNode(head *linkedNode, newNode *linkedNode) {
	tmpNode := head
	for {
		if tmpNode != nil {
			if newNode.data > tmpNode.data {
				if tmpNode.next == nil {
					// 已经到了结尾，直接追加即可
					tmpNode.next = newNode
				} else {
					if tmpNode.next.data >= newNode.data {
						//	找到合适位置，准备插入数据
						newNode.next = tmpNode.next
						tmpNode.next = newNode
						break
					}
				}
			}
		} else {
			break
		}
		tmpNode = tmpNode.next
	}
}

// RangeLinkedList 遍历链表
func RangeLinkedList(head *linkedNode) {
	tmpNode := head
	for {
		if tmpNode != nil {
			fmt.Println(tmpNode.data)
			tmpNode = tmpNode.next
		} else {
			break
		}
	}
}

// 删除节点
func deleteNode(head *linkedNode, data int) *linkedNode {
	tmpNode := head
	if head != nil && head.data == data {
		if head.next == nil {
			return nil
		}
		right := head.next
		tmpNode.next = right
		return right
	}
	for {
		if tmpNode.next == nil {
			break
		}
		right := tmpNode.next
		if right.data == data {
			//	找到要删除的节点，开始删除
			tmpNode.next = right.next
			right.next = nil // 内存回收
			return head
		}
		tmpNode = tmpNode.next
	}
	return head
}
