package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//1 2 4
//4 5 6

func main() {
	// node1 := &ListNode{
	// 	Val: 1,
	// }

	// node2 := &ListNode{
	// 	Val: 2,
	// }

	// node3 := &ListNode{
	// 	Val: 3,
	// }

	// fmt.Printf("head %p\n", head)

	// // head = node1
	// // prev = head
	// appendToResult(node1)
	// fmt.Printf("head %p\n", head)
	// fmt.Printf("node1 %p\n", node1)

	// fmt.Printf("node2 %p\n", node2)
	// // prev.Next = node2
	// // prev = node2
	// appendToResult(node2)
	// fmt.Printf("node2 %p\n", node2)

	// fmt.Printf("node3 %p\n", node3)
	// // prev.Next = node3
	// // prev = node3
	// appendToResult(node3)
	// fmt.Printf("node3 %p\n", node3)

	// fmt.Printf("head %p\n", head)
	// fmt.Println("head", head)
	// fmt.Println("head", head.Next)
	// fmt.Println("head", head.Next.Next)
	// fmt.Println("------------------------------")
	// return
	list1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	list2 := ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
			},
		},
	}

	fmt.Println(mergeTwoLists(&list1, &list2))
}

var prev *ListNode
var head *ListNode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	prev = nil
	head = nil
	i := 0
	for list1 != nil || list2 != nil {
		i++
		if list1 == nil {
			appendToResult(list2)
			list2 = list2.Next
			continue
		}
		if list2 == nil {
			appendToResult(list1)
			list1 = list1.Next
			continue
		}
		if list1.Val == list2.Val {
			appendToResult(list1)
			appendToResult(list2)
			list1 = list1.Next
			list2 = list2.Next
			continue
		}
		if list1.Val < list2.Val {
			appendToResult(list1)
			list1 = list1.Next
			continue
		}
		if list1.Val > list2.Val {
			appendToResult(list2)
			list2 = list2.Next
			continue
		}

	}

	return head
}

func appendToResult(list *ListNode) {
	if head == nil {
		head = new(ListNode)
		head.Val = list.Val
		prev = head
		return
	}

	tmp := new(ListNode)
	*tmp = *list
	prev.Next = tmp
	prev = tmp
}
