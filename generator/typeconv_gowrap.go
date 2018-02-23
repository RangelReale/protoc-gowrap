package generator

import (
	"fmt"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type Converter_GoWrap struct {
	TypeNames []string
}

func (c *Converter_GoWrap) GetTypeConverter(g *Generator, message *Descriptor, field *descriptor.FieldDescriptorProto) TypeConverter {
	if *field.Type == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
		desc := g.ObjectNamed(field.GetTypeName())

		d, ok := desc.(*Descriptor)
		if !ok {
			return nil
		}

		var typename string
		if d.File().Package != nil {
			typename = fmt.Sprintf("%s.%s", *d.File().Package, strings.Join(desc.TypeName(), "."))
		} else {
			typename = fmt.Sprintf("%s.%s", desc.PackageName(), strings.Join(desc.TypeName(), "."))
		}

		for _, t := range c.TypeNames {
			if t == typename {
				rgt, _ := g.RawGoType(message, field, nil)
				if strings.HasPrefix(rgt, "*") {
					rgt = strings.TrimPrefix(rgt, "*")
				}

				return &TypeConverter_GoWrap{
					rawGoType: rgt + "_GW",
				}
				break
			}
		}
	}
	return nil
}

type TypeConverter_GoWrap struct {
	rawGoType string
}

func (t *TypeConverter_GoWrap) GoType() (typ string, wire string) {
	return "*" + t.rawGoType, "byte"
}

func (t *TypeConverter_GoWrap) Imports() []TypeImport {
	return []TypeImport{
		/*
			{
				Alias: "tc_uuid",
				Path:  "github.com/RangelReale/go.uuid",
			},*/
	}
}

func (t *TypeConverter_GoWrap) EmptyValue() string {
	return "&" + t.rawGoType + "{}"
}

func (t *TypeConverter_GoWrap) GenerateImport(g *Generator, fieldname string, retName string, varName string) {
	g.P("{")

	g.P("i_imp, err := ", t.rawGoType, "_Import(", varName, ".", fieldname, ")")
	g.P("if err != nil {")
	g.P("return nil, err")
	g.P("}")
	g.P(retName, ".", fieldname, " = i_imp")

	g.P("}")
}

func (t *TypeConverter_GoWrap) GenerateExport(g *Generator, fieldname string, structName string, varName string) {
	g.P("{")

	g.P("i_exp, err := ", structName, ".", fieldname, ".Export()")
	g.P("if err != nil {")
	g.P("return nil, err")
	g.P("}")
	g.P(varName, ".", fieldname, " = i_exp")

	g.P("}")
}

func (t *TypeConverter_GoWrap) RecordTypeUse() bool {
	return true
}
