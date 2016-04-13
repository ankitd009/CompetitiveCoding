/**
 * CodeChef
 * Contest - March16
 * ProblemCode - CHEFSPL
 */
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;

/**
 * 
 * @author ankitd009
 */
class ChefAndSpecialDishedIndex {

    private static final String YES = "YES";
    private static final String NO = "NO";

    static PrintWriter pw;
    static BufferedReader br;

    static {
	br = new BufferedReader(new InputStreamReader(System.in));
	pw = new PrintWriter(System.out);
    }

    private String returnStringWithoutCharAt(String input, int index) {
	StringBuilder str = new StringBuilder();
	for (int i = 0; i < input.length(); i++) {
	    if (i != index) {
		str.append(input.charAt(i));
	    }
	}
	return str.toString();
    }

    public boolean oddCase(String firstPart, String secondPart) {
	boolean deletedChar = false;
	boolean ans = true;
	int i = 0, j = 0;
	while (i < firstPart.length() && j < secondPart.length()) {
	    if (firstPart.charAt(i) == secondPart.charAt(j)) {
		i++;
		j++;
	    } else if (!deletedChar) {
		deletedChar = true;
		// removing ith character from firstHalf
		String firstPartWithoutI = returnStringWithoutCharAt(firstPart, i);
		ans = firstPartWithoutI.equals(secondPart);
		// removing jth character from secondHalf
		String secondPartWithoutJ = returnStringWithoutCharAt(secondPart, j);
		ans |= secondPartWithoutJ.equals(firstPart);
		break;
	    } else {
		ans = false;
		break;
	    }
	}
	return ans;
    }

    public void letsRoll() throws IOException {
	int testCases = Integer.parseInt(br.readLine());
	while (testCases-- > 0) {
	    String input = br.readLine();
	    boolean ans;
	    int length = input.length();
	    if (length == 1) {
		ans = false;
	    } else {
		int half = length / 2;
		if (length % 2 == 0) {
		   /*
		    *if length is even, no deletion operation just check for
		    *equality
		    */
		    String firstHalf = input.substring(0, half);
		    String secondHalf = input.substring(half, length);
		    ans = firstHalf.equals(secondHalf);
		} else {
		    /*if length is odd, one deletion operation has to be
		     *performed secondPart has the odd letter
		     */
		    String firstPart = input.substring(0, half);
		    String secondPart = input.substring(half, length);
		    ans = oddCase(firstPart, secondPart);
		    if (!ans) {
		   /*
			*check if adding odd letter to firstHalf makes any
			*difference
			*/
			firstPart = input.substring(0, half + 1);
			secondPart = input.substring(half + 1, length);
			ans = oddCase(firstPart, secondPart);
		    }
		}
	    }
	    pw.println((ans ? YES : NO));
	}
	pw.flush();
	pw.close();
    }

    public static void main(String[] args) throws IOException {
	    new ChefAndSpecialDishedIndex().letsRoll();
    }

}
