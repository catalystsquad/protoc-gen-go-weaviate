package modules

import (
	"fmt"
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
		"className": p.getClassName,
		"dataType":  p.getPropertyDataType,
		"name":      p.getPropertyName,
	})

	p.tpl = template.Must(tpl.Parse(WeaviateTemplate))
}

func (p *WeaviateModule) getClassName(m pgs.Message) string {
	return m.Name().String()
}

func (p *WeaviateModule) getCrossReferenceType(field pgs.Field) string {
	embed := field.Type().Embed()
	if embed != nil {
		return embed.Name().String()
	}
	return field.Type().Element().Embed().Name().String()
}

func (p *WeaviateModule) getNonCrossReferenceType(field pgs.Field) string {
	return typeMap[field.Type().ProtoType()]
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

func isList(field pgs.Field) bool {
	return field.Type().IsRepeated()
}

func getFieldProtoType(field pgs.Field) pgs.ProtoType {
	return field.Type().ProtoType()
}
func (p *WeaviateModule) getPropertyName(field pgs.Field) string {
	return field.Name().String()
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
var {{ className . }}Class = models.Class{
  Class:       "{{ className . }}",
  Properties: []*models.Property{
  {{- range .Fields }}
    {
      DataType:    []string{"{{ dataType . }}"},
      Name:        "{{ name . }}",
    },
    {{- end }}
  },
}
{{ end }}
`

var typeMap = map[pgs.ProtoType]string{
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
