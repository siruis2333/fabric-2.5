/*
package main

import (

	"fmt"
	"sort"

)

// OrderShardFunctions processes each shard, sorting functions based on dependencies

	func OrderShardFunctions(shard []string, callGraph *Graph) []string {
		sortedFunctions := []string{}
		inDegree := make(map[string]int)
		adjList := make(map[string][]string)

		// Build dependency graph within the shard
		for _, function := range shard {
			if deps, exists := callGraph.AdjList[function]; exists {
				adjList[function] = deps
				for _, dep := range deps {
					inDegree[dep]++
				}
			}
		}

		// Collect functions with zero in-degree (independent functions)
		queue := []string{}
		for _, function := range shard {
			if inDegree[function] == 0 {
				queue = append(queue, function)
			}
		}

		// Topological sorting
		for len(queue) > 0 {
			sort.Strings(queue) // Ensures deterministic order
			current := queue[0]
			queue = queue[1:]
			sortedFunctions = append(sortedFunctions, current)

			for _, neighbor := range adjList[current] {
				inDegree[neighbor]--
				if inDegree[neighbor] == 0 {
					queue = append(queue, neighbor)
				}
			}
		}

		return sortedFunctions
	}

// ProcessShards sorts and prints execution order for each shard

	func ProcessShards(shards [][]string, callGraph *Graph) {
		for i, shard := range shards {
			ordered := OrderShardFunctions(shard, callGraph)
			fmt.Printf("Shard %d execution order: %v\n", i+1, ordered)
		}
	}
*/
package main

import (
	"fmt"
	"sort"
)

// OrderShardFunctions processes each shard, sorting functions based on dependencies
func OrderShardFunctions(shard []string, callGraph *Graph) []string {
	sortedFunctions := []string{}
	inDegree := make(map[string]int)
	adjList := make(map[string][]string)

	// Build dependency graph within the shard
	for _, function := range shard {
		if deps, exists := callGraph.AdjList[function]; exists {
			adjList[function] = deps
			for _, dep := range deps {
				inDegree[dep]++
			}
		}
	}

	// Collect functions with zero in-degree (independent functions)
	queue := []string{}
	for _, function := range shard {
		if inDegree[function] == 0 {
			queue = append(queue, function)
		}
	}

	// Topological sorting
	for len(queue) > 0 {
		sort.Strings(queue) // Ensures deterministic order
		current := queue[0]
		queue = queue[1:]
		sortedFunctions = append(sortedFunctions, current)

		for _, neighbor := range adjList[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return sortedFunctions
}

// ProcessShards orders functions within each shard
func ProcessShards(shards [][]string, callGraph *Graph) {
	for i, shard := range shards {
		ordered := OrderShardFunctions(shard, callGraph)
		fmt.Printf("Shard %d execution order: %v\n", i+1, ordered)
	}
}
