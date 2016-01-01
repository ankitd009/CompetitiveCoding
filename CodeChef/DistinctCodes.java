/**
 * CodeChef
 * Contest - LTIME31
 * ProblemCode - DISTCODE
 */
package codechef;

import java.io.IOException;
import java.io.InputStream;
import java.io.PrintWriter;
import java.util.HashSet;
import java.util.InputMismatchException;

/**
 * @author ankitd009
 * 
 */
class DistinctCodes {
    static PrintWriter pw;
    static Method br;

    static {
        pw = new PrintWriter(System.out);
        br = new Method(System.in);
    }

    void letsRoll() {
        int testCases = br.readInt();
        while (testCases-- > 0) {
            String input = br.readString();
            HashSet uniqueCodes = new HashSet();
            for (int i = 0, len = input.length() - 1; i < len; i++) {
                String temp = "" + input.charAt(i) + input.charAt(i + 1);
                if (!uniqueCodes.contains(temp)) {
                    uniqueCodes.add(temp);
                }
            }
            pw.println(uniqueCodes.size());
        }
    }

    /**
     * @param args
     */
    public static void main(String[] args) {
        // TODO Auto-generated method stub
        new DistinctCodes().letsRoll();
    }

}

class Method {

    private final InputStream stream;
    private final byte[] buf = new byte[1024];
    private int curChar;
    private int numChars;
    private SpaceCharFilter filter;

    Method(InputStream stream) {
        this.stream = stream;
    }

    public final int read() {

        if (numChars == -1) {
            throw new InputMismatchException();
        }
        if (curChar >= numChars) {
            curChar = 0;
            try {
                numChars = stream.read(buf);
            } catch (IOException e) {
                throw new InputMismatchException();
            }
            if (numChars <= 0) {
                return -1;
            }
        }
        return buf[curChar++];
    }

    public final int readInt() {
        int c = read();
        while (isSpaceChar(c)) {
            c = read();
        }
        int sgn = 1;
        if (c == '-') {
            sgn = -1;
            c = read();
        }
        int res = 0;
        do {
            if (c < '0' || c > '9') {
                throw new InputMismatchException();
            }
            res *= 10;
            res += c - '0';
            c = read();
        } while (!isSpaceChar(c));
        return res * sgn;
    }

    public final long readLong() {
        int c = read();
        while (isSpaceChar(c)) {
            c = read();
        }
        int sgn = 1;
        if (c == '-') {
            sgn = -1;
            c = read();
        }
        long res = 0;
        do {
            if (c < '0' || c > '9') {
                throw new InputMismatchException();
            }
            res *= 10;
            res += c - '0';
            c = read();
        } while (!isSpaceChar(c));
        return res * sgn;
    }

    public final String readString() {
        int c = read();
        while (isSpaceChar(c)) {
            c = read();
        }
        StringBuilder res = new StringBuilder();
        do {
            res.appendCodePoint(c);
            c = read();
        } while (!isSpaceChar(c));
        return res.toString();
    }

    public final boolean isSpaceChar(int c) {
        if (filter != null) {
            return filter.isSpaceChar(c);
        }
        return c == ' ' || c == '\n' || c == '\r' || c == '\t' || c == -1;
    }

    public final String next() {
        return readString();
    }

    public interface SpaceCharFilter {
        public boolean isSpaceChar(int ch);
    }

}
