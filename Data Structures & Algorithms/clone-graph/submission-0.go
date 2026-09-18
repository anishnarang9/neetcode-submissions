/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil{
		return nil 
	}

	clones := make(map[*Node]*Node)
	queue := []*Node{}

	clones[node] = &Node{Val: node.Val}

	queue = append(queue,node)

	for len(queue) > 0  {
		n := queue[0]
		queue = queue[1:]

		for _,v := range n.Neighbors{
			if _,exists := clones[v]; !exists{
				clones[v] = &Node{Val: v.Val}
				queue = append(queue,v)
			}
			clones[n].Neighbors = append(clones[n].Neighbors,clones[v])



		}


	}
	return clones[node]



    
}
