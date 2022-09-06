package leetcode

// https://leetcode.com/problems/remove-element/

func removeElement(nums []int, val int) int {
	count := 0
	length := len(nums)
	for i:=0; i< length-count; i++ {
		elem := nums[i]
		if elem == val {
			count++
		}
	}

	if count == 0 || count == length {
		return length - count
	}

	count = 0

	for i:=0; i< length-count; i++ {
		elem := nums[i]
		if elem == val {
			count++
			swap(nums, i, length - count)
		}
		if nums[i] == val {
			i--
		}
	}
	return length- count
}

func swap(nums []int, srcInd, destInd int){
	tmp := nums[srcInd]
	nums[srcInd] = nums[destInd]
	nums[destInd] = tmp
}
