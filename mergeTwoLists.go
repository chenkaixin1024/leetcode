package main

import "fmt"

//题目：合并两个有序链表
//leetcode链接：https://leetcode.cn/problems/merge-two-sorted-lists/


type ListNode struct{
	Val int
	Next *ListNode
}


//递归思路：取较小的返回，若遇到nil则返回另一个链表即可
func mergeTwoOrderListsByR(list1 *ListNode, list2 *ListNode) *ListNode{

	if list1 == nil{
		return list2
	}
	if list2 == nil{
		return list1
	}

	if list1.Val <= list2.Val{
		list1.Next = mergeTwoOrderListsByR(list1.Next, list2)
		return list1
	}else{
		list2.Next = mergeTwoOrderListsByR(list1, list2.Next)
		return list2
	}

}

//双指针迭代思路：l1,l2分别从list1和list2的头部开始，进行比较，较小的结点接入目标链表的next，且对应指针向后移，
//直到l1或l2遍历到对应链表末尾，则剩下的就是将另一方接入dst即可
func mergeTwoOrderLists(list1 *ListNode, list2 *ListNode) *ListNode{

	//链表引入空闲头结点，避免考虑实际插入时头结点问题
	dst := &ListNode{}
	l := dst
	l1, l2 := list1, list2
	for l1 !=nil && l2 != nil{
		if l1.Val <= l2.Val{
			l.Next = l1
			l1 = l1.Next
		}else{
			l.Next = l2
			l2 = l2.Next
		}
		l = l.Next
	}

	if l1 == nil{
		l.Next = l2
	}else{
		l.Next = l1
	}

	return dst.Next
}


func makeList(nums []int) *ListNode{
	head := &ListNode{}
	dst := head
	for _,num := range nums{
		node := &ListNode{
			Val: num,
		}
		head.Next = node
		head = node
	}

	return dst.Next

}

func PrintList(head *ListNode){
	for head != nil{
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {

	list1 := makeList([]int{1,2,5})
	list2 := makeList([]int{3,7,9})
	l := mergeTwoOrderListsByR(list1,list2)
	PrintList(l)

}
