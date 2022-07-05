// https://leetcode.com/problems/valid-parentheses

class Solution {
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<Character>();
        char[] input = s.toCharArray();
        for(char x : input){
            if(x == '(' || x == '{' || x == '['){
                stack.push(x);
            }
            else if(x == ')' || x == '}' || x == ']'){
                if(stack.isEmpty()){
                    return false;
                }
                char top = stack.pop();
                if(x == ')' && top != '('){
                    return false;
                }
                if(x == ']' && top != '['){
                    return false;
                }
                if(x == '}' && top != '{'){
                    return false;
                }
            }
        }
        return stack.isEmpty();
    }
}
