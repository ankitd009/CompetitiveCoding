package bfsshortestreach;

import java.io.IOException;
import java.io.InputStream;
import java.io.PrintWriter;
import java.util.Arrays;
import java.util.HashMap;
import java.util.HashSet;
import java.util.InputMismatchException;
import java.util.LinkedList;
import java.util.List;
import java.util.Set;

public class Solution {
	static final int EDGE_COST = 6, WHITE = 0, GRAY = 1, BLACK = 2, INF = Integer.MAX_VALUE;

	static Method br;
	static PrintWriter pw;
	static {
		br = new Method(System.in);
		pw = new PrintWriter(System.out);
	}

	void letsRoll() {
		int noOfNodes, noOfEdges, startNode;
		HashMap<Integer, Set<Integer>> adjList = new HashMap<>();
		noOfNodes = br.readInt();
		noOfEdges = br.readInt();
		for (int i = 0; i < noOfEdges; i++) {
			int start = br.readInt();
			int end = br.readInt();
			// add edge
			addEdge(adjList, start, end);
			// add reverse edge
			addEdge(adjList, end, start);
		}

		startNode = br.readInt();

		// call bfs
		int[] d = BFS(adjList, startNode, noOfNodes);
		// print the ans
		for (int i = 1; i <= noOfNodes; i++) {
			if (i != startNode) {
				if (d[i] == INF) {
					pw.print("-1 ");
				} else {
					pw.print(d[i] * EDGE_COST + " ");
				}
			}
		}
		pw.println();
	}

	public int[] BFS(HashMap<Integer, Set<Integer>> adjList, int startNode, int noOfNodes) {
		int[] color = new int[noOfNodes + 1];
		int[] d = new int[noOfNodes + 1];
		Arrays.fill(color, WHITE);
		Arrays.fill(d, INF);

		color[startNode] = GRAY;
		d[startNode] = 0;
		List<Integer> q = new LinkedList<Integer>();
		q.add(startNode);
		while (!q.isEmpty()) {
			int u = q.remove(0);
			if (adjList.get(u) != null) {
				for (int v : adjList.get(u)) {
					if (color[v] == WHITE) {
						color[v] = GRAY;
						d[v] = d[u] + 1;
						q.add(v);
					}
				}

			}
			color[u] = BLACK;
		}
		return d;
	}

	// Do not add duplicate edges
	public void addEdge(HashMap<Integer, Set<Integer>> adjList, int start, int dest) {
		Set<Integer> l = adjList.get(start);
		if (l == null) {
			l = new HashSet<Integer>();
		}
		l.add(dest);
		adjList.put(start, l);
	}

	public void init() {
		int numOfQueries = br.readInt();
		for (int i = 0; i < numOfQueries; i++) {
			letsRoll();
		}
		pw.flush();
		pw.close();
	}

	public static void main(String[] args) {
		new Solution().init();
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
