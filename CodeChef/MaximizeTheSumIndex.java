/**
 * CodeChef
 * Contest - Mar16
 */
import java.io.IOException;
import java.io.InputStream;
import java.io.PrintWriter;
import java.util.InputMismatchException;

class MaximizeTheSumIndex {
    static PrintWriter pw;
    static Method br;

    static {
	pw = new PrintWriter(System.out);
	br = new Method(System.in);
    }

    public void letsRoll() throws IOException {
	int testCases = br.readInt();
	long[] a, b;
	int n, k;
	while (testCases-- > 0) {
	    n = br.readInt();
	    k = br.readInt();
	    a = new long[n];
	    b = new long[n];
	    for (int i = 0; i < n; i++) {
		    a[i] = br.readLong();
	    }
	    for (int i = 0; i < n; i++) {
		    b[i] = br.readLong();
	    }
	    long sum = 0;
	    for (int i = 0; i < n; i++) {
		    sum += a[i] * b[i];
	    }
	    long maxOne = Long.MIN_VALUE;
	    for (int i = 0; i < n; i++) {
		    maxOne = Math.max(maxOne, k * b[i]);
	    }
	    k = -k;
	    for (int i = 0; i < n; i++) {
		    maxOne = Math.max(maxOne, k * b[i]);
	    }
	    pw.println(sum + maxOne);
	}
	pw.flush();
	pw.close();
    }

    public static void main(String[] args) throws IOException {
	new MaximizeTheSumIndex().letsRoll();
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
