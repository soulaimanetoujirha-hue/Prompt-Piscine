package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func RemoveValue(l *Node, val int) *Node {
	if l == nil {
		return nil
	}

	if l.Val == val {
		return l.Next
	}
	curr := l
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
			return l
		}
		curr = curr.Next
	}

	return l
}

func printList(l *Node) {
	for n := l; n != nil; n = n.Next {
		fmt.Printf("%d ", n.Val)
	}
	fmt.Println()
}

func main() {
	list := &Node{1, &Node{2, &Node{3, &Node{4, nil}}}}

	list = RemoveValue(list, 3)
	printList(list)
}
