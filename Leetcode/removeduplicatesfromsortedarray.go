// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
    i := 0
    swap := 0
    for i < len(nums) {
        index := getIndexOfNextDistinctElem(nums, i)
        if index == -1{
            break
        }
        nums[swap+1] = nums[index]
        i = index
        swap++
    }
    return swap + 1
}

func getIndexOfNextDistinctElem(nums []int, curIndex int)int{
    i := curIndex + 1
    for i < len(nums){
        if nums[curIndex] != nums[i]{
            return i
        }
        i++
    }
    return -1
}
