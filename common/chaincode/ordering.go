package analysis

import (
	"sort"
)

type ShardOrder struct {
	ShardID int
	Order   []string
}

// OrderWithinShards performs topological sort for each shard.
func OrderWithinShards(shards map[int][]string, graph FunctionCallGraph) []ShardOrder {
	var result []ShardOrder
	for id, funcs := range shards {
		subgraph := make(FunctionCallGraph)
		for _, fn := range funcs {
			subgraph[fn] = graph[fn]
		}
		order, err := TopologicalSort(subgraph)
		if err != nil {
			// fallback: use lexicographical order
			sort.Strings(funcs)
			order = funcs
		}
		result = append(result, ShardOrder{ShardID: id, Order: order})
	}
	return result
}
