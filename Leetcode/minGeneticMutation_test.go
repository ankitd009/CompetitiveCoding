package leetcode

import "testing"

func TestFindMinMutations(t *testing.T){

	res := minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"})
  	if res != 1 {
		t.Errorf("Expected: %+v, Got: %+v", 1 ,res)
  	}
	
}
