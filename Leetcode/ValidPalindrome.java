// https://leetcode.com/problems/valid-palindrome

class Solution {
   public boolean isPalindrome(String s) {
        String input = s.trim().toLowerCase();
        if(input.length() == 0 ){
            return true;
        }
        char[] chars = input.toCharArray();
        int inValidCount = 0;
        for(int i = 0 ; i < chars.length; i++){
            if(!isValid(chars[i])){
                inValidCount++;
            }
        }
        char[] validChars = new char[chars.length - inValidCount];
        int i = 0;
        for(char c: chars){
            if(isValid(c)){
                validChars[i] = c;
                i++;
            }
        }
        return isPalindrome(validChars);
    }

    public static boolean isValid(char c){
        return Character.isDigit(c) || Character.isLetter(c);
    }

    public static boolean isPalindrome(char[] chars){
        int i = 0, j = chars.length-1;
        while(i<j) {
            if (chars[i] != chars[j]){
                return false;
            }
            i++;
            j--;
        }
        return true;
    }
    
}
