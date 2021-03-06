package ds

import "testing"

func TestLinkedList(t *testing.T) {
	linkedList := CreateLinkedListFromArray([]int{5, 1, 4, 2, 5, 2, 2, 2, 3, 1})
	expectedLinkedList := CreateLinkedListFromArray([]int{5, 1, 4, 2, 3})

	linkedList.RemoveDuplicatesNoBuf()

	if !LinkedListEquals(linkedList, expectedLinkedList) {
		t.Errorf("Expect %s, actual %s", expectedLinkedList.Sprint(), linkedList.Sprint())
	}
}

func TestEmptyLinkedList(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.RemoveDuplicates()
}

func TestLinkedListRemoveMiddleElementEven(t *testing.T) {
	linkedList := CreateLinkedListFromArray([]int{5, 1, 4, 2, 5, 2, 2, 2, 3, 1})
	expectedLinkedList := CreateLinkedListFromArray([]int{5, 1, 4, 2, 2, 2, 2, 3, 1})

	linkedList.RemoveMiddleElement()

	if !LinkedListEquals(linkedList, expectedLinkedList) {
		t.Errorf("Expect %s, actual %s", expectedLinkedList.Sprint(), linkedList.Sprint())
	}
}

func TestLinkedListRemoveMiddleElementOdd(t *testing.T) {
	linkedList := CreateLinkedListFromArray([]int{5, 1, 4, 2, 5, 2, 2, 3, 1})
	expectedLinkedList := CreateLinkedListFromArray([]int{5, 1, 4, 2, 2, 2, 3, 1})

	linkedList.RemoveMiddleElement()

	if !LinkedListEquals(linkedList, expectedLinkedList) {
		t.Errorf("Expect %s, actual %s", expectedLinkedList.Sprint(), linkedList.Sprint())
	}
}

func TestLinkedListRemoveMiddleElementTwo(t *testing.T) {
	linkedList := CreateLinkedListFromArray([]int{5, 6})
	expectedLinkedList := CreateLinkedListFromArray([]int{6})

	linkedList.RemoveMiddleElement()

	if !LinkedListEquals(linkedList, expectedLinkedList) {
		t.Errorf("Expect %s, actual %s", expectedLinkedList.Sprint(), linkedList.Sprint())
	}
}

func TestLinkedListIsPalindrome(t *testing.T) {
	linkedListAsArray := []int{5, 6, 7, 8, 7, 6, 5}
	linkedList := CreateLinkedListFromArray(linkedListAsArray)

	if !linkedList.IsPalindrome() {
		t.Errorf("Expect %v is palindrome", linkedListAsArray)
	}

	linkedListAsArray = []int{5, 6, 7, 7, 6, 5}
	linkedList = CreateLinkedListFromArray(linkedListAsArray)

	if !linkedList.IsPalindrome() {
		t.Errorf("Expect %v is palindrome", linkedListAsArray)
	}

	linkedListAsArray = []int{5, 7, 7, 6, 5}
	linkedList = CreateLinkedListFromArray(linkedListAsArray)

	if linkedList.IsPalindrome() {
		t.Errorf("Expect %v is palindrome", linkedListAsArray)
	}

	linkedListAsArray = []int{5}
	linkedList = CreateLinkedListFromArray(linkedListAsArray)

	if !linkedList.IsPalindrome() {
		t.Errorf("Expect %v is palindrome", linkedListAsArray)
	}
}
