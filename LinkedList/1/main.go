package main

//LeetCode 第 24 题: 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]

type ListNode struct {
	Value int
	Next  *ListNode
}

// 递归
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs1(newHead.Next)
	newHead.Next = head
	return newHead
}

// 迭代
func swapPairs2(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}

//LeetCode 第 25 题：给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。k 是一个正整数，它的值小于或等于链表的长度。
//如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
