package leetcode

// https://leetcode.com/problems/sum-of-left-leaves/submissions/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    return sumLeftLeaves(root)
}

func sumLeftLeaves(node *TreeNode) int {
    if node == nil {
        return 0
    }
    sum := 0
    if isLeaf(node.Left){
        sum += node.Left.Val
    }
    return sum + sumLeftLeaves(node.Left) + sumLeftLeaves(node.Right)
}

func isLeaf(node *TreeNode)bool{
    return node != nil && node.Left == nil && node.Right == nil
}
