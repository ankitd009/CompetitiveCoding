package leetcode

// https://leetcode.com/problems/binary-tree-level-order-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	l := make(map[int][]int)
	traverse(root, l , 0)
	ar := make([][]int, len(l) ,len(l))
	for k,v := range l {
		ar[k] = v
	}
	return ar
}

func traverse(node *TreeNode, levelM map[int][]int, depth int) {
	if node == nil {
		return
	}
	if elems, exists := levelM[depth]; exists {
		elems = append(elems, node.Val)
		levelM[depth] = elems
	} else {
		elems = make([]int, 0 ,0)
		elems = append(elems, node.Val)
		levelM[depth] = elems
	}

	traverse(node.Left, levelM, depth + 1)
	traverse(node.Right, levelM, depth + 1)
}
