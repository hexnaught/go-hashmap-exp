package linkedlist

import "fmt"

type LinkedList struct {
	Head *node
}

type node struct {
	key   string
	value interface{}
	next  *node
}

func (ll *LinkedList) Print() {
	currNode := ll.Head
	for currNode != nil {
		fmt.Printf("\t\t\t[%p]:\n\t\t\t\tKey: %+v \tValue: %+v\n", currNode, currNode.key, currNode.value)
		currNode = currNode.next
	}
}

func (ll *LinkedList) Search(key string) (bool, *node) {
	currNode := ll.Head
	for currNode != nil {
		if currNode.key == key {
			return true, currNode
		}
		currNode = currNode.next
	}
	return false, nil
}

func (ll *LinkedList) Insert(key string, value interface{}) {
	newNode := &node{
		key:   key,
		value: value,
		next:  ll.Head,
	}
	ll.Head = newNode
}

func (ll *LinkedList) Set(key string, value interface{}) {
	_, node := ll.Search(key)
	node.value = value
}

func (ll *LinkedList) Get(key string) interface{} {
	currNode := ll.Head
	for currNode != nil {
		if currNode.key == key {
			return currNode.value
		}
		currNode = currNode.next
	}
	return nil
}

func (ll *LinkedList) Remove(key string) {
	if ll.Head.key == key {
		ll.Head = ll.Head.next
		return
	}

	prevNode := ll.Head
	for prevNode != nil {
		if prevNode.next == nil {
			break
		}

		if prevNode.next.key == key {
			prevNode.next = prevNode.next.next
			return
		}
		prevNode = prevNode.next
	}
}
