package datastructures

import (
	"testing"

	"github.com/arrwhidev/testutils-golang"
)

func Test_LinkedList_shouldSetInitialLengthCorrectly(t *testing.T) {
	ll := NewLinkedList(99)
	utils.AssertEqual(t, 1, ll.Len())
}

func Test_LinkedList_Len_shouldReturnZeroForInitialLength(t *testing.T) {
	ll := LinkedList{}
	utils.AssertEqual(t, 0, ll.Len())
}

func Test_LinkedList_Next_shouldReturnNextNode_andIncrementIter_andNotRemoveNodesFromList(t *testing.T) {
	ll := LinkedList{}
	ll.PushToHead(10)
	ll.PushToHead(20)
	utils.AssertEqual(t, 2, ll.Len())

	next := ll.Next()
	utils.AssertEqual(t, 20, next.value)
	utils.AssertEqual(t, 2, ll.Len())

	next = ll.Next()
	utils.AssertEqual(t, 10, next.value)
	utils.AssertEqual(t, 2, ll.Len())
}

func Test_LinkedList_Reset_shouldResetIter(t *testing.T) {
	ll := LinkedList{}
	ll.PushToHead(10)
	ll.PushToHead(20)

	ll.Next()
	utils.AssertEqual(t, 1, ll.iter)

	ll.Next()
	utils.AssertEqual(t, 2, ll.iter)

	ll.Reset()
	utils.AssertEqual(t, 0, ll.iter)
}

func Test_LinkedList_PushToHead_shouldAppendNodeToFrontOfList(t *testing.T) {
	ll := NewLinkedList(99)
	utils.AssertEqual(t, 99, ll.head.value)

	ll.PushToHead(10)
	utils.AssertEqual(t, 10, ll.head.value)
	utils.AssertEqual(t, 2, ll.Len())
}

func Test_LinkedList_PushToTail_shouldAppendNodeToTailOfList(t *testing.T) {
	ll := NewLinkedList(99)
	utils.AssertEqual(t, 99, ll.head.value)

	ll.PushToTail(10)
	utils.AssertEqual(t, 99, ll.head.value)
	utils.AssertEqual(t, 10, ll.head.next.value)
	utils.AssertEqual(t, 2, ll.Len())
}

func Test_LinkedList_PopFromHead_shouldRemoveNodeFromHeadOfList(t *testing.T) {
	ll := LinkedList{}
	ll.PushToHead(50)
	ll.PushToHead(100)

	ll.PopFromHead()
	utils.AssertEqual(t, 50, ll.head.value)
	utils.AssertEqual(t, 1, ll.Len())
}

func Test_LinkedList_PopFromTail_shouldRemoveNodeFromTailOfList(t *testing.T) {
	ll := LinkedList{}
	ll.PushToHead(50)
	ll.PushToHead(100)

	ll.PopFromTail()
	utils.AssertEqual(t, 100, ll.head.value)
	utils.AssertEqual(t, 1, ll.Len())
}

func Test_LinkedList_Print_shouldReturnStringRepresentationOfList(t *testing.T) {
	ll := LinkedList{}
	ll.PushToHead(50)
	ll.PushToHead(100)
	ll.PushToTail(999)
	ll.PushToTail(1000)

	utils.AssertEqual(t, "[100->50->999->1000]", ll.Print())
}
