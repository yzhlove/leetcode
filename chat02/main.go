package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	//998 + 998 = 1996
	//
	//999 + 999 = 1998
	//
	//999
	//999
	//
	//8991

	a := build(342)
	b := build(465)
	show(a)
	show(b)

	c := addTwoNumbers(a, b)
	show(c)

	// 4321
	// 765
	// 1081
}

func show(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val)
		l = l.Next
	}
	fmt.Println()
}

func build(number int) *ListNode {

	var topNode, tempNode *ListNode
	codes := strconv.Itoa(number)
	for i := len(codes) - 1; i >= 0; i-- {
		var node = new(ListNode)
		n, _ := strconv.Atoi(string(codes[i]))
		node.Val = n
		fmt.Println("build n --> ", n)
		if topNode == nil {
			topNode = node
		}
		if tempNode != nil {
			tempNode.Next = node
		}
		tempNode = node
	}
	return topNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var a, b, total int
	var topNode, tempNode *ListNode
	var overflow bool

	for {
		a, b, total = 0, 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		var node = new(ListNode)
		total = a + b
		if overflow {
			total += 1
		}

		if total > 9 {
			node.Val = total - 10
			overflow = true
		} else {
			node.Val = total
			overflow = false
		}
		if topNode == nil {
			topNode = node
		}
		if tempNode != nil {
			tempNode.Next = node
		}
		tempNode = node
		if l1 == nil && l2 == nil {
			break
		}
	}

	if overflow {
		node := new(ListNode)
		node.Val = 1
		tempNode.Next = node
	}
	return topNode
}
