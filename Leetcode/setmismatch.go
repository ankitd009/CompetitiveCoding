// https://leetcode.com/problems/set-mismatch

func findErrorNums(nums []int) []int {
    twice := -1
    missing := -1
    
    freq := make([]int, len(nums)+1, len(nums)+1)
    
    for _, i := range nums{
        freq[i]++
    }
    
    for i := 0; i < len(freq); i++ {
        c := freq[i]
        if c == 2 {
            twice = i
        }else if c == 0 {
            missing = i
        }
    }
    
    return []int{twice, missing}
}

func findErrorNumsMap(nums []int) []int {
    twice := -1
    missing := -1
    
    numFreq := make(map[int]int)
    
    for _, i := range nums {
        c := numFreq[i]
        c = c + 1
        numFreq[i] = c
    }
    
    for i := 1; i <= len(nums); i++ {
        c, exists := numFreq[i]
        if c == 2 {
            twice = i
        }else if !exists {
            missing = i
        }
    }
    
    return []int{twice, missing}
}
