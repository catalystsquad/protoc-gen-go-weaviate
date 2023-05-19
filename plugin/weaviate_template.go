package plugin

const WeaviateTemplate = `
package {{ package }}

import (
	"context"
	"fmt"
	"github.com/catalystsquad/app-utils-go/errorutils"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/data"
	"github.com/weaviate/weaviate/entities/models"
	"strconv"
)

{{ range .messages }}
type {{ structName . }} struct {
	{{- range .Fields }}
	{{ fieldComments . -}}
    {{ structField . }} {{ jsonTag . -}}
	{{ end }}
}

func (s {{ structName . }}) ToProto() *{{ protoStructName . }} {
    theProto := &{{ protoStructName . }}{}
	{{- range .Fields }}
	{{- if and (fieldIsMessage .) (fieldIsRepeated .) }}
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
    return theProto
}

func (s *{{ protoStructName . }}) ToWeaviateModel() {{ structName . }} {
    model := {{ structName . }}{}
	{{- range .Fields }}
	{{- if and (fieldIsMessage .) (fieldIsRepeated .) }}
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
        {{- if ne (propertyName .) "id" }}
        "{{ jsonFieldName . }}": {{ dataField . }},
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

func (s {{ structName . }}) EnsureClass(client *weaviate.Client) {
	ensureClass(client, s.WeaviateClassSchema())
}
{{ end }}

func ensureClass(client *weaviate.Client, class models.Class) {
	var exists bool
	if exists = classExists(client, class.Class); exists {
		updateClass(client, class)
	} else {
		createClass(client, class)
	}
}

func updateClass(client *weaviate.Client, class models.Class) {
	var fetchedClass *models.Class
	if fetchedClass = getClass(client, class.Class); fetchedClass == nil {
		return
	}
	for _, property := range class.Properties {
		// continue on error, weaviate doesn't support updating property data types so we don't try to do that on startup
		// because it requires reindexing and is non trivial
		if containsProperty(fetchedClass.Properties, property) {
			continue
		}
		createProperty(client, class.Class, property)
	}
}

func createProperty(client *weaviate.Client, className string, property *models.Property) {
	err := client.Schema().PropertyCreator().WithClassName(className).WithProperty(property).Do(context.Background())
	errorutils.LogOnErr(logrus.WithFields(logrus.Fields{"class_name": className, "property_name": property.Name, "property_data_type": property.DataType}), "error creating property", err)
	return
}

func getClass(client *weaviate.Client, name string) (class *models.Class) {
	var err error
	class, err = client.Schema().ClassGetter().WithClassName(name).Do(context.Background())
	errorutils.LogOnErr(logrus.WithField("class_name", name), "error getting class", err)
	return
}

func createClass(client *weaviate.Client, class models.Class) {
	// all classes use contextionary
	class.Vectorizer = "text2vec-contextionary"
	err := client.Schema().ClassCreator().WithClass(&class).Do(context.Background())
	errorutils.LogOnErr(logrus.WithField("class_name", class.Class), "error creating class", err)
}

func classExists(client *weaviate.Client, name string) (exists bool) {
	var err error
	exists, err = client.Schema().ClassExistenceChecker().WithClassName(name).Do(context.Background())
	errorutils.LogOnErr(logrus.WithField("class_name", name), "error checking class existence", err)
	return
}

func containsProperty(source []*models.Property, property *models.Property) bool {
	// todo maybe: use a map/set to avoid repeated loops
	return lo.ContainsBy(source, func(item *models.Property) bool {
		return item.Name == property.Name
	})
}
`
