package generator

import (
	"io"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/albenik/twirpex/internal/template"
)

type Generator struct {
	writer  io.Writer
	version string
}

func New(w io.Writer, ver string) *Generator {
	return &Generator{
		writer:  w,
		version: ver,
	}
}

func (g *Generator) Generate(file *protogen.File, service *protogen.Service) error {
	methods := make([]string, len(service.Methods))
	for i, m := range service.Methods {
		methods[i] = m.GoName
	}

	return template.Execute(g.writer, &template.Data{
		GeneratorVersion: g.version,
		GoPackage: string(file.GoPackageName),
		Proto: template.Proto{
			FileName:        file.Desc.Path(),
			PackageName:     string(service.Desc.Parent().FullName()),
			ServiceName:     string(service.Desc.Name()),
			ServiceFullName: string(service.Desc.FullName()),
			MethodsNames:    methods,
		},
	})
}
