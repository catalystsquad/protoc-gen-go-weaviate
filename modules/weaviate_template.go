package modules

const WeaviateTemplate2 = `
package {{ package . }}

import (
	"context"
	"fmt"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/data"
	"github.com/weaviate/weaviate/entities/models"
)

{{ range .AllMessages }}
type {{ structName . }} struct {
	{{- range .Fields }}
    {{ structField . -}}
	{{ end }}
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
			  Name:        "{{ propertyName . }}",
			  DataType:    []string{"{{ propertyDataType . }}"},
			},
            {{- end -}}
    	{{ end }}
	}
}

func (s {{ structName . }}) Data() map[string]interface{} {
	data := map[string]interface{}{
		{{- range .Fields }}
    	{{ dataField . }},
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
      reference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
	  data["{{ propertyName . }}"] = append(data["{{ propertyName . }}"].([]map[string]string), reference)
	}
    {{- end -}}
    {{- end -}}
	{{- end }}
	return data
}

func (s {{ structName . }}) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
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
{{ end }}
`
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

func (s {{ structName . }}) WeaviateClassName() string {
	return "{{ className . }}"
}

func (s {{ structName . }}) WeaviateModelDefinition() models.Model {
	return models.Class{
  		Class:       "{{ className . }}",
  		Data: []*models.Property{
  		{{- range .Fields }}
			{
			  Name:        "{{ propertyName . }}",
			  DataType:    []string{"{{ propertyDataType . }}"},
			},
    	{{- end }}
  	}
}

func (s {{ structName . }}) ToWeaviateObject() models.Object {
	return models.Object{
  		Class:       "{{ className . }}",
        Id: s.Id,
  		Data: s.getWeaviateProperties()
}
{{ end }}
`
