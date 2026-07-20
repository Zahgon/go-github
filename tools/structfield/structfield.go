package structfield

import (
	"go/ast"
	"go/token"
	"reflect"
	"regexp"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("structfield", New)
}

type StructFieldPlugin struct {
	allowedTagNames map[string]bool
	allowedTagTypes map[string]bool
}

type Settings struct {
	AllowedTagNames []string `json:"allowed-tag-names" yaml:"allowed-tag-names"`
	AllowedTagTypes []string `json:"allowed-tag-types" yaml:"allowed-tag-types"`
}

func New(cfg any) (register.LinterPlugin, error) {
	_ = "STUB: not implemented"
	return *new(register.LinterPlugin), nil
}

func (f *StructFieldPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *StructFieldPlugin) GetLoadMode() string { _ = "STUB: not implemented"; return "" }

func run(pass *analysis.Pass, allowedTagNames, allowedTagTypes map[string]bool) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func processStructField(structName string, field *ast.Field, pass *analysis.Pass, allowedTagNames, allowedTagTypes map[string]bool) {
	_ = "STUB: not implemented"
	return
}

func processTag(structName string, goField *ast.Ident, field *ast.Field, structTag reflect.StructTag, tagType string, pass *analysis.Pass, allowedTagNames, allowedTagTypes map[string]bool) {
	_ = "STUB: not implemented"
	return
}

func checkAndReportInvalidTypesForOmitzero(structName, tagType, goFieldName string, fieldType ast.Expr, tokenPos token.Pos, pass *analysis.Pass) bool {
	_ = "STUB: not implemented"
	return false
}

func checkGoFieldName(structName, goFieldName, tagType, tagName string, tokenPos token.Pos, pass *analysis.Pass, allowedNames map[string]bool) {
	_ = "STUB: not implemented"
	return
}

func checkGoFieldType(structName, goFieldName, tagType string, field *ast.Field, tokenPos token.Pos, pass *analysis.Pass, allowedTypes map[string]bool, omitempty, omitzero bool) {
	_ = "STUB: not implemented"
	return
}

func checkAndReportInvalidTypes(structName, tagType, goFieldName string, fieldType ast.Expr, tokenPos token.Pos, pass *analysis.Pass) (newFieldType string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}

func checkStructArrayType(structName, goFieldName string, arrType *ast.ArrayType, tokenPos token.Pos, pass *analysis.Pass) {
	_ = "STUB: not implemented"
	return
}

func isBuiltinType(typeName string) bool { _ = "STUB: not implemented"; return false }

func exprToString(e ast.Expr) string { _ = "STUB: not implemented"; return "" }

func splitTag(jsonTagName string) []string { _ = "STUB: not implemented"; return nil }

var camelCaseRE = regexp.MustCompile(`([a-z0-9])([A-Z])`)

func tagNameToPascal(tagName string) (want, alternate string) {
	_ = "STUB: not implemented"
	return "", ""
}

var initialisms = map[string]bool{
	"AI": true, "API": true, "ASCII": true, "AWS": true,
	"CAA": true, "CAS": true, "CLI": true, "CNAME": true, "CPU": true,
	"CSS": true, "CWE": true, "CVE": true, "CVSS": true,
	"DN": true, "DNS": true,
	"EOF": true, "EPSS": true,
	"GB": true, "GHSA": true, "GPG": true, "GUID": true,
	"HTML": true, "HTTP": true, "HTTPS": true,
	"ID": true, "IDE": true, "IDP": true, "IP": true, "JIT": true,
	"JSON": true,
	"OIDC": true,
	"LDAP": true, "LFS": true, "LHS": true, "LOC": true,
	"MCP": true, "MD5": true, "MS": true, "MX": true,
	"NPM": true, "NTP": true, "NVD": true,
	"OID": true, "OS": true,
	"PEM": true, "PR": true, "QPS": true,
	"RAM": true, "RHS": true, "RPC": true,
	"SAML": true, "SAS": true, "SBOM": true, "SCIM": true,
	"SHA": true, "SHA1": true, "SHA256": true,
	"SKU": true, "SLA": true, "SMTP": true, "SNMP": true,
	"SPDX": true, "SPDXID": true, "SQL": true, "SSH": true,
	"SSL": true, "SSO": true, "SVN": true,
	"TCP": true, "TFVC": true, "TLS": true, "TTL": true,
	"UDP": true, "UI": true, "UID": true, "UUID": true,
	"URI": true, "URL": true, "UTF8": true,
	"VCF": true, "VCS": true, "VM": true,
	"XML": true, "XMPP": true, "XSRF": true, "XSS": true,
}

var specialCases = map[string]string{
	"CPUS":    "CPUs",
	"CWES":    "CWEs",
	"JFROG":   "JFrog",
	"GRAPHQL": "GraphQL",
	"HREF":    "HRef",
	"IDS":     "IDs",
	"IPS":     "IPs",
	"OAUTH":   "OAuth",
	"OPENAPI": "OpenAPI",
	"URLS":    "URLs",
}

var possibleAlternates = map[string]string{
	"ORGANIZATION":  "Org",
	"ORGANIZATIONS": "Orgs",
	"REPOSITORY":    "Repo",
	"REPOSITORIES":  "Repos",
}
