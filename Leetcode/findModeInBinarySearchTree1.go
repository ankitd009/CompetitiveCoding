package leetcode

// https://leetcode.com/problems/find-mode-in-binary-search-tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
    
    freq := make(map[int]int)
    
    traverse(freq, root)
    
    highFreq := -1
    for _, v := range freq {
        if v > highFreq {
            highFreq = v
        }
    }
    
    nums := make([]int, 0,0)
    
    for k, v := range freq {
        if v == highFreq {
            nums = append(nums, k)
        }
    }
    
    return nums
}

func traverse(freq map[int]int, node *TreeNode){
    if node == nil {
        return
    }
    
    if count, exists := freq[node.Val]; exists {
        freq[node.Val] = count + 1
    }else {
        freq[node.Val] = 1
    }
    
    traverse(freq, node.Right)
    traverse(freq, node.Left)
}
