package leetcode

// https://leetcode.com/problems/remove-linked-list-elements

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *ListNode, val int) *ListNode {
	nList := &ListNode{
		Next: head,
	}
	itr := nList

	for itr != nil {
		if itr.Next != nil &&  itr.Next.Val == val {
			itr.Next = itr.Next.Next
		}else {
			itr = itr.Next
		}

	}

	return nList.Next
}
