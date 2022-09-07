package leetcode

import "testing"

func TestTwoSumSuccess(t *testing.T){
  result := twoSum([]int{3,3}, 6)
  if result[0] != 0 && result[1] == 1{
    t.Errorf("Expected: %+v, Got: %+v", []int{0,1}, result)
  }  
}

func TestTwoSumSuccessII(t *testing.T){
  result := twoSum([]int{2,7, 11, 15}, 9)
  if result[0] != 0 && result[1] == 1{
    t.Errorf("Expected: %+v, Got: %+v", []int{0,1}, result)
  }  
}
