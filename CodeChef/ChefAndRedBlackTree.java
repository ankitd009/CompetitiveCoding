/**
 * CodeChef
 * Contest - NOV14
 * ProblemCode - RBTREE
 */
import java.io.IOException;
import java.io.InputStream;
import java.io.PrintWriter;
import java.util.InputMismatchException;

/**
 * 
 * @author ankitd009
 */
class OtherMethod {

    static Method2 br;
    static PrintWriter pw;

    static {
        br = new Method2(System.in);
        pw = new PrintWriter(System.out);
    }

    public int nearestSmallestTwoPower(long num) {
        int power = 0;
        int currPower = 1;

        while (!(currPower > num)) {
            currPower <<= 1;
            power++;
        }

        return power - 1;
    }

    public int nearestLargestTwoPower(long num) {
        int power = 0;
        int currPower = 1;

        while (!(currPower > num)) {
            currPower <<= 1;
            power++;
        }

        return power;

    }

    public long pathLenght(long x, long y) {
        long path = 0;
        long xParent = x;
        long yParent = y;
        while (xParent != yParent) {
            if (xParent > yParent) {
                xParent >>= 1;
            } else {
                yParent >>= 1;
            }
            path++;
        }
        return path;
    }

    public void letsRoll() {

        int xlevel, ylevel, testCases = br.readInt();
        long redNodes, blackNodes, x, y, pathLength, pathByTwo;
        boolean rootColorBlack = true, currentRootColorBlack = true;
        while ((testCases--) > 0) {
            String type = br.readString();
            if (type.equals("Qi")) {
                currentRootColorBlack = !currentRootColorBlack;
            } else {
                x = br.readLong();
                y = br.readLong();
                xlevel = nearestSmallestTwoPower(x) + 1;
                ylevel = nearestSmallestTwoPower(y) + 1;
                pathLength = pathLenght(x, y);
                pathByTwo = pathLength >> 1;
                if (!((xlevel & 1) == 1) && !((ylevel & 1) == 1)) {
                    if (currentRootColorBlack == rootColorBlack) {
                        redNodes = pathByTwo + 1;
                        blackNodes = pathByTwo;
                    } else {
                        blackNodes = pathByTwo + 1;
                        redNodes = pathByTwo;
                    }
                } else if (((xlevel & 1) == 1) && ((ylevel & 1) == 1)) {
                    if (currentRootColorBlack == rootColorBlack) {
                        blackNodes = pathByTwo + 1;
                        redNodes = pathByTwo;
                    } else {
                        redNodes = pathByTwo + 1;
                        blackNodes = pathByTwo;
                    }
                } else {
                    redNodes = pathByTwo + 1;
                    blackNodes = pathByTwo + 1;
                }
                if (type.equals("Qb")) {
                    pw.println(blackNodes);
                } else {
                    pw.println(redNodes);
                }
            }
        }// testcases
        pw.flush();
    }// letsRoll

    public static void main(String[] args) {
        new OtherMethod().letsRoll();
    } // main
}

class Method2 {

    private final InputStream stream;

    private final byte[] buf = new byte[1024];

    private int curChar;

    private int numChars;

    private SpaceCharFilter filter;

    public Method2(InputStream stream) {

        this.stream = stream;

    }

    public int read() {

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

    public int readInt() {

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

    public long readLong() {
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

    public String readString() {

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

    public boolean isSpaceChar(int c) {

        if (filter != null) {
            return filter.isSpaceChar(c);
        }

        return c == ' ' || c == '\n' || c == '\r' || c == '\t' || c == -1;

    }

    public String next() {

        return readString();
    }

    public interface SpaceCharFilter {

        public boolean isSpaceChar(int ch);
    }
}
