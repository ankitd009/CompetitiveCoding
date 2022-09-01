// https://leetcode.com/problems/palindrome-number

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    return x == reverseNumber(x)
}

func reverseNumber(x int)int {
    resNum := 0
    for num:=x; num > 0; num=num/10 {
        digit := num%10
        resNum = resNum * 10 + digit
    }
    return resNum
}
