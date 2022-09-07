package leetcode

import "testing"

func TestTwoSumSuccess(t *testing.T){
  result := twoSum([]int{3,3}, 6)
  if result != []int{0,1} {
    t.Errof("Expected: %+v, Got: %+v", result)
  }  
}

func TestTwoSumSuccessII(t *testing.T){
  result := twoSum([]int{2,7, 11, 15}, 9)
  if result != []int{0,1} {
    t.Errof("Expected: %+v, Got: %+v", result)
  }  
}
