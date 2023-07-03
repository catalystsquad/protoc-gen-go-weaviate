package timestamppb

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/data"
	"github.com/weaviate/weaviate/entities/models"
	"strconv"
)

type TimestampWeaviateModel struct {

	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds int64 `json:"seconds,string"`

	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive.
	Nanos int32 `json:"nanos"`
}

func (s TimestampWeaviateModel) ToProto() *Timestamp {
	theProto := &Timestamp{}

	theProto.Seconds = s.Seconds

	theProto.Nanos = s.Nanos

	return theProto
}

func (s *Timestamp) ToWeaviateModel() TimestampWeaviateModel {
	model := TimestampWeaviateModel{}

	model.Seconds = s.Seconds

	model.Nanos = s.Nanos

	return model
}

func (s TimestampWeaviateModel) WeaviateClassName() string {
	return "Timestamp"
}

func (s TimestampWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s TimestampWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "seconds",
		DataType: []string{"string"},
	}, {
		Name:     "nanos",
		DataType: []string{"int"},
	},
	}
}

func (s TimestampWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"seconds": strconv.FormatInt(s.Seconds, 10),

		"nanos": s.Nanos,
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s TimestampWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	return data
}

func (s TimestampWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s TimestampWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s TimestampWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s TimestampWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
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
