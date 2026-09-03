package fmtpercentv

import (
	"go/ast"
	"go/token"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("fmtpercentv", New)
}

type FmtPercentVPlugin struct{}

func New(_ any) (register.LinterPlugin, error) {
	_ = "STUB: not implemented"
	return *new(register.LinterPlugin), nil
}

func (f *FmtPercentVPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FmtPercentVPlugin) GetLoadMode() string { _ = "STUB: not implemented"; return "" }

func run(pass *analysis.Pass) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func checkCallExpr(expr *ast.CallExpr, tokenPos token.Pos, pass *analysis.Pass) {
	_ = "STUB: not implemented"
	return
}
