package leetcode

// https://leetcode.com/problems/fizz-buzz

func fizzBuzz(n int) []string {
    sol := make([]string, 0, n)
    for i := 1 ; i <=n ;i++{
        sol = append(sol, getStr(i))
    }
    return sol
}

func getStr(n int)string{
    if n % 15 == 0 {
        return "FizzBuzz"
    }else if n % 5 == 0 {
        return "Buzz"
    }else if n % 3 == 0 {
        return "Fizz"
    }
    return strconv.Itoa(n)
}
