// https://www.hackerrank.com/challenges/designer-pdf-viewer/problem

import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;

public class Solution {

    static int getCharWeightage(int[] h, char ch){
        return h[ch - 97];
    }

    static int getMaxWeightage(int[] h, String word){
        char[] wordCharacters = word.toCharArray();
        int maxWeightage = 0;
        
        for(char ch : wordCharacters){
            int weightage = getCharWeightage(h, ch);
            if (maxWeightage < weightage){
                maxWeightage = weightage;
            }
        }
        return maxWeightage;
    }

    static int designerPdfViewer(int[] h, String word) {
        int maxWeightage = getMaxWeightage(h, word);
        return maxWeightage * word.length();
    }

    private static final Scanner scanner = new Scanner(System.in);

    public static void main(String[] args) throws IOException {
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int[] h = new int[26];

        String[] hItems = scanner.nextLine().split(" ");
        scanner.skip("(\r\n|[\n\r\u2028\u2029\u0085])?");

        for (int i = 0; i < 26; i++) {
            int hItem = Integer.parseInt(hItems[i]);
            h[i] = hItem;
        }

        String word = scanner.nextLine();

        int result = designerPdfViewer(h, word);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedWriter.close();

        scanner.close();
    }
}
