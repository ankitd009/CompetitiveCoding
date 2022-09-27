package leetcode

import "testing"

func TestFindErrorNums(t *testing.T){
	nums := []int{1,2,2,4}
	res := findErrorNums(nums)
  
  if res[0] != 2 && res[1] != 3{
		t.Errorf("Expected: %+v, Got: %+v", []int{2, 3} ,res)
	}
}
