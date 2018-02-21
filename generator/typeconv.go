package generator

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type Converter interface {
	GetTypeConverter(g *Generator, message *Descriptor, field *descriptor.FieldDescriptorProto) TypeConverter
}

type TypeConverter interface {
	GoType() (typ string, wire string)
	Imports() []TypeImport
	EmptyValue() string
	GenerateImport(g *Generator, fieldname string, structName string, varName string)
	GenerateExport(g *Generator, fieldname string, structName string, varName string)
}

type TypeImport struct {
	Alias string
	Path  string
}
