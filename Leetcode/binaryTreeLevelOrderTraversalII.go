package leetcode

// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	l := make(map[int][]int)
	traverse(root, l , 0)
	ar := make([][]int, len(l) ,len(l))
	for i := 0 ; i < len(l); i++ {
		ind := len(l)-1 -i
		ar[ind] = l[i]
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
