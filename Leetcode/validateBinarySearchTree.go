package leetcode

// https://leetcode.com/problems/validate-binary-search-tree/submissions/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	tO := inorderTraversal(root)
	if len(tO) == 1 {
		return true
	}
	for i := 0 ; i < len(tO) -1 ; i++ {
		if tO[i] >= tO[i+1]{
			return false
		}
	}
	return true
}

type  traverseOrder struct {
	vals []int
}

func inorderTraversal(root *TreeNode) []int {
	vals  := make([]int, 0, 0)
	if root == nil {
		return vals
	}
	tO := &traverseOrder {
		vals: vals,
	}
	traverse(root, tO)
	return tO.vals
}

func traverse(node *TreeNode, order *traverseOrder){
	if node == nil {
		return
	}

	if node.Left != nil {
		traverse(node.Left, order)
	}

	order.vals = append(order.vals, node.Val)

	if node.Right != nil {
		traverse(node.Right, order)
	}
}
