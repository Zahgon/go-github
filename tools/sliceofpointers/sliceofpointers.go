package sliceofpointers

import (
	"go/ast"
	"go/token"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("sliceofpointers", New)
}

type SliceOfPointersPlugin struct{}

func New(_ any) (register.LinterPlugin, error) {
	_ = "STUB: not implemented"
	return *new(register.LinterPlugin), nil
}

func (f *SliceOfPointersPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *SliceOfPointersPlugin) GetLoadMode() string { _ = "STUB: not implemented"; return "" }

func run(pass *analysis.Pass) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func checkArrayType(arrType *ast.ArrayType, tokenPos token.Pos, pass *analysis.Pass) {
	_ = "STUB: not implemented"
	return
}
