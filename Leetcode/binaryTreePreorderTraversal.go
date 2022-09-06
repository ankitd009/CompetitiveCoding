package leetcode

// https://leetcode.com/problems/binary-tree-preorder-traversal/submissions/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
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

type  traverseOrder struct {
    vals []int
}

func traverse(node *TreeNode, order *traverseOrder){
    if node == nil {
        return
    }
    order.vals = append(order.vals, node.Val)
    
    if node.Left != nil {
        traverse(node.Left, order)
    }
    
    if node.Right != nil {
        traverse(node.Right, order)
    }
}
