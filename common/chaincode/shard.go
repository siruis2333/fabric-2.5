package analysis

// GroupByDependency groups functions into shards based on shared dependencies.
func GroupByDependency(graph FunctionCallGraph) map[int][]string {
	shards := make(map[int][]string)
	shardID := 0
	visited := make(map[string]bool)

	var dfs func(string, int)
	dfs = func(fn string, id int) {
		if visited[fn] {
			return
		}
		visited[fn] = true
		shards[id] = append(shards[id], fn)
		for _, dep := range graph[fn] {
			dfs(dep, id)
		}
	}

	for fn := range graph {
		if !visited[fn] {
			dfs(fn, shardID)
			shardID++
		}
	}

	return shards
}
