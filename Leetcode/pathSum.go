package leetcode

// https://leetcode.com/problems/path-sum

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root != nil && root.Left == nil && root.Right == nil && root.Val == targetSum {
        return true
    }
    return hasSum(root, 0, targetSum)
}

func hasSum(node *TreeNode, curSum, targetSum int)bool{
    if node == nil {
        return false
    }
    sum := curSum + node.Val
    if sum == targetSum && node.Left == nil && node.Right == nil {
         return true
    }
    return hasSum(node.Left,sum , targetSum) || hasSum(node.Right, sum, targetSum)
}
