// https://leetcode.com/problems/symmetric-tree/submissions/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    return checkSymmetric(root.Left, root.Right)
}

func checkSymmetric(l, r *TreeNode) bool {
    if (l == nil && r != nil) || (l != nil && r == nil) {
        return false
    }
    if (l == nil && r == nil){
        return true
    }
    return l.Val == r.Val && checkSymmetric(l.Left, r.Right) && checkSymmetric(l.Right, r.Left)
}

