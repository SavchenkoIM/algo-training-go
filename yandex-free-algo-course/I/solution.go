package main

/** Comment it before submitting
type Node struct {
	data string
	Next *Node
}
**/

func Reverse(head *Node, left int, right int) *Node {
	realHead := head

	lastHead := (*Node)(nil)
	firstTail := (*Node)(nil)

	var nodes []*Node
	for i := 1; i <= right+1; i++ {
		if i == left-1 {
			lastHead = head
		}
		if i == right+1 {
			firstTail = head
			break
		}

		if i >= left && i <= right {
			nodes = append(nodes, head)
		}
		if head != nil {
			head = head.Next
		} else {
			break
		}
	}
	for i := len(nodes) - 1; i > 0; i-- {
		nodes[i].Next = nodes[i-1]
	}

	if lastHead != nil {
		lastHead.Next = nodes[len(nodes)-1]
	} else {
		realHead = nodes[len(nodes)-1]
	}

	if firstTail != nil {
		nodes[0].Next = firstTail
	} else {
		nodes[0].Next = nil
	}

	return realHead
}
