package example_example

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/data"
	"github.com/weaviate/weaviate/entities/models"
	"strconv"
)

type ThingWeaviateModel struct {

	// @gotags: fake:"{uuid}"
	Id string `json:"id"`

	// @gotags: fake:"{price:0.00,1000.00}"
	ADouble float64 `json:"aDouble"`

	// @gotags: fake:"{price:0.00,1000.00}"
	AFloat float32 `json:"aFloat"`

	// @gotags: fake:"{int32}"
	AnInt32 int32 `json:"anInt32"`

	// @gotags: fake:"{int64}"
	AnInt64 int64 `json:"anInt64,string"`

	// @gotags: fake:"{bool}"
	ABool bool `json:"aBool"`

	// @gotags: fake:"{hackerphrase}"
	AString string `json:"aString"`

	// @gotags: fake:"skip"
	ABytes []byte `json:"aBytes"`

	// @gotags: fake:"{hackerphrase}"
	RepeatedScalarField []string `json:"repeatedScalarField"`

	// @gotags: fake:"skip"
	OptionalScalarField *string `json:"optionalScalarField"`

	// @gotags: fake:"skip"
	AssociatedThing Thing2WeaviateModel `json:"associatedThing"`

	// @gotags: fake:"skip"
	OptionalAssociatedThing *Thing2WeaviateModel `json:"optionalAssociatedThing"`

	// @gotags: fake:"skip"
	RepeatedMessages []Thing2WeaviateModel `json:"repeatedMessages"`
}

func (s ThingWeaviateModel) ToProto() *Thing {
	theProto := &Thing{}

	theProto.Id = s.Id

	theProto.ADouble = s.ADouble

	theProto.AFloat = s.AFloat

	theProto.AnInt32 = s.AnInt32

	theProto.AnInt64 = s.AnInt64

	theProto.ABool = s.ABool

	theProto.AString = s.AString

	theProto.ABytes = s.ABytes

	theProto.RepeatedScalarField = s.RepeatedScalarField

	theProto.OptionalScalarField = s.OptionalScalarField

	theProto.AssociatedThing = s.AssociatedThing.ToProto()

	theProto.OptionalAssociatedThing = s.OptionalAssociatedThing.ToProto()

	for _, protoField := range s.RepeatedMessages {
		msg := protoField.ToProto()
		if theProto.RepeatedMessages == nil {
			theProto.RepeatedMessages = []*Thing2{msg}
		} else {
			theProto.RepeatedMessages = append(theProto.RepeatedMessages, msg)
		}
	}

	return theProto
}

func (s *Thing) ToWeaviateModel() ThingWeaviateModel {
	model := ThingWeaviateModel{}

	model.Id = s.Id

	model.ADouble = s.ADouble

	model.AFloat = s.AFloat

	model.AnInt32 = s.AnInt32

	model.AnInt64 = s.AnInt64

	model.ABool = s.ABool

	model.AString = s.AString

	model.ABytes = s.ABytes

	model.RepeatedScalarField = s.RepeatedScalarField

	model.OptionalScalarField = s.OptionalScalarField

	if s.AssociatedThing != nil {
		model.AssociatedThing = s.AssociatedThing.ToWeaviateModel()
	}

	if s.OptionalAssociatedThing != nil {
		model.OptionalAssociatedThing = lo.ToPtr(s.OptionalAssociatedThing.ToWeaviateModel())
	}

	for _, protoField := range s.RepeatedMessages {
		msg := protoField.ToWeaviateModel()
		if model.RepeatedMessages == nil {
			model.RepeatedMessages = []Thing2WeaviateModel{msg}
		} else {
			model.RepeatedMessages = append(model.RepeatedMessages, msg)
		}
	}

	return model
}

func (s ThingWeaviateModel) WeaviateClassName() string {
	return "Thing"
}

func (s ThingWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s ThingWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "aDouble",
		DataType: []string{"number"},
	}, {
		Name:     "aFloat",
		DataType: []string{"number"},
	}, {
		Name:     "anInt32",
		DataType: []string{"int"},
	}, {
		Name:     "anInt64",
		DataType: []string{"string"},
	}, {
		Name:     "aBool",
		DataType: []string{"boolean"},
	}, {
		Name:     "aString",
		DataType: []string{"text"},
	}, {
		Name:     "aBytes",
		DataType: []string{"blob"},
	}, {
		Name:     "repeatedScalarField",
		DataType: []string{"text[]"},
	}, {
		Name:     "optionalScalarField",
		DataType: []string{"text"},
	}, {
		Name:     "associatedThing",
		DataType: []string{"Thing2"},
	}, {
		Name:     "optionalAssociatedThing",
		DataType: []string{"Thing2"},
	}, {
		Name:     "repeatedMessages",
		DataType: []string{"Thing2"},
	}, {
		Name:     "anIgnoredField",
		DataType: []string{"text"},
	},
	}
}

func (s ThingWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"aDouble": s.ADouble,

		"aFloat": s.AFloat,

		"anInt32": s.AnInt32,

		"anInt64": strconv.FormatInt(s.AnInt64, 10),

		"aBool": s.ABool,

		"aString": s.AString,

		"aBytes": s.ABytes,

		"repeatedScalarField": s.RepeatedScalarField,

		"optionalScalarField": s.OptionalScalarField,

		"associatedThing": []map[string]string{},

		"optionalAssociatedThing": []map[string]string{},

		"repeatedMessages": []map[string]string{},
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s ThingWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	if s.AssociatedThing.Id != "" {
		AssociatedThingReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", s.AssociatedThing.Id)}
		data["associatedThing"] = append(data["associatedThing"].([]map[string]string), AssociatedThingReference)
	}

	if s.OptionalAssociatedThing != nil {
		if s.OptionalAssociatedThing.Id != "" {
			OptionalAssociatedThingReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", s.OptionalAssociatedThing.Id)}
			data["optionalAssociatedThing"] = append(data["optionalAssociatedThing"].([]map[string]string), OptionalAssociatedThingReference)
		}
	}
	for _, crossReference := range s.RepeatedMessages {
		RepeatedMessagesReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["repeatedMessages"] = append(data["repeatedMessages"].([]map[string]string), RepeatedMessagesReference)
	}
	return data
}

func (s ThingWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s ThingWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s ThingWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s ThingWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type Thing2WeaviateModel struct {

	// @gotags: fake:"{uuid}"
	Id string `json:"id"`

	// @gotags: fake:"{name}"
	Name string `json:"name"`
}

func (s Thing2WeaviateModel) ToProto() *Thing2 {
	theProto := &Thing2{}

	theProto.Id = s.Id

	theProto.Name = s.Name

	return theProto
}

func (s *Thing2) ToWeaviateModel() Thing2WeaviateModel {
	model := Thing2WeaviateModel{}

	model.Id = s.Id

	model.Name = s.Name

	return model
}

func (s Thing2WeaviateModel) WeaviateClassName() string {
	return "Thing2"
}

func (s Thing2WeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s Thing2WeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "name",
		DataType: []string{"text"},
	},
	}
}

func (s Thing2WeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"name": s.Name,
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s Thing2WeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	return data
}

func (s Thing2WeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s Thing2WeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s Thing2WeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s Thing2WeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
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
