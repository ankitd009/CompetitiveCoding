package leetcode

// https://leetcode.com/problems/sum-of-even-numbers-after-queries/

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
    
    results := make([]int,len(queries),len(queries))
    evenSum := 0
    
    for _, num := range nums {
        if (num % 2 == 0 ){
            evenSum = evenSum + num    
        }        
    }
    
    for i, query := range queries {
        
        valToAdd := query[0]
        index := query[1]
        
        orgNum := nums[index]
        orgNumIsEven := orgNum%2 == 0
        valToAddIsEven := valToAdd%2 == 0
        
        if (orgNumIsEven && valToAddIsEven) {
            evenSum = evenSum + valToAdd
        } else if (!orgNumIsEven && !valToAddIsEven) {
            evenSum = evenSum + orgNum + valToAdd            
        } else if (orgNumIsEven && !valToAddIsEven){
            evenSum = evenSum - orgNum
        }
        nums[index] = nums[index] + valToAdd
        results[i] = evenSum
    }
    
    return results
}
