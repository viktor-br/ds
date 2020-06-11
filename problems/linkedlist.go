package problems

import (
	"bytes"
	"fmt"
)

type LinkedListNode struct {
	Data int
	Next *LinkedListNode
}

func NewLinkedListNode(data int, nextNode *LinkedListNode) *LinkedListNode {
	return &LinkedListNode{data, nextNode}
}

func CreateLinkedListFromArray(elements []int) *LinkedListNode {
	if len(elements) == 0 {
		return nil
	}
	var node *LinkedListNode
	for i := len(elements) - 1; i >= 0; i-- {
		node = NewLinkedListNode(elements[i], node)
	}

	return node
}

func LinkedListLength(node *LinkedListNode) int {
	i := 0
	current := node

	for current != nil {
		current = current.Next
		i++
	}

	return i
}

func LinkedListEquals(node1 *LinkedListNode, node2 *LinkedListNode) bool {
	current1 := node1
	current2 := node2
	for current1 != nil {
		if current2 == nil {
			return false
		}

		if current1.Data != current2.Data {
			return false
		}
		current1 = current1.Next
		current2 = current2.Next
	}

	if current2 != nil {
		return false
	}

	return true
}

func LinkedListToString(node *LinkedListNode) string {
	if node == nil {
		return fmt.Sprint("LinkedList is empty")
	}
	var buffer bytes.Buffer
	element := node
	for element != nil {
		buffer.WriteString(fmt.Sprintf(" -> %d", element.Data))
		element = element.Next
	}

	return buffer.String()
}

func Partition(node *LinkedListNode, x int) *LinkedListNode {
	head := node
	tail := node
	var next *LinkedListNode

	for node != nil {
		next = node.Next
		if node.Data < x {
			node.Next = head
			head = node
		} else {
			tail.Next = node
			tail = node
		}
		node = next
	}

	tail.Next = nil

	return head
}

func Intersected(node1 *LinkedListNode, node2 *LinkedListNode) bool {
	len1 := LinkedListLength(node1)
	len2 := LinkedListLength(node2)

	p1 := node1
	p2 := node2

	if len1 > len2 {
		p1 = MoveToKth(p1, len1-len2)
	} else if len1 < len2 {
		p2 = MoveToKth(p1, len2-len1)
	}

	for p1 != nil {
		if p1 == p2 {
			return true
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return false
}

func MoveToKth(node *LinkedListNode, k int) *LinkedListNode {
	current := node
	for i := 0; i < k; i++ {
		current = current.Next
	}

	return current
}

func SumLists(node1 *LinkedListNode, node2 *LinkedListNode) *LinkedListNode {
	head := NewLinkedListNode(-1, nil)
	tail := head
	var current *LinkedListNode
	current1 := node1
	current2 := node2
	var data int = 0
	var data1 int = current1.Data
	var data2 int = current2.Data
	for current1 != nil || current2 != nil || data > 0 {
		data = data1 + data2 + data
		current = NewLinkedListNode(data%10, nil)
		tail.Next = current
		tail = current

		if data > 9 {
			data = 1
		} else {
			data = 0
		}

		if current1 != nil {
			current1 = current1.Next
		}

		if current2 != nil {
			current2 = current2.Next
		}

		if current1 == nil {
			data1 = 0
		} else {
			data1 = current1.Data
		}

		if current2 == nil {
			data2 = 0
		} else {
			data2 = current2.Data
		}
	}

	return head.Next
}

func SumListsReverse(node1 *LinkedListNode, node2 *LinkedListNode) *LinkedListNode {
	len1 := LinkedListLength(node1)
	len2 := LinkedListLength(node2)

	current1 := node1
	current2 := node2

	if len1 > len2 {
		current2 = AddZeroNodes(current2, len1-len2)
	} else if len2 > len1 {
		current1 = AddZeroNodes(current1, len2-len1)
	}

	node, data := createNode(current1, current2)
	if data > 0 {
		node = NewLinkedListNode(data, node)
	}

	return node
}

func createNode(node1 *LinkedListNode, node2 *LinkedListNode) (*LinkedListNode, int) {
	if node1 == nil && node2 == nil {
		return nil, 0
	}

	current, data := createNode(node1.Next, node2.Next)

	data = node1.Data + node2.Data + data

	node := NewLinkedListNode(data%10, current)

	if data > 9 {
		data = 1
	} else {
		data = 0
	}

	return node, data
}

func AddZeroNodes(node *LinkedListNode, k int) *LinkedListNode {
	current := NewLinkedListNode(0, node)
	for i := 1; i < k; i++ {
		current = NewLinkedListNode(0, current)
	}

	return current
}
