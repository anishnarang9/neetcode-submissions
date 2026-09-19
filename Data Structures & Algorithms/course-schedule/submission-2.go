func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int,numCourses)

	for _,pre := range prerequisites{
		course := pre[0]
		prereq := pre[1]

		adj[prereq] = append(adj[prereq], course)

		
	}
	visited := make([]bool, numCourses)
	path := make ([]bool, numCourses)

	var hasCycle func (course int)bool

	hasCycle = func (course int ) bool {
		if path[course] {
			return true
		}
		if visited[course]{
			return false 
		}
		visited[course] = true 
		path[course] = true 

		for _, next := range adj[course]{
			if hasCycle(next){
				return true 
			}
		}
		path[course] = false
		return false  

	}
	for course := 0; course < numCourses; course++{
		if !visited[course]{
			if hasCycle(course) {
				return false 
			}
		}
	}
	return true 
    
}
