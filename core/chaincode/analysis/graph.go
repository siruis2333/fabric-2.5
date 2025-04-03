/*package main

import (
	"fmt"
)

// Graph represents function dependencies
type Graph struct {
	AdjList   map[string][]string // Function -> Dependencies
	InDegree  map[string]int      // In-degree (number of incoming edges)
	Functions []string            // All function names
}

// NewGraph initializes an empty graph
func NewGraph() *Graph {
	return &Graph{
		AdjList:  make(map[string][]string),
		InDegree: make(map[string]int),
	}
}

// BuildGraph constructs the dependency graph from a call graph
func (g *Graph) BuildGraph(callGraph *CallGraph) {
	// Ensure all functions exist in the graph, even if they don't call others
	for funcName, calls := range callGraph.Calls {
		if _, exists := g.AdjList[funcName]; !exists {
			g.AdjList[funcName] = []string{}
		}
		if _, exists := g.InDegree[funcName]; !exists {
			g.InDegree[funcName] = 0
		}
		for _, calledFunc := range calls {
			g.AdjList[funcName] = append(g.AdjList[funcName], calledFunc)
			g.InDegree[calledFunc]++ // Increase in-degree
		}
		g.Functions = append(g.Functions, funcName)
	}
}*/

// Groups functions into shards based on common dependencies
package main

import "fmt"

// Graph represents function dependencies
type Graph struct {
	AdjList   map[string][]string // Function -> Dependencies
	InDegree  map[string]int      // In-degree (number of incoming edges)
	Functions []string            // All function names
}

// NewGraph initializes an empty graph
func NewGraph() *Graph {
	return &Graph{
		AdjList:  make(map[string][]string),
		InDegree: make(map[string]int),
	}
}

// CallGraph represents the function call relationships
type CallGraph struct {
	Calls map[string][]string
}

// PrintGraph displays the dependency graph
func (g *Graph) PrintGraph() {
	fmt.Println("Dependency Graph:")
	for funcName, deps := range g.AdjList {
		fmt.Printf("%s -> %v\n", funcName, deps)
	}
}
