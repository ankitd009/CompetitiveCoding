package leetcode

// https://leetcode.com/problems/balanced-binary-tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return validate(root)
}

func validate(root *TreeNode) bool {
    if root == nil {
        return true
    }
    lH := maxDepth(root.Left)
    rH := maxDepth(root.Right)
    if lH > rH && lH - rH > 1 {
        return false
    }else if rH > lH && rH - lH > 1{
        return false
    }
    return validate(root.Left) && validate(root.Right)
}

func maxDepth(root *TreeNode) int{
    if root == nil {
        return 0;
    }
    return 1 + max(maxDepth(root.Left), maxDepth(root.Right));
}

func max(x,y int)int{
    if x > y {
        return x
    }
    return y
}
