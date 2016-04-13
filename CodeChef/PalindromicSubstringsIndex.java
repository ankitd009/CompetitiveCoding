/**
 * CodeChef
 * Contest - Mar16
 */
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.util.HashSet;

class PalindromicSubstringsIndex {

    private static final String YES = "Yes";
    private static final String NO = "No";

    static PrintWriter pw;
    static BufferedReader br;

    static {
	br = new BufferedReader(new InputStreamReader(System.in));
	pw = new PrintWriter(System.out);
    }

    void letsRoll() throws IOException {
	int testCases = Integer.parseInt(br.readLine());
	while (testCases-- > 0) {
	    String one = br.readLine();
	    String two = br.readLine();
	    HashSet<Character> uniqueCharacterSetForOne = new HashSet<Character>();
	    for (char c : one.toCharArray()) {
		uniqueCharacterSetForOne.add(c);
	    }
	    boolean ans = false;
	    for (char c : two.toCharArray()) {
		if (uniqueCharacterSetForOne.contains(c)) {
		    ans = true;
		    break;
		}
	    }
	    pw.println((ans ? YES : NO));
	}
	pw.flush();
    }

    public static void main(String[] args) throws IOException {
	new PalindromicSubstringsIndex().letsRoll();
    }
}
