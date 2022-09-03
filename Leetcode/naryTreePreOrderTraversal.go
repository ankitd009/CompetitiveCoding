https://leetcode.com/problems/n-ary-tree-preorder-traversal

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
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

func traverse(node *Node, order *traverseOrder){
    if node == nil {
        return
    }
    order.vals = append(order.vals, node.Val)
    
    if node.Children != nil {
        for _, v := range node.Children {
            traverse(v, order)    
        }
    }
}
