package leetcode
// https://leetcode.com/problems/find-the-town-judge/submissions/


import "fmt"


func initMatrix(n int) [][]bool {
	matrix := make([][]bool, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]bool, n)
	}
	return matrix
}

func isTrustedByEveryone(lookup map[int]bool,index int, trustMatrix [][]bool) bool {

	val, found := lookup[index]
	if found {
		return val
	}

	for i := 1; i < len(trustMatrix); i++ {
		if !trustMatrix[i][index] && index != i {
			lookup[index] = false
			return false
		}
	}
	lookup[index] = true
	return true
}

func doesNotTrustAnyone(lookup map[int]bool, index int, trustMatrix [][]bool) bool {
	val, found := lookup[index]
	if found {
		return val
	}
	for i := 1; i < len(trustMatrix); i++ {
		if trustMatrix[index][i] && index != i {
			lookup[index] = false
			return false
		}
	}
	lookup[index] = true
	return true
}

func findJudge(N int, trusts [][]int) int {
	if N == 1 {
		return 1
	}

	trustMatrix := initMatrix(N + 1)

	for _, t := range trusts {
		trustMatrix[t[0]][t[1]] = true
	}

	trustLookup := make(map[int]bool)
	distrustLookup := make(map[int]bool)

	for i := 1; i <= N; i++ {
		isT := isTrustedByEveryone(trustLookup,i, trustMatrix)
		dOesNOt := doesNotTrustAnyone(distrustLookup, i, trustMatrix)
		if isT && dOesNOt {
			return i
		}
	}

	return -1
}
