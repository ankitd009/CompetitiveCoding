package leetcode

// https://leetcode.com/problems/merge-sorted-array/submissions/

func merge(nums1 []int, m int, nums2 []int, n int)  {
    
    min := -1000000001
    
    if m == 0 && n == 0 {
        return
    }else if m == 0 {
        for j := 0 ; j < n ; j++ {
            nums1[j] = nums2[j]
        }
        return
    }else if n == 0 {
        return
    }
    
    l := m + n
    
    t := l-1
    i1 := m - 1
    i2 := n - 1
    
    for i := 0  ; i < l; i++ {
        elem1 := min
        elem2 := min
        if i1 >= 0 {
            elem1 = nums1[i1]
        }
        if i2 >= 0 {
            elem2 = nums2[i2]
        }
        
        e := elem1
        i1--
        if elem2 > elem1{
            e = elem2
            i1++
            i2--
        }
        
        nums1[t] = e
        t--
    }
}
