//go:build ignore

package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"text/template"
)

const (
	fileSuffix = "-iterators.go"
)

var (
	check   = flag.Bool("check", false, "Check whether generated files are up to date")
	verbose = flag.Bool("v", false, "Print verbose log messages")

	sourceTmpl = template.Must(template.New("source").Funcs(template.FuncMap{
		"hasPrefix": strings.HasPrefix,
	}).Parse(source))

	testTmpl = template.Must(template.New("test").Parse(test))
)

func isCheck() bool { _ = "STUB: not implemented"; return false }

func logf(fmt string, args ...any) { _ = "STUB: not implemented"; return }

func main() {
	flag.Parse()
	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, ".", sourceFilter, 0)
	if err != nil {
		log.Fatal(err)
		return
	}

	for pkgName, pkg := range pkgs {
		t := &templateData{
			filename: pkgName + fileSuffix,
			Package:  pkgName,
			Methods:  []*method{},
			Structs:  make(map[string]*structDef),
		}

		for _, f := range pkg.Files {
			t.processStructs(f)
		}

		for _, f := range pkg.Files {
			if err := t.processMethods(f); err != nil {
				log.Fatal(err)
			}
		}

		if err := t.dump(); err != nil {
			log.Fatal(err)
		}
	}
	logf("Done.")
}

func sourceFilter(fi os.FileInfo) bool { _ = "STUB: not implemented"; return false }

type templateData struct {
	filename string
	Package  string
	Methods  []*method
	Structs  map[string]*structDef
}

type structDef struct {
	Name      string
	Fields    map[string]string
	FieldJSON map[string]string
	Embeds    []string
}

type method struct {
	RecvType             string
	RecvVar              string
	ClientField          string
	MethodName           string
	IterMethod           string
	Args                 string
	CallArgs             string
	TestCallArgs         string
	ZeroArgs             string
	ReturnType           string
	OptsType             string
	OptsName             string
	OptsIsPtr            bool
	UseListCursorOptions bool
	UseListOptions       bool
	UsePage              bool
	UseAfter             bool
	UseCursor            bool
	WrappedItemsField    string
	TestJSON1            string
	TestJSON2            string
	TestJSON3            string
}

type methodInfo struct {
	RecvTypeRaw          string
	RecvType             string
	RecvVar              string
	ClientField          string
	Args                 string
	CallArgs             string
	TestCallArgs         string
	ZeroArgs             string
	OptsType             string
	OptsName             string
	OptsIsPtr            bool
	UseListCursorOptions bool
	UseListOptions       bool
	UsePage              bool
	UseAfter             bool
	UseCursor            bool
}

var useCursorPagination = map[string]bool{
	"AppsService.ListHookDeliveries":          true,
	"OrganizationsService.ListHookDeliveries": true,
	"RepositoriesService.ListHookDeliveries":  true,
}

var customNames = map[string]string{
	"RepositoriesService.GetCommit":         "ListCommitFiles",
	"RepositoriesService.CompareCommits":    "ListCommitComparisonFiles",
	"RepositoriesService.GetCombinedStatus": "ListCombinedStatus",
}

var sliceToBeUsedForIteration = map[string]string{
	"RepositoriesService.GetCommit":      "Files",
	"RepositoriesService.CompareCommits": "Files",
}

var customTestJSON = map[string]string{
	"ListAllTopics":         `{"names": []}`,
	"ListUserInstallations": `{"installations": []}`,
}

func (t *templateData) processStructs(f *ast.File) { _ = "STUB: not implemented"; return }

func (t *templateData) hasListCursorOptions(structName string) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *templateData) hasListOptions(structName string) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *templateData) hasOptions(structName, optionsType string) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *templateData) hasIntPage(structName string) bool { _ = "STUB: not implemented"; return false }

func (t *templateData) hasStringAfter(structName string) bool {
	_ = "STUB: not implemented"
	return false
}

func getZeroValue(typeStr string) string { _ = "STUB: not implemented"; return "" }

func (t *templateData) processMethods(f *ast.File) error { _ = "STUB: not implemented"; return nil }

func (t *templateData) isMethodIterable(fd *ast.FuncDecl) (*methodInfo, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func validateMethodShape(fd *ast.FuncDecl) bool { _ = "STUB: not implemented"; return false }

func (t *templateData) collectMethodInfo(fd *ast.FuncDecl) (*methodInfo, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func getIterName(methodInfo *methodInfo, methodName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (t *templateData) processReturnArrayType(fd *ast.FuncDecl, sliceRet *ast.ArrayType, methodInfo *methodInfo) {
	_ = "STUB: not implemented"
	return
}

func (t *templateData) processReturnStarExpr(fd *ast.FuncDecl, starRet *ast.StarExpr, methodInfo *methodInfo) {
	_ = "STUB: not implemented"
	return
}

func findSinglePointerSliceField(sd *structDef) (fieldName, fieldType string, ok bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

func lowerFirst(s string) string { _ = "STUB: not implemented"; return "" }

func typeToString(expr ast.Expr) string { _ = "STUB: not implemented"; return "" }

func (t *templateData) dump() error { _ = "STUB: not implemented"; return nil }

const doNotEditHeader = `// Code generated by gen-iterators; DO NOT EDIT.
// Instead, please run "go generate ./..." as described here:
// https://github.com/google/go-github/blob/master/CONTRIBUTING.md#submitting-a-patch

// Copyright 2026 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
`

const source = doNotEditHeader + `
package {{.Package}}

import (
	"context"
	"iter"
)

{{range .Methods}}
// {{.IterMethod}} returns an iterator that paginates through all results of {{.MethodName}}.
func ({{.RecvVar}} *{{.RecvType}}) {{.IterMethod}}({{.Args}}) iter.Seq2[{{.ReturnType}}, error] {
	return func(yield func({{.ReturnType}}, error) bool) {
		{{if .OptsIsPtr -}}
		// Create a copy of opts to avoid mutating the caller's struct
		if {{.OptsName}} == nil {
			{{.OptsName}} = &{{.OptsType}}{}
		} else {
			{{.OptsName}} = Ptr(*{{.OptsName}})
		}

		{{end}}
		for {
			results, resp, err := {{.RecvVar}}.{{.MethodName}}({{.CallArgs}})
			if err != nil {
				yield({{if hasPrefix .ReturnType "*"}}nil{{else}}*new({{.ReturnType}}){{end}}, err)
				return
			}

			{{if .WrappedItemsField -}}
			var iterItems []{{.ReturnType}}
			if results != nil {
				iterItems = results.{{.WrappedItemsField}}
			}
			for _, item := range iterItems {
			{{else -}}
			for _, item := range results {
			{{end -}}
				if !yield(item, nil) {
					return
				}
			}

			{{if and .UseListCursorOptions .UseListOptions}}
			if resp.After == "" && resp.NextPage == 0 {
				break
			}
			{{.OptsName}}.ListCursorOptions.After = resp.After
			{{.OptsName}}.ListOptions.Page = resp.NextPage
			{{else if .UseListCursorOptions}}
			if resp.After == "" {
				break
			}
			{{.OptsName}}.ListCursorOptions.After = resp.After
			{{else if .UseListOptions}}
			if resp.NextPage == 0 {
				break
			}
			{{.OptsName}}.ListOptions.Page = resp.NextPage
			{{else if .UsePage}}
			if resp.NextPage == 0 {
				break
			}
			{{.OptsName}}.Page = resp.NextPage
			{{else if .UseAfter}}
			if resp.After == "" {
				break
			}
			{{.OptsName}}.After = resp.After
			{{else if .UseCursor}}
			if resp.Cursor == "" {
				break
			}
			{{.OptsName}}.Cursor = resp.Cursor
			{{end -}}
		}
	}
}
{{end}}
`

const test = doNotEditHeader + `
package {{.Package}}

import (
	"fmt"
	"net/http"
	"testing"
)

{{range .Methods}}
func Test{{.RecvType}}_{{.IterMethod}}(t *testing.T) {
	t.Parallel()
	client, mux, _ := setup(t)
	var callNum int
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		callNum++
		switch callNum {
		case 1:
			{{- if .UseCursor}}
			w.Header().Set("Link", ` + "`" + `<https://api.github.com/?cursor=yo>; rel="next"` + "`" + `)
			{{else if or .UseListCursorOptions .UseAfter}}
			w.Header().Set("Link", ` + "`" + `<https://api.github.com/?after=yo>; rel="next"` + "`" + `)
			{{else}}
			w.Header().Set("Link", ` + "`" + `<https://api.github.com/?page=1>; rel="next"` + "`" + `)
			{{end -}}
			fmt.Fprint(w, ` + "`" + `{{.TestJSON1}}` + "`" + `)
		case 2:
			fmt.Fprint(w, ` + "`" + `{{.TestJSON2}}` + "`" + `)
		case 3:
			fmt.Fprint(w, ` + "`" + `{{.TestJSON3}}` + "`" + `)
		case 4:
			w.WriteHeader(http.StatusNotFound)
		case 5:
			fmt.Fprint(w, ` + "`" + `{{.TestJSON3}}` + "`" + `)
		}
	})

	iter := client.{{.ClientField}}.{{.IterMethod}}({{.ZeroArgs}})
	var gotItems int
	for _, err := range iter {
		gotItems++
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}
	if want := 7; gotItems != want {
		t.Errorf("client.{{.ClientField}}.{{.IterMethod}} call 1 got %v items; want %v", gotItems, want)
	}

	{{.OptsName}} := &{{.OptsType}}{}
	iter = client.{{.ClientField}}.{{.IterMethod}}({{.TestCallArgs}})
	gotItems = 0
	for _, err := range iter {
		gotItems++
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}
	if want := 2; gotItems != want {
		t.Errorf("client.{{.ClientField}}.{{.IterMethod}} call 2 got %v items; want %v", gotItems, want)
	}

	iter = client.{{.ClientField}}.{{.IterMethod}}({{.ZeroArgs}})
	gotItems = 0
	for _, err := range iter {
		gotItems++
		if err == nil {
			t.Error("expected error; got nil")
		}
	}
	if gotItems != 1 {
		t.Errorf("client.{{.ClientField}}.{{.IterMethod}} call 3 got %v items; want 1 (an error)", gotItems)
	}

	iter = client.{{.ClientField}}.{{.IterMethod}}({{.ZeroArgs}})
	gotItems = 0
	iter(func(item {{.ReturnType}}, err error) bool {
		gotItems++
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		return false
	})
	if gotItems != 1 {
		t.Errorf("client.{{.ClientField}}.{{.IterMethod}} call 4 got %v items; want 1 (an error)", gotItems)
	}
}
{{end}}
`
