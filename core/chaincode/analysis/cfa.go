/*
package main

import (

	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"

)

// CallGraph store the dependency

	type CallGraph struct {
		Calls map[string][]string
	}

// NewCallGraph creat dependency graph

	func NewCallGraph() *CallGraph {
		return &CallGraph{
			Calls: make(map[string][]string),
		}
	}

// Analyze and parse the smart contract

	func (cg *CallGraph) Analyze(filePath string) error {
		// read smart contract
		src, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		// parse go
		fs := token.NewFileSet()
		node, err := parser.ParseFile(fs, "", src, parser.AllErrors)
		if err != nil {
			return err
		}

		// Traverse AST
		ast.Inspect(node, func(n ast.Node) bool {
			if fn, ok := n.(*ast.FuncDecl); ok {
				funcName := fn.Name.Name
				cg.Calls[funcName] = []string{} // Initialize the call list

				// find out the calls within the function body
				ast.Inspect(fn.Body, func(expr ast.Node) bool {
					if call, ok := expr.(*ast.CallExpr); ok {
						// Exclude function calls through pointers
						if _, ok := call.Fun.(*ast.Ident); ok {
							cg.Calls[funcName] = append(cg.Calls[funcName], call.Fun.(*ast.Ident).Name)
						}
					}
					return true
				})
			}
			return true
		})

		return nil
	}

// Print

	func (cg *CallGraph) Print() {
		for funcName, calls := range cg.Calls {
			fmt.Printf("%s calls: %s\n", funcName, strings.Join(calls, ", "))
		}
	}

// main function

	func main() {
		if len(os.Args) < 2 {
			fmt.Println("need path")
			return
		}
		filePath := os.Args[1]

		// Create a call graph
		callGraph := NewCallGraph()
		err := callGraph.Analyze(filePath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Print call graph
		callGraph.Print()

		// Initialize Graph for Sharding
		graph := NewGraph()
		graph.BuildGraph(callGraph)

		// Call sharding function
		fmt.Println("DEBUG: Calling PrintShards...")
		graph.PrintShards() // This ensures sharding is executed
		fmt.Println("DEBUG: Finished PrintShards...")

}
*/

package main
