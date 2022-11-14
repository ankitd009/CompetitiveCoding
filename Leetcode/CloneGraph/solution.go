// https://leetcode.com/problems/clone-graph/

package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	adjList := getAdjList(node)
	return getNewGraph(adjList, node.Val)
}

func newNode(val int) *Node {
	return &Node{
		Val:       val,
		Neighbors: make([]*Node, 0, 0),
	}
}

func getNewGraph(adjList map[int][]int, rootVal int) *Node {
	nodeRef := make(map[int]*Node)

	for nodeVal := range adjList {
		node := newNode(nodeVal)
		nodeRef[nodeVal] = node
	}

	for nodeVal, neighVals := range adjList {
		node := nodeRef[nodeVal]
		neg := node.Neighbors
		for _, val := range neighVals {
			n := nodeRef[val]
			neg = append(neg, n)
		}
		node.Neighbors = neg
	}

	return nodeRef[rootVal]
}

func getAdjList(node *Node) map[int][]int {
	adjList := make(map[int][]int)
	seen := make(map[int]bool)

	queue := make([]*Node, 0, 0)
	queue = append(queue, node)

	for len(queue) > 0 {
		curN := queue[0]
		queue = queue[1:]

		if _, saw := seen[curN.Val]; saw {
			continue
		} else {
			seen[curN.Val] = true
		}

		if len(curN.Neighbors) > 0 {
			for _, n := range curN.Neighbors {
				putEdge(adjList, curN.Val, n.Val)
				queue = append(queue, n)
			}
		} else if len(curN.Neighbors) == 0 {
			adjList[curN.Val] = []int{}
		}

	}
	return adjList
}

func putEdge(adjList map[int][]int, src, dest int) {

	if elems, exists := adjList[src]; exists {
		elems = append(elems, dest)
		adjList[src] = elems
		return
	}

	elems := make([]int, 0, 0)
	elems = append(elems, dest)
	adjList[src] = elems
}
