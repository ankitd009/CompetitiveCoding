package leetcode

// https://leetcode.com/problems/merge-two-sorted-lists/submissions/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	itr1 := list1
	itr2 := list2

	var resList *ListNode

	for itr1 != nil || itr2 != nil  {
		var val1, val2 *int

		if itr1 != nil {
			val1 = &itr1.Val
		}

		if itr2 != nil {
			val2 = &itr2.Val
		}

		if val1 != nil && val2 != nil {
			if *val1 < *val2 {
				resList = appendNode(resList, *val1)
				if itr1 != nil {
					itr1 = itr1.Next
				}
			}else{
				resList= appendNode(resList, *val2)
				if itr2 != nil {
					itr2 = itr2.Next
				}
			}
		}else if val1 == nil {
			resList = appendNode(resList, *val2)
			if itr2 != nil {
				itr2 = itr2.Next
			}
		}else if val2 == nil {
			resList = appendNode(resList, *val1)
			if itr1 != nil {
				itr1 = itr1.Next
			}
		}
	}

	return resList
}

func appendNode(list *ListNode, val int) *ListNode {
	node := &ListNode{
		Next: nil,
		Val: val,
	}

	if list == nil {
		list = node
		return list
	}


	itr := list
	for itr.Next != nil {
		itr = itr.Next
	}
	itr.Next = node
	return list
}
