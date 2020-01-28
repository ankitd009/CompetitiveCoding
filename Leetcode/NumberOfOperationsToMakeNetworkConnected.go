// https://leetcode.com/problems/number-of-operations-to-make-network-connected/submissions/

// using union by rank and path compression

type subset struct {
    parent int
    rank int
}

func find(subsets []subset, i int) int { 
    if (subsets[i].parent != i) {
        subsets[i].parent = find(subsets, subsets[i].parent)
    }
    return subsets[i].parent; 
}

func Union(subsets []subset, x, y int) { 
    xroot := find(subsets, x); 
    yroot := find(subsets, y); 
  
    if subsets[xroot].rank < subsets[yroot].rank {
        subsets[xroot].parent = yroot;    
    }else if subsets[xroot].rank > subsets[yroot].rank {
        subsets[yroot].parent = xroot; 
    }else{
        subsets[yroot].parent = xroot; 
        subsets[xroot].rank = subsets[xroot].rank + 1; 
    }
}

func makeConnected(n int, connections [][]int) int {
    
    if n-1 > len(connections) {
        return -1
    }
    
    subsets := make([]subset, n , n)
    for i := 0 ; i<n ; i++{
        subsets[i].parent = i
        subsets[i].rank = 0
    }
    
    for _, connection := range connections {
        x := find(subsets, connection[0])
        y := find(subsets, connection[1])
        Union(subsets, x, y)
    }
    
    disjoint := 0
    for i := 0 ; i < n; i++{
        if subsets[i].parent == i {
            disjoint++
        }
    }
    return disjoint - 1
}
