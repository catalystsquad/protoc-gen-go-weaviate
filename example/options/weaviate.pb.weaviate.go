package weaviate

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/data"
	"github.com/weaviate/weaviate/entities/models"
	"strconv"
)

type WeaviateFileOptionsWeaviateModel struct {
}

func (s WeaviateFileOptionsWeaviateModel) ToProto() *WeaviateFileOptions {
	theProto := &WeaviateFileOptions{}
	return theProto
}

func (s *WeaviateFileOptions) ToWeaviateModel() WeaviateFileOptionsWeaviateModel {
	model := WeaviateFileOptionsWeaviateModel{}
	return model
}

func (s WeaviateFileOptionsWeaviateModel) WeaviateClassName() string {
	return "WeaviateFileOptions"
}

func (s WeaviateFileOptionsWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s WeaviateFileOptionsWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{}
}

func (s WeaviateFileOptionsWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{}

	data = s.addCrossReferenceData(data)

	return data
}

func (s WeaviateFileOptionsWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	return data
}

func (s WeaviateFileOptionsWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s WeaviateFileOptionsWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s WeaviateFileOptionsWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s WeaviateFileOptionsWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type WeaviateMessageOptionsWeaviateModel struct {
}

func (s WeaviateMessageOptionsWeaviateModel) ToProto() *WeaviateMessageOptions {
	theProto := &WeaviateMessageOptions{}
	return theProto
}

func (s *WeaviateMessageOptions) ToWeaviateModel() WeaviateMessageOptionsWeaviateModel {
	model := WeaviateMessageOptionsWeaviateModel{}
	return model
}

func (s WeaviateMessageOptionsWeaviateModel) WeaviateClassName() string {
	return "WeaviateMessageOptions"
}

func (s WeaviateMessageOptionsWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s WeaviateMessageOptionsWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{}
}

func (s WeaviateMessageOptionsWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{}

	data = s.addCrossReferenceData(data)

	return data
}

func (s WeaviateMessageOptionsWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	return data
}

func (s WeaviateMessageOptionsWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s WeaviateMessageOptionsWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s WeaviateMessageOptionsWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s WeaviateMessageOptionsWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type WeaviateFieldOptionsWeaviateModel struct {
	Ignore bool `json:"ignore"`
}

func (s WeaviateFieldOptionsWeaviateModel) ToProto() *WeaviateFieldOptions {
	theProto := &WeaviateFieldOptions{}

	theProto.Ignore = s.Ignore

	return theProto
}

func (s *WeaviateFieldOptions) ToWeaviateModel() WeaviateFieldOptionsWeaviateModel {
	model := WeaviateFieldOptionsWeaviateModel{}

	model.Ignore = s.Ignore

	return model
}

func (s WeaviateFieldOptionsWeaviateModel) WeaviateClassName() string {
	return "WeaviateFieldOptions"
}

func (s WeaviateFieldOptionsWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s WeaviateFieldOptionsWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "ignore",
		DataType: []string{"boolean"},
	},
	}
}

func (s WeaviateFieldOptionsWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"ignore": s.Ignore,
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s WeaviateFieldOptionsWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	return data
}

func (s WeaviateFieldOptionsWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s WeaviateFieldOptionsWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s WeaviateFieldOptionsWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s WeaviateFieldOptionsWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

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
