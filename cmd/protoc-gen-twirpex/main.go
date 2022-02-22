package main

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/albenik/twirpex/internal/generator"
)

var version string // set via -ldflags -X

func main() {
	pgen := &protogen.Options{}
	pgen.Run(func(plugin *protogen.Plugin) error {
		for _, file := range plugin.Files {
			if !file.Generate {
				continue
			}

			for _, service := range file.Services {
				filename := file.GeneratedFilenamePrefix + ".twirpex.go"
				g := generator.New(plugin.NewGeneratedFile(filename, file.GoImportPath), version)
				if err := g.Generate(file, service); err != nil {
					return fmt.Errorf("%s: %s schema: %w", file.Desc.Path(), service.Desc.FullName(), err)
				}
			}
		}

		return nil
	})
}
