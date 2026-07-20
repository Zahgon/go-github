package paramcheck

import (
	"go/ast"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("paramcheck", New)
}

type ParamCheckPlugin struct {
	bodyAllowedPointerTypes map[string]bool

	bodyAllowedWrongNames map[string]bool
}

func New(cfg any) (register.LinterPlugin, error) {
	_ = "STUB: not implemented"
	return *new(register.LinterPlugin), nil
}

func (p *ParamCheckPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *ParamCheckPlugin) GetLoadMode() string { _ = "STUB: not implemented"; return "" }

func (p *ParamCheckPlugin) run(pass *analysis.Pass) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (p *ParamCheckPlugin) analyzeFunc(pass *analysis.Pass, fn *ast.FuncDecl) {
	_ = "STUB: not implemented"
	return
}

func bodyReportName(pass *analysis.Pass, fn *ast.FuncDecl, name *ast.Ident) {
	_ = "STUB: not implemented"
	return
}

func queryReportName(pass *analysis.Pass, fn *ast.FuncDecl, name *ast.Ident) {
	_ = "STUB: not implemented"
	return
}

func (p *ParamCheckPlugin) bodyReportPass(pass *analysis.Pass, field *ast.Field, name *ast.Ident) {
	_ = "STUB: not implemented"
	return
}

func (p *ParamCheckPlugin) queryReportPass(pass *analysis.Pass, field *ast.Field, name *ast.Ident) {
	_ = "STUB: not implemented"
	return
}

func (p *ParamCheckPlugin) queryReportSuffix(pass *analysis.Pass, field *ast.Field) {
	_ = "STUB: not implemented"
	return
}

func (p *ParamCheckPlugin) bodyReportSuffix(pass *analysis.Pass, field *ast.Field) {
	_ = "STUB: not implemented"
	return
}

func isAddOptions(call *ast.CallExpr) bool { _ = "STUB: not implemented"; return false }

func isClientNewRequest(call *ast.CallExpr) bool { _ = "STUB: not implemented"; return false }

func isMutatingMethod(call *ast.CallExpr) bool { _ = "STUB: not implemented"; return false }

func findParam(fn *ast.FuncDecl, name string) (*ast.Field, *ast.Ident) {
	_ = "STUB: not implemented"
	return nil, nil
}

func typeNameIdent(expr ast.Expr) *ast.Ident { _ = "STUB: not implemented"; return nil }

func renameEdits(fn *ast.FuncDecl, old, newName string) []analysis.TextEdit {
	_ = "STUB: not implemented"
	return nil
}
