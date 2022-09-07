func twoSum(nums []int, target int) []int {
    freqMap := make(map[int]int)
    indexMap := make(map[int][]int)
    for i, n := range nums {
        if count, exists := freqMap[n]; exists {
            freqMap[n] = count + 1
        }else {
            freqMap[n] = 1
        }
        
        if indices, exists := indexMap[n]; exists {
            indexMap[n] = append(indices, i)
        }else {
            indices = make([]int, 0 ,0)
            indices = append(indices, i)
            indexMap[n] = indices
        }
    }
    for _, n := range nums {
        borrow(freqMap, n)
        canBorrow := borrow(freqMap, target - n)
        if canBorrow {
            return []int{getIndex(indexMap, n), getIndex(indexMap, target-n)}
        }
        putBack(freqMap, n)
    }
    return []int{}
}

func borrow(freqMap map[int]int, num int) bool{
    if c, exists := freqMap[num]; exists {
        freqMap[num] = c - 1
        return c > 0
    }
    return false
}

func putBack(freqMap map[int]int, num int) {
    c := freqMap[num]
    freqMap[num] = c + 1
}

func getIndex(indexMap map[int][]int, num int)int {
    indices, exists := indexMap[num]
    if !exists {
        return -1
    }
    ind := indices[0]
    
    indexMap[num] = indices[1:]
    
    return ind
}
