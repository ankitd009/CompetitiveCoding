package leetcode

import "testing"

func TestContainsDuplicate_Success(t *testing.T){
	nums := []int{1,2,3,1}
	res := containsDuplicate(nums)
        if !res {
		t.Errorf("Expected: true, Got: %+v", res)
	}
}
