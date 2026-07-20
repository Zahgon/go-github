package extraneousnew

import (
	"go/ast"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("extraneousnew", New)
}

type ExtraneousNewPlugin struct {
	ignoredMethods map[string]bool
}

func New(cfg any) (register.LinterPlugin, error) {
	_ = "STUB: not implemented"
	return *new(register.LinterPlugin), nil
}

func (f *ExtraneousNewPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *ExtraneousNewPlugin) GetLoadMode() string { _ = "STUB: not implemented"; return "" }

func run(pass *analysis.Pass, ignoredMethods map[string]bool) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func inspectAllBlocks(pass *analysis.Pass, root ast.Node) { _ = "STUB: not implemented"; return }

func inspectBlock(pass *analysis.Pass, block *ast.BlockStmt) { _ = "STUB: not implemented"; return }

func getFunctionName(expr ast.Expr) string { _ = "STUB: not implemented"; return "" }

func lookAhead(pass *analysis.Pass, block *ast.BlockStmt, startIndex int, lhsIdent *ast.Ident, typeName string) {
	_ = "STUB: not implemented"
	return
}

type valueVarInfo struct {
	ident     *ast.Ident
	typeName  string
	stmtIndex int
}

func isUsedElsewhere(block *ast.BlockStmt, declIndex, useIndex int, name string) bool {
	_ = "STUB: not implemented"
	return false
}

func isReadAfterCall(block *ast.BlockStmt, callIndex int, name string) bool {
	_ = "STUB: not implemented"
	return false
}

func isIdentOrAddressOfIdent(expr ast.Expr, name string) bool {
	_ = "STUB: not implemented"
	return false
}
