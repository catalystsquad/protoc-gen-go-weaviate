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

var summaryRegex = regexp.MustCompile("[^a-zA-Z0-9 -_]+")

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
	{{- if fieldIsOptional . }}
	if s.{{ .GoName }} != nil {
	{{- end }}
	{{- if isTimestamp . }}
	{{- if fieldIsOptional . }}
		theProto.{{ .GoName }} = timestamppb.New(*s.{{ .GoName }})
	{{- else }}
		theProto.{{ .GoName }} = timestamppb.New(s.{{ .GoName }})
	{{- end }}
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
	{{- if fieldIsOptional . }}
	}
	{{- end }}
	{{ end }}
	{{ end }}
    return
}

func (s *{{ protoStructName . }}) ToWeaviateModel() (model *{{ structName . }}, err error) {
    model = &{{ structName . }}{}
	{{- range .Fields }}
	{{ if includeField . }}
	{{- if isTimestamp . }}
	if s.{{ .GoName }} != nil {
		{{- if fieldIsOptional . }}
		model.{{ .GoName }} = lo.ToPtr(s.{{ .GoName }}.AsTime())
		{{- else }}
		model.{{ .GoName }} = s.{{ .GoName }}.AsTime()
		{{- end }}
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
		model.{{ structFieldName . }} = model{{ structFieldName . }}
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

func (s {{ structName . }}) FullWeaviateClassSchema() models.Class {
    class := models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.AllWeaviateClassSchemaProperties(),
	}
	{{ if classModuleConfig . }}
	var {{ structName . }}ModuleConfig map[string]interface{}
	{{ structName . }}ModuleConfigBytes := []byte(` + "`" + `{{ classModuleConfig . }}` + "`" + `)
	{{ structName . }}Err := json.Unmarshal({{ structName . }}ModuleConfigBytes, &{{ structName . }}ModuleConfig)
	if {{ structName . }}Err != nil {
	  panic({{ structName . }}Err)
	}
	class.ModuleConfig = {{ structName . }}ModuleConfig
	{{ end }}
	return class
}

func (s {{ structName . }}) CrossReferenceWeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaCrossReferenceProperties(),
	}
}

func (s {{ structName . }}) NonCrossReferenceWeaviateClassSchema() models.Class {
	class := models.Class{
		Class:      s.WeaviateClassName(),
        {{- if vectorizer . }}
		Vectorizer: "{{ vectorizer . }}",
        {{- end }}
		Properties: s.WeaviateClassSchemaNonCrossReferenceProperties(),
	}
	
	{{ if classModuleConfig . }}
	var {{ structName . }}ModuleConfig map[string]interface{}
	{{ structName . }}ModuleConfigBytes := []byte(` + "`" + `{{ classModuleConfig . }}` + "`" + `)
	{{ structName . }}Err := json.Unmarshal({{ structName . }}ModuleConfigBytes, &{{ structName . }}ModuleConfig)
	if {{ structName . }}Err != nil {
	  panic({{ structName . }}Err)
	}
	class.ModuleConfig = {{ structName . }}ModuleConfig
	{{ end }}

	return class
}

func (s {{ structName . }}) WeaviateClassSchemaNonCrossReferenceProperties() []*models.Property {
	properties := []*models.Property{}
  		{{ range .Fields -}}
		    {{ if and (includeField . ) (ne (propertyName .) "id") (ne (fieldIsCrossReference .) true) }}
			{{ structFieldName . }}Property := &models.Property{
			  Name:        "{{ jsonFieldName . }}",
			  DataType:    []string{"{{ propertyDataType . }}"},
              {{- if and (eq (propertyDataType .) "text") (tokenization .) }}
              Tokenization: "{{ tokenization . }}",
              {{- end }}
			}
            {{ if moduleConfig . }}
            var {{ structFieldName . }}ModuleConfig map[string]interface{}
            {{ structFieldName . }}ModuleConfigBytes := []byte(` + "`" + `{{ moduleConfig . }}` + "`" + `)
            {{ structFieldName . }}Err := json.Unmarshal({{ structFieldName . }}ModuleConfigBytes, &{{ structFieldName . }}ModuleConfig)
            if {{ structFieldName . }}Err != nil {
              panic({{ structFieldName . }}Err)
            }
            {{ structFieldName . }}Property.ModuleConfig = {{ structFieldName . }}ModuleConfig
            {{ end }}
            properties = append(properties, {{ structFieldName . }}Property)
			{{ if and  (ne (propertyDataType .) "text") (ne (propertyDataType .) "blob") }}
			{{ structFieldName . }}WTextProperty := &models.Property{
			  Name:        "{{ jsonFieldName . }}Text",
			  DataType:    []string{"text"},
			}
			{{ if moduleConfig . }}
            var {{ structFieldName . }}TextModuleConfig map[string]interface{}
            {{ structFieldName . }}TextModuleConfigBytes := []byte(` + "`" + `{{ moduleConfig . }}` + "`" + `)
            {{ structFieldName . }}TextErr := json.Unmarshal({{ structFieldName . }}TextModuleConfigBytes, &{{ structFieldName . }}TextModuleConfig)
            if {{ structFieldName . }}TextErr != nil {
              panic({{ structFieldName . }}TextErr)
            }
            {{ structFieldName . }}WTextProperty.ModuleConfig = {{ structFieldName . }}TextModuleConfig
            {{ end }}
            properties = append(properties, {{ structFieldName . }}WTextProperty)
			{{- end -}}
            {{- end -}}
    	{{ end }}
        {{- if summaryEnabled . }}
        summaryProperty := &models.Property{
          Name: "_summary",
          DataType: []string{"text"},
          Tokenization: "field",
        }
        {{- if summaryModuleConfig . }}
        var summaryModuleConfig map[string]interface{}
	    summaryModuleConfigBytes := []byte(` + "`" + `{{ summaryModuleConfig . }}` + "`" + `)
	    summaryModuleConfigErr := json.Unmarshal(summaryModuleConfigBytes, &summaryModuleConfig)
	    if summaryModuleConfigErr != nil {
	      panic(summaryModuleConfigErr)
	    }
	    summaryProperty.ModuleConfig = summaryModuleConfig
        {{ end }}
        properties = append(properties, summaryProperty)
        {{- end }}
	return properties
}

func (s {{ structName . }}) WeaviateClassSchemaCrossReferenceProperties() []*models.Property {
	properties := []*models.Property{}
  		{{- range .Fields }}
		    {{ if and (includeField . ) (ne (propertyName .) "id") (fieldIsCrossReference .) -}}
			properties = append(properties, &models.Property{
			  Name:        "{{ jsonFieldName . }}",
			  DataType:    []string{"{{ propertyDataType . }}"},
			})
            {{- end -}}
    	{{ end }}
	return properties
}

func (s {{ structName . }}) AllWeaviateClassSchemaProperties() []*models.Property {
	return append(s.WeaviateClassSchemaNonCrossReferenceProperties(), s.WeaviateClassSchemaCrossReferenceProperties()...)
}

func (s {{ structName . }}) Data() (map[string]interface{}, error) {
	data := map[string]interface{}{}
	{{- range .Fields }}
	{{ if includeField . }}
	{{- if ne (propertyName .) "id" }}
    {{- if and (ne (propertyDataType .) "text") (ne (propertyDataType .) "blob") (eq (fieldIsCrossReference .) false) }}
      {{- if fieldIsOptional . }}
	    data["{{ jsonFieldName . }}Text"] = fmt.Sprintf("%v", lo.FromPtr(s.{{ structFieldName . }}))
	  {{- else }}
        data["{{ jsonFieldName . }}Text"] = fmt.Sprintf("%v", s.{{ structFieldName . }})
      {{- end }}
	{{- end -}}
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
	{{- if summaryEnabled . }}
    summary, err := s.SummaryData()
    if err != nil {
      return data, err
    }
    data["_summary"] = summary
    {{- end }}
	data = s.addCrossReferenceData(data)
	
	return data, nil
}

func (s {{ structName . }}) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
    {{- range .Fields }}
    {{- if and (includeField . ) (fieldIsCrossReference .) }}
    {{- if fieldIsRepeated . }}
	for _, crossReference := range s.{{ structFieldName . }} {
      if crossReference != nil {
        {{- if crossReferenceIdFieldIsOptional . }}
        id := lo.FromPtr(crossReference.Id)
	    {{- else }}
        id := crossReference.Id
        {{ end }}
        {{ structFieldName . }}Reference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", id)}
	    data["{{ jsonFieldName . }}"] = append(data["{{ jsonFieldName . }}"].([]map[string]string), {{ structFieldName . }}Reference)
      }
	}
    {{- else }}
	if s.{{ structFieldName . }} != nil {
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
	}
    {{- end -}}
    {{- end -}}
	{{- end }}
	return data
}

func (s {{ structName . }}) exists(ctx context.Context, client *weaviate.Client) (bool, error) {
	return client.Data().Checker().WithID(lo.FromPtr(s.Id)).WithClassName(s.WeaviateClassName()).Do(ctx)
}

func (s {{ structName . }}) Upsert(ctx context.Context, client *weaviate.Client, consistencyLevel string) (data *data.ObjectWrapper, err error) {
	data, err = s.Create(ctx, client, consistencyLevel)
	if err != nil && strings.Contains(err.Error(), "already exists") {
		err = s.Update(ctx, client, consistencyLevel)
	}
	return
}

func (s {{ structName . }}) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (data *data.ObjectWrapper, err error) {
	{{- range .Fields }}
	  {{- if and (includeField . ) (fieldIsCrossReference .) }}
        {{- if fieldIsRepeated . }}
          for _, crossReference := range s.{{ structFieldName . }} {
			if crossReference != nil {
			  _, err = crossReference.Upsert(ctx, client, consistencyLevel)
              if err != nil {
                return
              }
			}
	      }
        {{- else }}
        if s.{{ structFieldName . }} != nil {
		  _, err = s.{{ structFieldName . }}.Upsert(ctx, client, consistencyLevel)
		  if err != nil {
		    return
		  }
        }
		{{- end }}
	  {{- end }}
    {{- end }}
    var dataMap map[string]interface{}
    if dataMap, err = s.Data(); err != nil {
      return
    }
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(dataMap).
		{{- if idFieldIsOptional . }}
		WithID(lo.FromPtr(s.Id)).
		{{- else }}
		WithID(s.Id).
		{{- end }}
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s {{ structName . }}) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) (err error) {
	{{- range .Fields }}
	  {{- if and (includeField . ) (fieldIsCrossReference .) }}
        {{- if fieldIsRepeated . }}
          for _, crossReference := range s.{{ structFieldName . }} {
            if crossReference != nil {
              _, err = crossReference.Upsert(ctx, client, consistencyLevel)
              if err != nil {
                return
              }
            }
	      }
        {{- else }}
        if s.{{ structFieldName . }} != nil {
		  _, err = s.{{ structFieldName . }}.Upsert(ctx, client, consistencyLevel)
		  if err != nil {
		    return
		  }
        }
		{{- end }}
	  {{- end }}
    {{- end }}
    var dataMap map[string]interface{}
    if dataMap, err = s.Data(); err != nil {
      return
    }
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		{{- if idFieldIsOptional . }}
		WithID(lo.FromPtr(s.Id)).
		{{- else }}
		WithID(s.Id).
		{{- end }}
		WithProperties(dataMap).
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

func (s {{ structName . }}) EnsureFullClass(client *weaviate.Client, continueOnError bool) (err error) {
	if err = s.EnsureClassWithoutCrossReferences(client, continueOnError); err != nil {
		return
	}
	return s.EnsureClassWithCrossReferences(client, continueOnError)
}

func (s {{ structName . }}) EnsureClassWithoutCrossReferences(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.NonCrossReferenceWeaviateClassSchema(), continueOnError)
}

func (s {{ structName . }}) EnsureClassWithCrossReferences(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.CrossReferenceWeaviateClassSchema(), continueOnError)
}

func (s {{ structName . }}) SummaryData() (string, error) {
	return getStringValue(s)
}
{{ end }}
{{ end }}

func EnsureClasses(client *weaviate.Client, continueOnError bool) (err error) {
	// create classes without cross references first so there are no errors about missing classes
	{{- range .messages }}
	{{- if shouldGenerateMessage . }}
	err = {{ structName . }}{}.EnsureClassWithoutCrossReferences(client, continueOnError)
	if !continueOnError && err != nil {
		return
	}
	{{- end }}
	{{- end }}
	// update classes including cross references
	{{- range .messages }}
	{{- if shouldGenerateMessage . }}
	err = {{ structName . }}{}.EnsureClassWithCrossReferences(client, continueOnError)
	if !continueOnError && err != nil {
		return
	}
	{{- end }}
	{{- end }}
	return
}

func ensureClass(client *weaviate.Client, class models.Class, continueOnError bool) (err error) {
	var exists bool
	exists, err = classExists(client, class.Class)
	if !continueOnError && err != nil {
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
	if fetchedClass == nil || (!continueOnError && err != nil) {
		return
	}
	for _, property := range class.Properties {
		// continue on error, weaviate doesn't support updating property data types so we don't try to do that on startup
		// because it requires reindexing and is non trivial
		if containsProperty(fetchedClass.Properties, property) {
			continue
		}
		err = createProperty(client, class.Class, property)
		if !continueOnError && err != nil {
				return
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

func getStringValue(x interface{}) (value string, err error) {
	var jsonBytes []byte
	if jsonBytes, err = json.Marshal(x); err != nil {
		return
	}
	summaryString := summaryRegex.ReplaceAllString(string(jsonBytes), " ")
	value = strings.ToLower(summaryString)
	return
}
`
