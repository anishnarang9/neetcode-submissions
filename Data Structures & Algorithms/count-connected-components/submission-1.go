func countComponents(n int, edges [][]int) int {

	visited := make([]bool,n)
	adj := make([][]int,n)
	comps := 0

	for _, edge := range edges{
		a := edge[0]
		b := edge[1]

		adj[a] = append(adj[a],b)
		adj[b] = append(adj[b],a)

	}

	var dfs func(node int)

	dfs = func(node int){
		if visited[node]{
			return 
		}
		visited[node] = true 

		for _,nei := range adj[node]{
			dfs(nei)
		}

	}
	for node:=0; node < n; node++{
		if !visited[node]{
			comps ++
			dfs(node)
		}
	}
	return comps 
    
}
