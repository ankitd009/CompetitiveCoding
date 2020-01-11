// https://leetcode.com/problems/convert-bst-to-greater-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    convert(root, 0)
    return root
}

func convert(node *TreeNode, sum int)int{
    if node == nil {
        return 0
    }
    t := node.Val
    rtSum := convert(node.Right, sum)
    node.Val = node.Val+ rtSum + sum
    ltSum := convert(node.Left, sum + rtSum + t)
    return rtSum + ltSum + t
}
