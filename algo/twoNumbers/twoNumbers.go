package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//1000000000000000000000000000001
//  1 0 0 0 0 0 0 0 0 0 0 0 1
//4 5 6

func main() {
	list1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 0,
										Next: &ListNode{
											Val: 0,
											Next: &ListNode{
												Val: 0,
												Next: &ListNode{
													Val: 0,
													Next: &ListNode{
														Val: 0,
														Next: &ListNode{
															Val: 0,
															Next: &ListNode{
																Val: 0,
																Next: &ListNode{
																	Val: 0,
																	Next: &ListNode{
																		Val: 0,
																		Next: &ListNode{
																			Val: 0,
																			Next: &ListNode{
																				Val: 0,
																				Next: &ListNode{
																					Val: 0,
																					Next: &ListNode{
																						Val: 0,
																						Next: &ListNode{
																							Val: 0,
																							Next: &ListNode{
																								Val: 1,
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	list1 = ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	list2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	fmt.Println(addTwoNumbers1(&list1, &list2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	firstNumber := make([]int, 0, 100)
	secondNumber := make([]int, 0, 100)
	for l1 != nil {
		firstNumber = append(firstNumber, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		secondNumber = append(secondNumber, l2.Val)
		l2 = l2.Next
	}

	firstLen := len(firstNumber)
	secondLen := len(secondNumber)
	maxLen := secondLen
	if firstLen > maxLen {
		maxLen = firstLen
	}

	// if total.Cmp(big.NewInt(int64(0))) == 0 {
	// 	return &ListNode{Val: 0}
	// }

	var head *ListNode
	var prev *ListNode
	var shift bool
	var total int
	// 2 4 3
	// 5 6 4
	for i := 0; i < maxLen; i++ {
		if firstLen > i && secondLen > i {
			total = firstNumber[i] + secondNumber[i]
		} else if secondLen > i {
			total = secondNumber[i]
		} else if firstLen > i {
			total = firstNumber[i]
		}
		if shift {
			total += 1
			shift = false
		}
		if total >= 10 {
			total = total % 10
			shift = true
		}

		if head == nil {
			head = new(ListNode)
			head.Val = total
			prev = head
			continue
		}
		list := new(ListNode)
		list.Val = total
		prev.Next = list
		prev = list

	}
	if shift {
		list := new(ListNode)
		list.Val = 1
		prev.Next = list
	}

	return head
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	dummy := new(ListNode)

	for node := dummy; l1 != nil || l2 != nil || carry > 0; node = node.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode{carry % 10, nil}
		carry /= 10
	}

	return dummy.Next
}
