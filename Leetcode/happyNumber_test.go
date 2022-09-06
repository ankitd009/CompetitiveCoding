package leetcode

import "testing"

func TestIsHappySuccess(t *testing.T){
  res := isHappy(19)
  if !res {
	  t.Errorf("Expected: true, Got: %v", res)
  }
}

func TestIsHappyFailure(t *testing.T){
  res := isHappy(2)
  if res {
	  t.Errorf("Expected: false, Got: %v", res)
  }
}
