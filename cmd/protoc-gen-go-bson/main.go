package main

import (
	"google.golang.org/protobuf/compiler/protogen"
)

const filenameSuffix = "_bson.pb.go"

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f)
		}
		return nil;
	})
}

func generateFile(gen *protogen.Plugin, file *protogen.File) {
	filename := file.GeneratedFilenamePrefix + filenameSuffix
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P(("// CODE GENERATED by protoc-bson-gen. DO NOT EDIT."))
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
}