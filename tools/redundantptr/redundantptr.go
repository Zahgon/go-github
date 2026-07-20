package redundantptr

import (
	"go/ast"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("redundantptr", New)
}

type RedundantPtrPlugin struct{}

func New(_ any) (register.LinterPlugin, error) {
	_ = "STUB: not implemented"
	return *new(register.LinterPlugin), nil
}

func (p *RedundantPtrPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *RedundantPtrPlugin) GetLoadMode() string { _ = "STUB: not implemented"; return "" }

func run(pass *analysis.Pass) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func analyzeFunction(pass *analysis.Pass, fn ast.Node, body *ast.BlockStmt) {
	_ = "STUB: not implemented"
	return
}

func redundantPtrCall(pass *analysis.Pass, call *ast.CallExpr) (rootName, argText string) {
	_ = "STUB: not implemented"
	return "", ""
}

func rootIdentName(expr ast.Expr) string { _ = "STUB: not implemented"; return "" }

func exprString(pass *analysis.Pass, expr ast.Expr) string { _ = "STUB: not implemented"; return "" }

func collectLocals(fn ast.Node, body *ast.BlockStmt) map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

func addFieldListNames(locals map[string]bool, fields *ast.FieldList) {
	_ = "STUB: not implemented"
	return
}

func shouldIgnoreDeprecatedPtrWrapper(fn ast.Node) bool { _ = "STUB: not implemented"; return false }
