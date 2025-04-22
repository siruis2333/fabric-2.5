package analysis

import (
	"fmt"
)

// TopologicalSort performs topological sorting on the call graph.
func TopologicalSort(graph FunctionCallGraph) ([]string, error) {
	visited := make(map[string]bool)
	temp := make(map[string]bool)
	var result []string
	var visit func(string) error

	visit = func(n string) error {
		if temp[n] {
			return fmt.Errorf("cycle detected at %s", n)
		}
		if !visited[n] {
			temp[n] = true
			for _, dep := range graph[n] {
				if err := visit(dep); err != nil {
					return err
				}
			}
			visited[n] = true
			temp[n] = false
			result = append(result, n)
		}
		return nil
	}

	for n := range graph {
		if err := visit(n); err != nil {
			return nil, err
		}
	}
	return result, nil
}
