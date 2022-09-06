// https://leetcode.com/problems/happy-number

func isHappy(n int) bool {
    num := n
    unique := make(map[int]bool)
    for {
        num = digitSquareSum(num)
        if num == 1 {
            return true
        }
        if _, exists := unique[num]; exists {
            return false
        }
        unique[num] = true
    }
    return false
}

func digitSquareSum(n int)int{
    num := n
    sum := 0
    for num > 0 {
        digit := num % 10
        sum = sum + digit * digit
        num = num/10
    }
    return sum
}
