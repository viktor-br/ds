package ds

import "fmt"
import "bytes"

// LinkedList represents LL itself
type LinkedList struct {
	head  *LinkedListElement
	count int
}

// LinkedListElement represent element of LL
type LinkedListElement struct {
	data int
	next *LinkedListElement
}

// NewLinkedList creates LinkedList
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// NewLinkedListElement creates LinkedListElement
func NewLinkedListElement(value int) *LinkedListElement {
	return &LinkedListElement{data: value}
}

// CreateLinkedListFromArray obvious
func CreateLinkedListFromArray(elements []int) *LinkedList {
	linkedList := NewLinkedList()
	for i := len(elements) - 1; i >= 0; i-- {
		linkedList.Add(NewLinkedListElement(elements[i]))
	}

	return linkedList
}

// LinkedListEquals check if two linked lists are equal
func LinkedListEquals(linkedList1, linkedList2 *LinkedList) bool {
	currentElement1 := linkedList1.head
	currentElement2 := linkedList2.head
	if linkedList1.count != linkedList2.count {
		return false
	}

	for currentElement1 != nil && currentElement2 != nil {
		if currentElement1.data != currentElement2.data {
			return false
		}
		currentElement1 = currentElement1.next
		currentElement2 = currentElement2.next
	}

	return true
}

// Add element to a linked list.
func (linkedList *LinkedList) Add(element *LinkedListElement) {
	if linkedList.head != nil {
		element.next = linkedList.head
	}

	linkedList.head = element
	linkedList.count++
}

// Remove element from top
func (linkedList *LinkedList) Remove() {
	if linkedList.count == 0 {
		return
	}
	linkedList.head = linkedList.head.next
	linkedList.count--
}

// Print displays linked list elements.
func (linkedList *LinkedList) Print() {
	fmt.Println(linkedList.Sprint())
}

// Sprint return string representation of the linked list
func (linkedList *LinkedList) Sprint() string {
	if linkedList.head == nil {
		return fmt.Sprintf("LinkedList is empty")
	}
	var buffer bytes.Buffer
	element := linkedList.head
	buffer.WriteString(fmt.Sprintf("(count=%d)", linkedList.count))
	for element != nil {
		buffer.WriteString(fmt.Sprintf(" -> %d", element.data))
		element = element.next
	}

	return buffer.String()
}

// Count linked list elements.
func (linkedList *LinkedList) Count() int {
	return linkedList.count
}

// RemoveDuplicates removes duplicates
func (linkedList *LinkedList) RemoveDuplicates() {
	buffer := map[int]bool{}

	currentElement := linkedList.head
	var prevElement *LinkedListElement

	for currentElement != nil {
		_, ok := buffer[currentElement.data]
		if ok {
			prevElement.next = currentElement.next
			linkedList.count--
		} else {
			prevElement = currentElement
			buffer[currentElement.data] = false
		}

		currentElement = currentElement.next
	}
}

// RemoveDuplicatesNoBuf removes duplicates
func (linkedList *LinkedList) RemoveDuplicatesNoBuf() {
	currentElement := linkedList.head
	var nextElement *LinkedListElement
	var prevElement *LinkedListElement

	for currentElement != nil {
		prevElement = currentElement
		nextElement = currentElement.next

		for nextElement != nil {
			if currentElement.data == nextElement.data {
				nextElement = nextElement.next
				prevElement.next = nil
				linkedList.count--
				continue
			}
			prevElement.next = nextElement
			prevElement = nextElement
			nextElement = nextElement.next
		}

		currentElement = currentElement.next
	}
}

// RemoveMiddleElement removes middle element of the linked list
func (linkedList *LinkedList) RemoveMiddleElement() {
	currentIndex := 0
	var middleElementIndex int
	currentMiddleIndex := 0
	middleElementParent := linkedList.head
	currentElement := linkedList.head

	if linkedList.count <= 2 {
		linkedList.Remove()
		return
	}

	for currentElement != nil {
		middleElementIndex = currentIndex / 2
		if middleElementIndex-1 > currentMiddleIndex {
			middleElementParent = middleElementParent.next
			currentMiddleIndex++
		}
		currentIndex++
		currentElement = currentElement.next
	}

	middleElementParent.next = middleElementParent.next.next
	linkedList.count--
}

// IsPalindrome checks if linked list is the same if are reading in both directions
func (linkedList *LinkedList) IsPalindrome() bool {
	_, isPalindrome := linkedList.palindrom(linkedList.head)

	return isPalindrome
}

func (linkedList *LinkedList) palindrom(linkedListElement *LinkedListElement) (*LinkedListElement, bool) {
	if linkedListElement == nil {
		return linkedList.head, true
	}

	linkedListElementFromHead, isPalindrome := linkedList.palindrom(linkedListElement.next)
	if !isPalindrome {
		return linkedListElementFromHead.next, isPalindrome
	}

	if linkedListElementFromHead.data != linkedListElement.data {
		return linkedListElementFromHead.next, false
	}

	return linkedListElementFromHead.next, true
}
