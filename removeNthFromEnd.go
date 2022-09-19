package main

import "fmt"

//题目：删除链表的倒数第n个节点
//leetcode链接：https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

type ListNode struct{
	val int
	Next *ListNode
}

//栈方法思路：借助slice保存链表的所有节点，从slice尾部开始计算，对倒数第n个节点进行删除
func removeNthNodeFromEndByStack(head *ListNode, n int) *ListNode{

	stack := make([]*ListNode, n)
	//头部插入一个空节点连接到head，目的在于方便处理删除头节点的情况
	free := &ListNode{Next: head}
	head = free
	for head != nil{
		stack = append(stack, head)
		head = head.Next
	}
	del := stack[len(stack)-n-1]
	del.Next = del.Next.Next

	return free.Next
}


//快慢指针法思路：设置快慢指针，快指针先走，等到快指针遍历了n个节点，再让慢指针开始走，这样快慢指针间隔就会有n个节点
//当快指针走链表尾部nil时，慢指针恰好走到倒数第n个节点的上一个节点，这样就可以删除倒数第n个节点了
func removeNthNodeFromEndByFastSlow(head *ListNode, n int) *ListNode{
	//头部插入一个空节点连接到head，目的在于方便处理删除头节点的情况
	free := &ListNode{Next: head}
	fast, slow := head, free

	for i := 0; i < n; i++{
		fast = fast.Next
	}

	for ;fast != nil; fast = fast.Next{
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return free.Next
}

func makeList(nums []int) *ListNode{
	head := &ListNode{}
	dst := head
	for _,num := range nums{
		node := &ListNode{
			val: num,
		}
		head.Next = node
		head = node
	}

	return dst.Next

}

func PrintList(head *ListNode){
	for head != nil{
		fmt.Printf("%d->", head.val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	nums := []int{1,2,3,4,5}
	list := makeList(nums)
	PrintList(list)
	list = removeNthNodeFromEndByStack(list,5)
	PrintList(list)
}
