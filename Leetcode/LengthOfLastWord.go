package leetcode

// https://leetcode.com/problems/length-of-last-word/submissions/

func lengthOfLastWord(s string) int {
	s3 := strings.Fields(strings.TrimSpace(s))
	return len(s3[len(s3)-1])
}
