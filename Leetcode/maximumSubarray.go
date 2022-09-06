// https://leetcode.com/problems/maximum-subarray/

// Kadane’s Algorithm
// https://www.geeksforgeeks.org/largest-sum-contiguous-subarray/

func maxSubArray(nums []int) int {
    
    maxSoFar := -1000000
    maxEndHere := 0
    for i := 0; i < len(nums); i++ {
        maxEndHere = maxEndHere + nums[i]
        if maxSoFar < maxEndHere{
            maxSoFar = maxEndHere;
        }
        if maxEndHere < 0{
            maxEndHere = 0;
        }
                
    }
    return maxSoFar
}
