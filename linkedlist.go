package main

import (
	"fmt"
)

type Node struct {
	value int
	next *Node
}

type LinkedList struct {
	len, iter int
	head *Node
}

/**
	Returns next value.
	Only moves an internal iterator, does alter underlying list.
	Can return nil when list is empty or when at the end of the list.
*/
func (ll *LinkedList) Next() *Node {
	if ll.len == 0 || ll.iter >= ll.len {
		return nil
	} else {
		for i, tmp := 0, ll.head; tmp != nil; tmp = tmp.next {
			if ll.iter == i {
				ll.iter++
				return tmp
			}
			i++
		}

		return nil
	}
}

/**
	Reset iterator to start of list.
*/
func (ll *LinkedList) Reset() {
	ll.iter = 0
}

/**
	Add node to head.
	Constant time operation O(1).
*/
func (ll *LinkedList) PushToHead(value int) {
	node := &Node{value, nil}
	ll.head, node.next = node, ll.head
	ll.len++
}

/**
	Add node to tail.
	Linear time operation O(n).
*/
func (ll *LinkedList) PushToTail(value int) {
	node := &Node{value, nil}
	for tmp := ll.head; tmp != nil; tmp = tmp.next {
		if tmp.next == nil {
			tmp.next = node
			ll.len++
			break
		}
	}	
}

/**
	Remove node from head.
	Constant time operation O(1).
*/
func (ll *LinkedList) PopFromHead() {
	if ll.len == 0 {
		fmt.Println("List is aleady empty.")
	} else if ll.len == 1 {
		ll.head = nil
		ll.len--
	} else {
		ll.head = ll.head.next
		ll.len--
	}
}

/**
	Remove node from tail.
	Linear time operation O(n).
*/
func (ll *LinkedList) PopFromTail() {
	if ll.len == 0 {
		fmt.Println("List is aleady empty.")
	} else if ll.len == 1 {
		ll.head = nil
		ll.len--
	} else {
		for tmp := ll.head; tmp != nil; tmp = tmp.next {
			if tmp.next.next == nil {
				tmp.next = nil
				ll.len--
				break;
			}
		}
	}
}

/**
	Returns length of the list.
*/
func (ll *LinkedList) Len() int {
	return ll.len
}

/**
	Prints visual representation of list.
	Truncates output after 8 nodes.
*/
func (ll *LinkedList) Print() {
	fmt.Print("[")
	for i, tmp := 1, ll.head; tmp != nil; tmp = tmp.next {
		if (i > 8) {
			fmt.Print("...")
			break;
		}

		fmt.Print(tmp.value)
		if tmp.next != nil {
			fmt.Print("->")
		}
		i++
	}
	fmt.Print("]\n")
}

func main() {
	ll := LinkedList{}
	ll.PushToHead(10)
	ll.PushToHead(30)
	ll.PushToTail(999)
	ll.PushToTail(1000)

	next := ll.Next()
	next = ll.Next()
	next = ll.Next()
	if (next != nil) {
		fmt.Println(next.value)
	}

	// fmt.Println(ll.Len())
	// ll.Print()
	// ll.PopFromHead()
	// ll.Print()
	// ll.PopFromHead()
	// ll.Print()
	// ll.PopFromTail()
	// ll.Print()
	// ll.PopFromHead()
	// ll.Print()
	// ll.PopFromTail()
	// ll.Print()
}