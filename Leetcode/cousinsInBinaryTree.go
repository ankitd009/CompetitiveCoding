// https://leetcode.com/problems/cousins-in-binary-tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
    
    nX, dX := findParentDepth(root, x, 0)
    nY, dY := findParentDepth(root, y, 0)
    
    return dX == dY && nX != nY
    
}

func findParentDepth(node *TreeNode, val int, depth int)(*TreeNode, int){
    if node == nil {
        return nil, -1
    }
    if node.Left != nil && node.Left.Val == val{
        return node, depth + 1
    }
    if node.Right != nil && node.Right.Val == val{
        return node, depth + 1
    }
    nL, dL := findParentDepth(node.Left, val, depth +1)
    if nL != nil && dL != -1 {
        return nL, dL
    }
    nR, dR := findParentDepth(node.Right, val, depth +1)
    if nR != nil && dR != -1 {
        return nR, dR
    }
    return nil, -1
}
