package plugin

const WeaviateTemplate = `
package {{ package }}

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/data"
	"github.com/weaviate/weaviate/entities/models"
	"strconv"
)

{{ range .messages }}
{{ if shouldGenerate . }}
type {{ structName . }} struct {
	{{- range .Fields }}
	{{ if includeField . }}
	{{ fieldComments . -}}
    {{ structField . }} {{ jsonTag . -}}
	{{ end }}
	{{ end }}
}

func (s {{ structName . }}) ToProto() *{{ protoStructName . }} {
    theProto := &{{ protoStructName . }}{}
	{{- range .Fields }}
	{{ if includeField . }}
	{{- if isTimestamp . }}
	if s.{{ .GoName }} != nil {
		theProto.{{ .GoName }} = timestamppb.New(*s.{{ .GoName }})
	}
	{{- else if and (fieldIsMessage .) (fieldIsRepeated .) }}
    for _, protoField := range s.{{ structFieldName . }} {
		msg := protoField.ToProto()
		if theProto.{{ structFieldName . }} == nil {
			theProto.{{ structFieldName . }} = []*{{ protoStructTypeFromField . }}{msg}
		} else {
			theProto.{{ structFieldName . }} = append(theProto.{{ structFieldName . }}, msg)
		}
	}
    {{- else if fieldIsMessage . }}
	theProto.{{ structFieldName . }} = s.{{ structFieldName . }}.ToProto()
    {{- else }}
    theProto.{{ structFieldName . }} = s.{{ structFieldName . }}
    {{ end }}
	{{ end }}
	{{ end }}
    return theProto
}

func (s *{{ protoStructName . }}) ToWeaviateModel() {{ structName . }} {
    model := {{ structName . }}{}
	{{- range .Fields }}
	{{ if includeField . }}
	{{- if isTimestamp . }}
	if s.{{ .GoName }} != nil {
		model.{{ .GoName }} = lo.ToPtr(s.{{ .GoName }}.AsTime())
	}
	{{- else if and (fieldIsMessage .) (fieldIsRepeated .) }}
    for _, protoField := range s.{{ structFieldName . }} {
		msg := protoField.ToWeaviateModel()
		if model.{{ structFieldName . }} == nil {
			model.{{ structFieldName . }} = {{ structFieldType . }}{msg}
		} else {
			model.{{ structFieldName . }} = append(model.{{ structFieldName . }}, msg)
		}
	}
    {{- else if and (fieldIsMessage .) (not (fieldIsOptional .)) }}
	if s.{{ structFieldName . }} != nil {
	model.{{ structFieldName . }} = s.{{ structFieldName . }}.ToWeaviateModel()
	}
	{{- else if and (fieldIsMessage .) (fieldIsOptional .) }}
	if s.{{ structFieldName . }} != nil {
		model.{{ structFieldName . }} = lo.ToPtr(s.{{ structFieldName . }}.ToWeaviateModel())
	}
    {{- else }}
    model.{{ structFieldName . }} = s.{{ structFieldName . }}
    {{ end }}
	{{ end }}
	{{ end }}
    return model
}

func (s {{ structName . }}) WeaviateClassName() string {
	return "{{ className . }}"
}

func (s {{ structName . }}) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s {{ structName . }}) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{
  		{{- range .Fields -}}
		    {{ if ne (propertyName .) "id" -}}
			{
			  Name:        "{{ jsonFieldName . }}",
			  DataType:    []string{"{{ propertyDataType . }}"},
			},
            {{- end -}}
    	{{ end }}
	}
}

func (s {{ structName . }}) Data() map[string]interface{} {
	data := map[string]interface{}{
		{{- range .Fields }}
		{{ if includeField . }}
        {{- if ne (propertyName .) "id" }}
        "{{ jsonFieldName . }}": {{ dataField . }},
		{{- end }}
		{{- end }}
		{{- end }}
	}

	data = s.addCrossReferenceData(data)
	
	return data
}

func (s {{ structName . }}) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
    {{- range .Fields }}
    {{- if fieldIsCrossReference . -}}
    {{- if fieldIsRepeated . }}
	for _, crossReference := range s.{{ structFieldName . }} {
      {{ structFieldName . }}Reference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
	  data["{{ jsonFieldName . }}"] = append(data["{{ jsonFieldName . }}"].([]map[string]string), {{ structFieldName . }}Reference)
	}
    {{- else }}
	{{ if fieldIsOptional . -}}
	if s.{{ structFieldName . }} != nil {
    {{- end -}}
    if s.{{ structFieldName . }}.Id != "" {
	{{ structFieldName . }}Reference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", s.{{ structFieldName . }}.Id)}
    data["{{ jsonFieldName . }}"] = append(data["{{ jsonFieldName . }}"].([]map[string]string), {{ structFieldName . }}Reference)
    }
	{{ if fieldIsOptional . -}}
	}
    {{- end -}}
    {{- end -}}
    {{- end -}}
	{{- end }}
	return data
}

func (s {{ structName . }}) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s {{ structName . }}) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s {{ structName . }}) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s {{ structName . }}) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}
{{ end }}
{{ end }}

func ensureClass(client *weaviate.Client, class models.Class, continueOnError bool) (err error) {
	var exists bool
	exists, err = classExists(client, class.Class)
	if err != nil {
		return
	}
	if exists {
		return updateClass(client, class, continueOnError)
	} else {
		return createClass(client, class)
	}
}

func updateClass(client *weaviate.Client, class models.Class, continueOnError bool) (err error) {
	var fetchedClass *models.Class
	fetchedClass, err = getClass(client, class.Class)
	if err != nil || fetchedClass == nil {
		return
	}
	for _, property := range class.Properties {
		// continue on error, weaviate doesn't support updating property data types so we don't try to do that on startup
		// because it requires reindexing and is non trivial
		if containsProperty(fetchedClass.Properties, property) {
			continue
		}
		err = createProperty(client, class.Class, property)
		if err != nil {
			if !continueOnError {
				return
			}
		}
	}
	return
}

func createProperty(client *weaviate.Client, className string, property *models.Property) (err error) {
	return client.Schema().PropertyCreator().WithClassName(className).WithProperty(property).Do(context.Background())
}

func getClass(client *weaviate.Client, name string) (class *models.Class, err error) {
	return client.Schema().ClassGetter().WithClassName(name).Do(context.Background())
}

func createClass(client *weaviate.Client, class models.Class) (err error) {
	return client.Schema().ClassCreator().WithClass(&class).Do(context.Background())
}

func classExists(client *weaviate.Client, name string) (exists bool, err error) {
	return client.Schema().ClassExistenceChecker().WithClassName(name).Do(context.Background())
}

func containsProperty(source []*models.Property, property *models.Property) bool {
	// todo maybe: use a map/set to avoid repeated loops
	return lo.ContainsBy(source, func(item *models.Property) bool {
		return item.Name == property.Name
	})
}
`
