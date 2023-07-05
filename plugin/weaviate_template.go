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
{{ if shouldGenerateMessage . }}
type {{ structName . }} struct {
	{{- range .Fields }}
	{{ if includeField . }}
	{{ fieldComments . -}}
    {{ structField . }} {{ jsonTag . -}}
	{{ end }}
	{{ end }}
}

func (s {{ structName . }}) ToProto() (theProto *{{ protoStructName . }}, err error) {
    theProto = &{{ protoStructName . }}{}
	{{- range .Fields }}
	{{ if includeField . }}
	{{- if isTimestamp . }}
	if s.{{ .GoName }} != nil {
		theProto.{{ .GoName }} = timestamppb.New(*s.{{ .GoName }})
	}
	{{- else if isStructPb . }}
	if s.{{ .GoName }} != "" {
		if err = json.Unmarshal([]byte(s.{{ .GoName }}), &theProto.{{ .GoName }}); err != nil {
			return
		}
	}
	{{- else if and (fieldIsMessage .) (fieldIsRepeated .) }}
    for _, protoField := range s.{{ structFieldName . }} {
		msg, err := protoField.ToProto()
		if err != nil {
			return nil, err
		}
		if theProto.{{ structFieldName . }} == nil {
			theProto.{{ structFieldName . }} = []*{{ protoStructTypeFromField . }}{msg}
		} else {
			theProto.{{ structFieldName . }} = append(theProto.{{ structFieldName . }}, msg)
		}
	}
    {{- else if and (fieldIsMessage .) }}
	if theProto.{{ structFieldName . }}, err = s.{{ structFieldName . }}.ToProto(); err != nil {
		return
	}
    {{- else }}
    theProto.{{ structFieldName . }} = s.{{ structFieldName . }}
    {{ end }}
	{{ end }}
	{{ end }}
    return
}

func (s *{{ protoStructName . }}) ToWeaviateModel() (model {{ structName . }}, err error) {
    model = {{ structName . }}{}
	{{- range .Fields }}
	{{ if includeField . }}
	{{- if isTimestamp . }}
	if s.{{ .GoName }} != nil {
		model.{{ .GoName }} = lo.ToPtr(s.{{ .GoName }}.AsTime())
	}
	{{- else if isStructPb . }}
	{{ structFieldName . }}Bytes, err := s.{{ structFieldName . }}.MarshalJSON()
	if err != nil {
		return model, err
	}
	if {{ structFieldName . }}Bytes != nil {
		model.{{ structFieldName . }} = string({{ structFieldName . }}Bytes)
	}
	{{- else if and (fieldIsMessage .) (fieldIsRepeated .) (not (isStructPb .)) }}
    for _, protoField := range s.{{ structFieldName . }} {
		msg, err := protoField.ToWeaviateModel()
		if err != nil {
			return model, err
		}
		if model.{{ structFieldName . }} == nil {
			model.{{ structFieldName . }} = {{ structFieldType . }}{msg}
		} else {
			model.{{ structFieldName . }} = append(model.{{ structFieldName . }}, msg)
		}
	}
    {{- else if and (fieldIsMessage .) (not (fieldIsOptional .)) (not (isStructPb .)) }}
	if s.{{ structFieldName . }} != nil {
		if model.{{ structFieldName . }}, err = s.{{ structFieldName . }}.ToWeaviateModel(); err != nil {
			return
		}
	}
	{{- else if and (fieldIsMessage .) (fieldIsOptional .) (not (isStructPb .)) }}
	if s.{{ structFieldName . }} != nil {
		model{{ structFieldName . }}, err := s.{{ structFieldName . }}.ToWeaviateModel()
		if err != nil {
			return model, err
		}
		model.{{ structFieldName . }} = lo.ToPtr(model{{ structFieldName . }})
	}
    {{- else }}
    model.{{ structFieldName . }} = s.{{ structFieldName . }}
    {{ end }}
	{{ end }}
	{{ end }}
    return
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
	data := map[string]interface{}{}
	{{- range .Fields }}
	{{ if includeField . }}
	{{- if ne (propertyName .) "id" }}
	{{- if fieldIsOptional . }}
	if s.{{ .GoName }} != nil {
	data["{{ jsonFieldName . }}"] = {{ dataField . }}
	{{- else }}
	data["{{ jsonFieldName . }}"] = {{ dataField . }}
	{{ end }}
	{{- if fieldIsOptional . }}
	}
	{{ end }}
	{{- end }}
	{{- end }}
	{{- end }}

	data = s.addCrossReferenceData(data)
	
	return data
}

func (s {{ structName . }}) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
    {{- range .Fields }}
    {{- if fieldIsCrossReference . -}}
    {{- if fieldIsRepeated . }}
	for _, crossReference := range s.{{ structFieldName . }} {
      {{- if crossReferenceIdFieldIsOptional . }}
      id := lo.FromPtr(crossReference.Id)
	  {{- else }}
      id := crossReference.Id
      {{ end }}
      {{ structFieldName . }}Reference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", id)}
	  data["{{ jsonFieldName . }}"] = append(data["{{ jsonFieldName . }}"].([]map[string]string), {{ structFieldName . }}Reference)
	}
    {{- else }}
	{{ if fieldIsOptional . -}}
	if s.{{ structFieldName . }} != nil {
    {{- end -}}
	{{- if crossReferenceIdFieldIsOptional . }}
	if lo.FromPtr(s.{{ structFieldName . }}.Id) != "" {
	{{- else }}
	if s.{{ structFieldName . }}.Id != "" {
	{{- end }}
	{{- if crossReferenceIdFieldIsOptional . }}
    id := lo.FromPtr(s.{{ structFieldName . }}.Id)
    {{- else }}
    id := s.{{ structFieldName . }}.Id
    {{ end }}
	{{ structFieldName . }}Reference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", id)}
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
		{{- if idFieldIsOptional . }}
		WithID(lo.FromPtr(s.Id)).
		{{- else }}
		WithID(s.Id).
		{{- end }}
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s {{ structName . }}) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		{{- if idFieldIsOptional . }}
		WithID(lo.FromPtr(s.Id)).
		{{- else }}
		WithID(s.Id).
		{{- end }}
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s {{ structName . }}) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		{{- if idFieldIsOptional . }}
		WithID(lo.FromPtr(s.Id)).
		{{- else }}
		WithID(s.Id).
		{{- end }}
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
