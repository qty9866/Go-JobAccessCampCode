package main

import "fmt"

type LinkedNode struct {
	data     int
	next     *LinkedNode
	previous *LinkedNode
}

func main() {
	n1 := buildDLink()
	RangeLinkedList(n1)

	fmt.Println("------插入3/7/11/0------")
	insertInto(n1, &LinkedNode{data: 3})
	insertInto(n1, &LinkedNode{data: 7})
	insertInto(n1, &LinkedNode{data: 11})
	n1 = insertInto(n1, &LinkedNode{data: 0})

	RangeLinkedList(n1) // 0 1 3 5 7 10 11

	fmt.Println("-------删除11------")
	n1, _ = DeleteNode(n1, 11)
	RangeLinkedList(n1)

	n11 := &LinkedNode{data: 11}
	_, ok := DeleteNode(n11, 11)
	fmt.Println(ok)
}

func buildDLink() *LinkedNode {
	n1 := &LinkedNode{data: 1}
	n5 := &LinkedNode{data: 5}
	n10 := &LinkedNode{data: 10}

	n1.next = n5
	n1.previous = nil

	n5.next = n10
	n5.previous = n1

	n10.previous = n5
	return n1
}

func insertInto(head *LinkedNode, newNode *LinkedNode) *LinkedNode {
	tmpNode := head
	// 链表为空
	if head == nil {
		head = newNode
		return head
	}
	// 在头结点插入
	if head.data >= newNode.data {
		tmpNode.previous = newNode
		newNode.next = tmpNode
		return newNode
	}
	// 在后面插入
	for {
		if tmpNode.next == nil {
			//已经到达结尾，在最后添加节点
			tmpNode.next = newNode
			newNode.previous = tmpNode
			return head
		} else {
			if tmpNode.next.data >= newNode.data {
				//	TODO : 找到了合适的位置，开始插入
				newNode.next = tmpNode.next
				tmpNode.next.previous = newNode
				newNode.previous = tmpNode
				tmpNode.next = newNode
				return head
			}
		}
		tmpNode = tmpNode.next
	}
}

func RangeLinkedList(head *LinkedNode) {
	fmt.Println("从头读到尾")
	tmpNode := head
	for {
		fmt.Println(tmpNode.data)
		if tmpNode.next == nil {
			break
		}
		tmpNode = tmpNode.next
	}

	fmt.Println("从尾到头读")
	for {
		fmt.Println(tmpNode.data)
		if tmpNode.previous == nil {
			break
		}
		tmpNode = tmpNode.previous
	}
}

// DeleteNode 删除指定Node
func DeleteNode(head *LinkedNode, value int) (*LinkedNode, bool) {
	tmpNode := head
	// 空链表
	if head == nil {
		return nil, false
	}
	// 删除的是头结点
	if head.data == value {
		if head.next == nil {
			return nil, true
		} else {
			tmpNode = head.next
			head.next.previous = nil
			head.next = nil
			return tmpNode, true
		}

	}
	for {
		if tmpNode.next == nil {
			return head, false
		} else {
			if tmpNode.next.data == value {
				right := tmpNode.next
				if right.next == nil {
					tmpNode.next = nil
					right.previous = nil
					return head, true
				} else {
					right.next.previous = tmpNode
					tmpNode.next = right.next
					right.previous = nil
					right.next = nil
					return head, true
				}

			}
		}
		tmpNode = tmpNode.next
	}
}
