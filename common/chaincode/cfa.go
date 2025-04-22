package analysis

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type FunctionCallGraph map[string][]string

// AnalyzeSmartContract parses Go source files and constructs function-level call graph.
func AnalyzeSmartContract(dir string) (FunctionCallGraph, error) {
	graph := make(FunctionCallGraph)
	fs := token.NewFileSet()

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || !strings.HasSuffix(path, ".go") {
			return nil
		}

		node, err := parser.ParseFile(fs, path, nil, 0)
		if err != nil {
			return err
		}

		ast.Inspect(node, func(n ast.Node) bool {
			fn, ok := n.(*ast.FuncDecl)
			if !ok || fn.Body == nil {
				return true
			}
			caller := fn.Name.Name
			graph[caller] = []string{}
			ast.Inspect(fn.Body, func(n ast.Node) bool {
				call, ok := n.(*ast.CallExpr)
				if ok {
					if sel, ok := call.Fun.(*ast.Ident); ok {
						graph[caller] = append(graph[caller], sel.Name)
					}
				}
				return true
			})
			return true
		})
		return nil
	})

	return graph, err
}
