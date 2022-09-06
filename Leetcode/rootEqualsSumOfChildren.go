package leetcode

// https://leetcode.com/problems/root-equals-sum-of-children/submissions/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func checkTree(root *TreeNode) bool {
    return root.Val == root.Left.Val + root.Right.Val
}
