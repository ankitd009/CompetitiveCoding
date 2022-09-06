// https://leetcode.com/problems/contains-duplicate

package leetcode

func containsDuplicate(nums []int) bool {
    unique := make(map[int]bool)
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        if _, exists := unique[num]; exists {
            return true
        } else{  
            unique[num] = true
        }
    }  
    return false
}
