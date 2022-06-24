package main

import "fmt"

func main() {
	var l1 = &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	var l2 = &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	var result = addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Print(result.Val)
		result = result.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	//var node ListNode 声明了一个类型为ListNode的变量node,node指向结构体里面都是对应类型的默认值，只有指针才可能为null
	var result *ListNode
	var nextNode *ListNode
	var carry = 0
	for l1 != nil && l2 != nil {
		nextNode = generateAddedNode(l1.Val, l2.Val, &carry, nextNode)
		if result == nil {
			result = nextNode
		}
		// 指针后移
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		nextNode = generateAddedNode(l1.Val, 0, &carry, nextNode)
		if result == nil {
			result = nextNode
		}
		// 指针后移
		l1 = l1.Next
	}
	for l2 != nil {
		nextNode = generateAddedNode(0, l2.Val, &carry, nextNode)
		if result == nil {
			result = nextNode
		}
		l2 = l2.Next
	}
	if carry != 0 && nextNode != nil {
		nextNode.Next = &ListNode{
			carry, nil,
		}
	}
	return result
}

func generateAddedNode(num1 int, num2 int, carry *int, nextNode *ListNode) *ListNode {
	sum := addTowNumber(num1, num2, carry)
	if nextNode == nil {
		nextNode = &ListNode{sum, nil}
	} else {
		nextNode.Next = &ListNode{sum, nil}
		nextNode = nextNode.Next
	}
	return nextNode
}

// 多返回值需要在后面定义 (ListNode, int)
// 返回新增加的ListNode和进位
func addTowNumber(num1 int, num2 int, carry *int) int {
	var result = 0
	var sum = num1 + num2 + *carry
	if sum >= 10 {
		result = sum % 10
		*carry = 1
	} else {
		result = sum
		*carry = 0
	}
	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}
