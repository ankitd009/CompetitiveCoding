// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/submissions/

package leetcode

func findDisappearedNumbers(nums []int) []int {
    
    freq := make([]int, len(nums)+1, len(nums)+1)
    
    for _, i := range nums{
        freq[i]++
    }
    
    res := make([]int, 0, 0)
    
    for i := 1; i < len(freq); i++ {
        c := freq[i]
        if c == 0 {
            res = append(res, i)
        }
    }
    
    return res
}
