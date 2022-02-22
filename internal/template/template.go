package template

import (
	"io"
	"text/template"
	"unicode"
)

const tpl = `// Code generated by protoc-gen-twirpex {{ .GeneratorVersion }}, DO NOT EDIT.
// source: {{ .Proto.FileName }}

package {{ .GoPackage }}

import (
	"github.com/albenik/twirpex"
)

func New{{ .Proto.ServiceName }}ServerEx(svc {{ .Proto.ServiceName }}, opts ...interface{}) twirpex.TwirpServer {
	return New{{ .Proto.ServiceName }}Server(svc, opts...).(*{{ .Proto.ServiceName | lcfirst }}Server)
}

func (*{{ .Proto.ServiceName | lcfirst }}Server) TwirpServiceMeta() *twirpex.ServiceMeta {
	return &twirpex.ServiceMeta {
		PackageName:     "{{ .Proto.PackageName }}",
		ServiceName:     "{{ .Proto.ServiceName }}",
		ServiceFullName: "{{ .Proto.ServiceFullName }}",
		MethodsNames:    []string {
		{{- range .Proto.MethodsNames }}
			"{{ . }}",
		{{- end }}
		},
	}
}
`

var t = template.Must(template.New("src").Funcs(template.FuncMap{
	"lcfirst": func(s string) string {
		r := []rune(s)
		r[0] = unicode.ToLower(r[0])
		return string(r)
	},
}).Parse(tpl))

type Data struct {
	GeneratorVersion string
	GoPackage        string
	Proto            Proto
}

type Proto struct {
	FileName        string
	PackageName     string
	ServiceName     string
	ServiceFullName string
	MethodsNames    []string
}

func Execute(w io.Writer, data *Data) error {
	return t.Execute(w, data)
}
