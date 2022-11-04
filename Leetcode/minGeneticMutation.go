// https://leetcode.com/problems/minimum-genetic-mutation

package leetcode

var (
    allowedTransitions = []byte{'A', 'C', 'G', 'T'}
)

type QueueNode struct {
    gene string
    level int
}

func minMutation(startGene string, endGene string, bank []string) int {
    bankMap := getBankMap(bank)
    bankMap[startGene] = true
    
    seen := make(map[string]bool)
    
    queue := make([]QueueNode, 0, 0)
    queue = append(queue, QueueNode{gene: startGene, level: 0})
    
    for len(queue) > 0 {
        n := queue[0]
        queue = queue[1:]
        
        if n.gene == endGene {
            return n.level
        }
        
        for i := 0 ; i < len(n.gene); i++ {
            s := n.gene
            for _, ch := range allowedTransitions {
                s1 := replaceAtIndex(s, rune(ch), i)
                _, saw := seen[s1]
                _, exists := bankMap[s1]
                if exists && !saw{
                        queue = append(queue, QueueNode{gene: s1, level: n.level+1})
                        seen[s1] = true
                }
            }
        }
    }
    
    return -1
}

func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
}

func getBankMap(bank []string)map[string]bool{
    
    bm := make(map[string]bool)   
    for _, s := range bank {
        bm[s] = true
    }
    return bm
}