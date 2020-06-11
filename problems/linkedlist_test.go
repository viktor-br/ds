package problems

import "testing"

func TestIntersection(t *testing.T) {
	node4 := NewLinkedListNode(1, nil)
	node3 := NewLinkedListNode(1, node4)
	node2 := NewLinkedListNode(1, node3)
	node1 := NewLinkedListNode(1, node2)
	node22 := NewLinkedListNode(1, node1)
	node21 := NewLinkedListNode(1, node22)
	node14 := NewLinkedListNode(1, node1)
	node13 := NewLinkedListNode(1, node14)
	node12 := NewLinkedListNode(1, node13)
	node11 := NewLinkedListNode(1, node12)

	if !Intersected(node11, node21) {
		t.Errorf("expect true, actual false")
	}
}

func TestNoIntersection(t *testing.T) {
	node23 := NewLinkedListNode(1, nil)
	node22 := NewLinkedListNode(1, node23)
	node21 := NewLinkedListNode(1, node22)
	node14 := NewLinkedListNode(1, nil)
	node13 := NewLinkedListNode(1, node14)
	node12 := NewLinkedListNode(1, node13)
	node11 := NewLinkedListNode(1, node12)

	if Intersected(node11, node21) {
		t.Errorf("expect false, actual true")
	}
}

func TestPartition(t *testing.T) {
	node := CreateLinkedListFromArray([]int{3, 5, 8, 5, 10, 2, 1})
	expectedNode := CreateLinkedListFromArray([]int{1, 2, 3, 5, 8, 5, 10})

	actualNode := Partition(node, 5)

	if !LinkedListEquals(expectedNode, actualNode) {
		t.Errorf("expect: %s, actual: %s", LinkedListToString(expectedNode), LinkedListToString(actualNode))
	}
}

func TestSumLists(t *testing.T) {
	testCases := []struct {
		Name     string
		Node1    *LinkedListNode
		Node2    *LinkedListNode
		Expected *LinkedListNode
	}{
		{"sum list, same length", CreateLinkedListFromArray([]int{7, 1, 6}), CreateLinkedListFromArray([]int{5, 9, 2}), CreateLinkedListFromArray([]int{2, 1, 9})},
		{"sum list, one is longer", CreateLinkedListFromArray([]int{7, 1, 6}), CreateLinkedListFromArray([]int{5, 9, 2, 1}), CreateLinkedListFromArray([]int{2, 1, 9, 1})},
		{"sum list, additional shift", CreateLinkedListFromArray([]int{7, 0, 9}), CreateLinkedListFromArray([]int{3, 0, 2, 9}), CreateLinkedListFromArray([]int{0, 1, 1, 0, 1})},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actualNode := SumLists(tc.Node1, tc.Node2)

			if !LinkedListEquals(tc.Expected, actualNode) {
				t.Errorf("expect: %s, actual: %s", LinkedListToString(tc.Expected), LinkedListToString(actualNode))
			}
		})
	}
}

func TestSumListsReverse(t *testing.T) {
	testCases := []struct {
		Name     string
		Node1    *LinkedListNode
		Node2    *LinkedListNode
		Expected *LinkedListNode
	}{
		{"sum list, same length", CreateLinkedListFromArray([]int{6, 1, 7}), CreateLinkedListFromArray([]int{2, 9, 5}), CreateLinkedListFromArray([]int{9, 1, 2})},
		{"sum list, one is longer", CreateLinkedListFromArray([]int{6, 1, 7}), CreateLinkedListFromArray([]int{1, 2, 9, 5}), CreateLinkedListFromArray([]int{1, 9, 1, 2})},
		{"sum list, additional shift", CreateLinkedListFromArray([]int{9, 0, 7}), CreateLinkedListFromArray([]int{9, 2, 0, 3}), CreateLinkedListFromArray([]int{1, 0, 1, 1, 0})},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actualNode := SumListsReverse(tc.Node1, tc.Node2)

			if !LinkedListEquals(tc.Expected, actualNode) {
				t.Errorf("expect: %s, actual: %s", LinkedListToString(tc.Expected), LinkedListToString(actualNode))
			}
		})
	}
}

func TestAddZeroNodes(t *testing.T) {
	node := CreateLinkedListFromArray([]int{6, 1, 7})
	actual := AddZeroNodes(node, 3)
	expected := CreateLinkedListFromArray([]int{0, 0, 0, 6, 1, 7})
	if !LinkedListEquals(expected, actual) {
		t.Errorf("expect: %s, actual: %s", LinkedListToString(expected), LinkedListToString(actual))
	}
}
