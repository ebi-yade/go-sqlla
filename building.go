package genbase

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"strings"
)

type Generator struct {
	Package *PackageInfo

	Buf             bytes.Buffer // Accumulated output.
	RequiredImports []*Import
}

type Import struct {
	Ident string // e.g. "gb"
	Path  string // e.g. "github.com/favclip/genbase"
}

func NewGenerator(pkg *PackageInfo) *Generator {
	return &Generator{
		Package: pkg,
	}
}

func (g *Generator) Printf(format string, args ...interface{}) {
	g.Buf.WriteString(fmt.Sprintf(format, args...))
}

func (g *Generator) AddImport(path string, ident string) {
	if strings.HasPrefix(path, `"`) && strings.HasSuffix(path, `"`) {
		path = path[1 : len(path)-1]
	}
	g.RequiredImports = append(g.RequiredImports, &Import{Ident: ident, Path: path})
}

func (g *Generator) PrintHeader(cmdName string, args *[]string) {
	if cmdName == "" && args != nil {
		cmdName = (*args)[0]
	} else if cmdName == "" {
		cmdName = "???"
	}

	var as []string
	if args == nil {
		as = os.Args[1:]
	} else {
		as = (*args)
	}

	// Print the header and package clause.
	g.Printf("// generated by %s %s; DO NOT EDIT\n", cmdName, strings.Join(as, " "))
	g.Printf("\n")
	g.Printf("package %s\n", g.Package.Name())
	g.Printf("import (\n")
	for _, imp := range g.RequiredImports {
		g.Printf("%s \"%s\"\n", imp.Ident, imp.Path)
	}
	g.Printf(")\n")
}

func (g *Generator) Format() ([]byte, error) {
	src, err := format.Source(g.Buf.Bytes())
	if err != nil {
		return g.Buf.Bytes(), err
	}

	return src, nil
}
