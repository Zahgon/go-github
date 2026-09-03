package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/token"
	"log"
)

func main() {
	log.SetFlags(0)

	omitempty := flag.Bool("omitempty", false, "reduce the report to structs that have at least one field tagged with \"omitempty\" or \"omitzero\"")
	flag.Usage = func() {
		log.Print("Usage: list-return-structs [-omitempty] [dir]")
		flag.PrintDefaults()
	}
	flag.Parse()

	dir := "github"
	if flag.NArg() > 0 {
		dir = flag.Arg(0)
	}

	structs, err := analyze(dir, *omitempty)
	if err != nil {
		log.Fatalf("analyzing %s: %v", dir, err)
	}

	for _, s := range structs {
		fmt.Printf("%v:%v:type %v struct {\n", s.relPath, s.line, s.name)
	}
}

type structInfo struct {
	name       string
	relPath    string
	line       int
	structType *ast.StructType
}

func parsePackage(dir string) (*token.FileSet, []*ast.File, string, error) {
	_ = "STUB: not implemented"
	return nil, nil, "", nil
}

func collectStructs(fset *token.FileSet, files []*ast.File, baseDir string) map[string]*structInfo {
	_ = "STUB: not implemented"
	return nil
}

func collectDirectUses(files []*ast.File) (returns, inputs map[string]bool) {
	_ = "STUB: not implemented"
	return nil, nil
}

func transitiveClosure(start map[string]bool, structs map[string]*structInfo) map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

func baseStructName(typ ast.Expr) string { _ = "STUB: not implemented"; return "" }

func hasJSONOmitempty(st *ast.StructType) bool { _ = "STUB: not implemented"; return false }

func structTagLookup(tag, key string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func unquoteStructTag(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func analyze(dir string, omitempty bool) ([]*structInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
