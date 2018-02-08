package sqlla

import (
	"io"
	"io/ioutil"
	"log"
	"strings"
	"text/template"
	"unicode"

	_ "github.com/mackee/go-sqlla/statik"

	"github.com/rakyll/statik/fs"
	"github.com/serenize/snaker"
)

//go:generate statik -src=template -m

var templates = []string{
	"/table.tmpl",
	"/select.tmpl",
	"/select_column.tmpl",
	"/update.tmpl",
	"/update_column.tmpl",
	"/insert.tmpl",
	"/insert_column.tmpl",
	"/delete.tmpl",
	"/delete_column.tmpl",
}

var tmpl = template.New("sqlla")

func init() {
	tmpl = tmpl.Funcs(
		template.FuncMap{
			"Title": strings.Title,
			"Exprize": func(s string) string {
				s = strings.TrimPrefix(s, "sql.")
				s = strings.TrimPrefix(s, "time.")
				s = strings.TrimPrefix(s, "mysql.")

				return s
			},
			"Untitle": func(s string) string {
				s0 := rune(s[0])
				if !unicode.IsUpper(s0) {
					return s
				}
				s0l := unicode.ToLower(rune(s[0]))
				return string(s0l) + s[1:]
			},
			"toLower": strings.ToLower,
			"toSnake": snaker.CamelToSnake,
			"toCamel": snaker.SnakeToCamel,
		},
	)

	for _, filename := range templates {
		afs, err := fs.New()
		if err != nil {
			log.Fatalf("failed open bundled filesystem: %s", err)
		}

		af, err := afs.Open(filename)
		if err != nil {
			log.Fatalf("failed open bundled templates: %s", err)
		}

		bs, err := ioutil.ReadAll(af)
		if err != nil {
			log.Fatalf("failed read bundled templates: %s", err)
		}

		tmpl, err = tmpl.Parse(string(bs))
		if err != nil {
			log.Fatalf("failed parse bundled templates: %s", err)
		}
	}
}

func WriteCode(w io.Writer, table *Table) error {
	return tmpl.Execute(w, table)
}
