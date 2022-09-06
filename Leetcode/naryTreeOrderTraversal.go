package leetcode

// https://leetcode.com/problems/n-ary-tree-level-order-traversal

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	l := make(map[int][]int)
	traverse(root, l , 0)
	ar := make([][]int, len(l) ,len(l))
	for k,v := range l {
		ar[k] = v
	}
	return ar
}

func traverse(node *Node, levelM map[int][]int, depth int) {
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

    if node.Children != nil {
        for _,v := range node.Children {
            traverse(v, levelM, depth + 1)        
        }
    }
}
