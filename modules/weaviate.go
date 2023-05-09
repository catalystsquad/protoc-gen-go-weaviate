package modules

import (
	"fmt"
	"github.com/iancoleman/strcase"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

type WeaviateModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	tpl *template.Template
}

func NewWeaviateModule() *WeaviateModule { return &WeaviateModule{ModuleBase: &pgs.ModuleBase{}} }

func (p *WeaviateModule) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	p.ctx = pgsgo.InitContext(c.Parameters())

	tpl := template.New("weaviate").Funcs(map[string]interface{}{
		"structName":       p.getStructName,
		"structFieldName":  p.getStructFieldName,
		"structFieldType":  p.getStructFieldType,
		"className":        p.getClassName,
		"propertyName":     p.getPropertyName,
		"propertyDataType": p.getPropertyDataType,
	})

	p.tpl = template.Must(tpl.Parse(WeaviateTemplate))
}

func (p *WeaviateModule) getClassName(m pgs.Message) string {
	return m.Name().String()
}

func (p *WeaviateModule) getStructName(m pgs.Message) string {
	return p.getClassName(m)
}

func (p *WeaviateModule) getCrossReferenceType(field pgs.Field) string {
	embed := field.Type().Embed()
	if embed != nil {
		return embed.Name().String()
	}
	return field.Type().Element().Embed().Name().String()
}

func (p *WeaviateModule) getStructFieldStructType(field pgs.Field) string {
	embed := field.Type().Embed()
	if embed != nil {
		return embed.Name().String()
	}
	return field.Type().Element().Embed().Name().String()
}

func (p *WeaviateModule) getNonCrossReferenceType(field pgs.Field) string {
	return weaviateTypeMap[field.Type().ProtoType()]
}

func (p *WeaviateModule) getNonStructFieldType(field pgs.Field) string {
	return goTypeMap[field.Type().ProtoType()]
}

func (p *WeaviateModule) getStructFieldType(field pgs.Field) (datatype string) {
	if isStructType(field) {
		datatype = p.getStructFieldStructType(field)
	} else {
		datatype = p.getNonStructFieldType(field)
	}
	if isList(field) {
		datatype = fmt.Sprintf("[]%s", datatype)
	}
	if isPointer(field) {
		datatype = fmt.Sprintf("*%s", datatype)
	}
	return
}

func (p *WeaviateModule) getPropertyDataType(field pgs.Field) (dataType string) {
	if isCrossReference(field) {
		dataType = p.getCrossReferenceType(field)
		return
	}
	dataType = p.getNonCrossReferenceType(field)
	if isList(field) {
		dataType = fmt.Sprintf("%s[]", dataType)
	}
	return
}

func isCrossReference(field pgs.Field) bool {
	return field.Type().ProtoType() == pgs.MessageT
}

func isStructType(field pgs.Field) bool {
	return field.Type().ProtoType() == pgs.MessageT
}

func isPointer(field pgs.Field) bool {
	label := field.Descriptor().Proto3Optional
	return label != nil && *label
}

func isList(field pgs.Field) bool {
	return field.Type().IsRepeated()
}

func (p *WeaviateModule) getPropertyName(field pgs.Field) string {
	return field.Name().String()
}

func (p *WeaviateModule) getStructFieldName(field pgs.Field) string {
	return strcase.ToCamel(p.getPropertyName(field))
}

func (p *WeaviateModule) Name() string { return "weaviate" }

func (p *WeaviateModule) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	for _, t := range targets {
		p.generate(t)
	}

	return p.Artifacts()
}

func (p *WeaviateModule) generate(f pgs.File) {
	if len(f.Messages()) == 0 {
		return
	}

	name := p.ctx.OutputPath(f).SetExt(".weaviate.go")
	p.AddGeneratorTemplateFile(name.String(), p.tpl, f)
}

func (p *WeaviateModule) marshaler(m pgs.Message) pgs.Name {
	return p.ctx.Name(m) + "JSONMarshaler"
}

func (p *WeaviateModule) unmarshaler(m pgs.Message) pgs.Name {
	return p.ctx.Name(m) + "JSONUnmarshaler"
}

const WeaviateTemplate = `
import (
  "github.com/weaviate/weaviate/entities/models"
)
{{ range .AllMessages }}
type {{ structName . }} struct {
	{{ range .Fields }}
	{{- structFieldName . }} {{ structFieldType . }}
	{{ end -}}
}

func ({{ structName . }}) WeaviateClassName() string {
	return "{{ className . }}"
}

func ({{ structName . }}) WeaviateModelDefinition() models.Model {
	return models.Class{
  		Class:       "{{ className . }}",
  		Properties: []*models.Property{
  		{{- range .Fields }}
			{
			  Name:        "{{ propertyName . }}",
			  DataType:    []string{"{{ propertyDataType . }}"},
			},
    	{{- end }}
  	}
} 
{{ end }}
`

var weaviateTypeMap = map[pgs.ProtoType]string{
	pgs.StringT:  "text",
	pgs.EnumT:    "text",
	pgs.BoolT:    "boolean",
	pgs.BytesT:   "blob",
	pgs.DoubleT:  "number",
	pgs.FloatT:   "number",
	pgs.Int64T:   "int",
	pgs.UInt64T:  "int",
	pgs.Int32T:   "int",
	pgs.Fixed64T: "int",
	pgs.Fixed32T: "int",
	pgs.UInt32T:  "int",
	pgs.SFixed32: "int",
	pgs.SFixed64: "int",
	pgs.SInt32:   "int",
	pgs.SInt64:   "int",
}

var goTypeMap = map[pgs.ProtoType]string{
	pgs.StringT:  "string",
	pgs.EnumT:    "int",
	pgs.BoolT:    "bool",
	pgs.BytesT:   "[]byte",
	pgs.DoubleT:  "float64",
	pgs.FloatT:   "float64",
	pgs.Int64T:   "int64",
	pgs.UInt64T:  "uint64",
	pgs.Int32T:   "int32",
	pgs.Fixed64T: "int64",
	pgs.Fixed32T: "int32",
	pgs.UInt32T:  "uint32",
	pgs.SFixed32: "int32",
	pgs.SFixed64: "int64",
	pgs.SInt32:   "int32",
	pgs.SInt64:   "int64",
}
