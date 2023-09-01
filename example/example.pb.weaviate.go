package example_example

import (
	json "encoding/json"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	regexp "regexp"
	strings "strings"
	time "time"
)

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/data"
	"github.com/weaviate/weaviate/entities/models"
	"strconv"
)

var summaryRegex = regexp.MustCompile("[^-'\\\\.a-zA-Z0-9 ]+")

type ThingWeaviateModel struct {

	// @gotags: fake:"{uuid}"
	Id *string `json:"id" fake:"{uuid}"`

	// @gotags: fake:"{price:0.00,1000.00}"
	ADouble float64 `json:"aDouble" fake:"{price:0.00,1000.00}"`

	// @gotags: fake:"{price:0.00,1000.00}"
	AFloat float32 `json:"aFloat" fake:"{price:0.00,1000.00}"`

	// @gotags: fake:"{int32}"
	AnInt32 int32 `json:"anInt32" fake:"{int32}"`

	// @gotags: fake:"{int64}"
	AnInt64 int64 `json:"anInt64,string" fake:"{int64}"`

	// @gotags: fake:"{bool}"
	ABool bool `json:"aBool" fake:"{bool}"`

	// @gotags: fake:"{hackerphrase}"
	AString string `json:"aString" fake:"{hackerphrase}"`

	// @gotags: fake:"skip"
	ABytes []byte `json:"aBytes" fake:"skip"`

	// @gotags: fake:"{hackerphrase}"
	RepeatedScalarField []string `json:"repeatedScalarField" fake:"{hackerphrase}"`

	// @gotags: fake:"skip"
	OptionalScalarField *string `json:"optionalScalarField" fake:"skip"`

	// @gotags: fake:"skip"
	AssociatedThing *Thing2WeaviateModel `json:"associatedThing" fake:"skip"`

	// @gotags: fake:"skip"
	OptionalAssociatedThing *Thing2WeaviateModel `json:"optionalAssociatedThing" fake:"skip"`

	// @gotags: fake:"skip"
	RepeatedMessages []*Thing2WeaviateModel `json:"repeatedMessages" fake:"skip"`

	// @gotags: fake:"skip"
	ATimestamp time.Time `json:"aTimestamp" fake:"skip"`

	// @gotags: fake:"skip"
	AStructField string `json:"aStructField" fake:"skip"`

	// @gotags: fake:"{number:1,2}"
	AnEnum AnEnumType `json:"anEnum" fake:"{number:1,2}"`

	// @gotags: fake:"{number: 100,1000}"
	AnOptionalInt *int32 `json:"anOptionalInt" fake:"{number: 100,1000}"`

	// @gotags: fake:"skip"
	OptionalTimestamp *time.Time `json:"optionalTimestamp" fake:"skip"`
}

func (s ThingWeaviateModel) ToProto() (theProto *Thing, err error) {
	theProto = &Thing{}

	if s.Id != nil {
		theProto.Id = s.Id

	}

	theProto.ADouble = s.ADouble

	theProto.AFloat = s.AFloat

	theProto.AnInt32 = s.AnInt32

	theProto.AnInt64 = s.AnInt64

	theProto.ABool = s.ABool

	theProto.AString = s.AString

	theProto.ABytes = s.ABytes

	theProto.RepeatedScalarField = s.RepeatedScalarField

	if s.OptionalScalarField != nil {
		theProto.OptionalScalarField = s.OptionalScalarField

	}

	if theProto.AssociatedThing, err = s.AssociatedThing.ToProto(); err != nil {
		return
	}

	if s.OptionalAssociatedThing != nil {
		if theProto.OptionalAssociatedThing, err = s.OptionalAssociatedThing.ToProto(); err != nil {
			return
		}
	}

	for _, protoField := range s.RepeatedMessages {
		msg, err := protoField.ToProto()
		if err != nil {
			return nil, err
		}
		if theProto.RepeatedMessages == nil {
			theProto.RepeatedMessages = []*Thing2{msg}
		} else {
			theProto.RepeatedMessages = append(theProto.RepeatedMessages, msg)
		}
	}

	theProto.ATimestamp = timestamppb.New(s.ATimestamp)

	if s.AStructField != "" {
		if err = json.Unmarshal([]byte(s.AStructField), &theProto.AStructField); err != nil {
			return
		}
	}

	theProto.AnEnum = s.AnEnum

	if s.AnOptionalInt != nil {
		theProto.AnOptionalInt = s.AnOptionalInt

	}

	if s.OptionalTimestamp != nil {
		theProto.OptionalTimestamp = timestamppb.New(*s.OptionalTimestamp)
	}

	return
}

func (s *Thing) ToWeaviateModel() (model *ThingWeaviateModel, err error) {
	model = &ThingWeaviateModel{}

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
		if model.AssociatedThing, err = s.AssociatedThing.ToWeaviateModel(); err != nil {
			return
		}
	}

	if s.OptionalAssociatedThing != nil {
		modelOptionalAssociatedThing, err := s.OptionalAssociatedThing.ToWeaviateModel()
		if err != nil {
			return model, err
		}
		model.OptionalAssociatedThing = modelOptionalAssociatedThing
	}

	for _, protoField := range s.RepeatedMessages {
		msg, err := protoField.ToWeaviateModel()
		if err != nil {
			return model, err
		}
		if model.RepeatedMessages == nil {
			model.RepeatedMessages = []*Thing2WeaviateModel{msg}
		} else {
			model.RepeatedMessages = append(model.RepeatedMessages, msg)
		}
	}

	if s.ATimestamp != nil {
		model.ATimestamp = s.ATimestamp.AsTime()
	}

	AStructFieldBytes, err := s.AStructField.MarshalJSON()
	if err != nil {
		return model, err
	}
	if AStructFieldBytes != nil {
		model.AStructField = string(AStructFieldBytes)
	}

	model.AnEnum = s.AnEnum

	model.AnOptionalInt = s.AnOptionalInt

	if s.OptionalTimestamp != nil {
		model.OptionalTimestamp = lo.ToPtr(s.OptionalTimestamp.AsTime())
	}

	return
}

func (s ThingWeaviateModel) WeaviateClassName() string {
	return "Thing"
}

func (s ThingWeaviateModel) FullWeaviateClassSchema() models.Class {
	class := models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.AllWeaviateClassSchemaProperties(),
	}

	var ThingWeaviateModelModuleConfig map[string]interface{}
	ThingWeaviateModelModuleConfigBytes := []byte(`{"reranker-cohere": {}}`)
	ThingWeaviateModelErr := json.Unmarshal(ThingWeaviateModelModuleConfigBytes, &ThingWeaviateModelModuleConfig)
	if ThingWeaviateModelErr != nil {
		panic(ThingWeaviateModelErr)
	}
	class.ModuleConfig = ThingWeaviateModelModuleConfig

	return class
}

func (s ThingWeaviateModel) CrossReferenceWeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaCrossReferenceProperties(),
	}
}

func (s ThingWeaviateModel) NonCrossReferenceWeaviateClassSchema() models.Class {
	class := models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaNonCrossReferenceProperties(),
	}

	var ThingWeaviateModelModuleConfig map[string]interface{}
	ThingWeaviateModelModuleConfigBytes := []byte(`{"reranker-cohere": {}}`)
	ThingWeaviateModelErr := json.Unmarshal(ThingWeaviateModelModuleConfigBytes, &ThingWeaviateModelModuleConfig)
	if ThingWeaviateModelErr != nil {
		panic(ThingWeaviateModelErr)
	}
	class.ModuleConfig = ThingWeaviateModelModuleConfig

	return class
}

func (s ThingWeaviateModel) WeaviateClassSchemaNonCrossReferenceProperties() []*models.Property {
	properties := []*models.Property{}

	ADoubleProperty := &models.Property{
		Name:     "aDouble",
		DataType: []string{"number"},
	}

	properties = append(properties, ADoubleProperty)

	ADoubleWTextProperty := &models.Property{
		Name:     "aDoubleText",
		DataType: []string{"text"},
	}

	properties = append(properties, ADoubleWTextProperty)
	AFloatProperty := &models.Property{
		Name:     "aFloat",
		DataType: []string{"number"},
	}

	properties = append(properties, AFloatProperty)

	AFloatWTextProperty := &models.Property{
		Name:     "aFloatText",
		DataType: []string{"text"},
	}

	properties = append(properties, AFloatWTextProperty)
	AnInt32Property := &models.Property{
		Name:     "anInt32",
		DataType: []string{"int"},
	}

	var AnInt32ModuleConfig map[string]interface{}
	AnInt32ModuleConfigBytes := []byte(`{"text2vec-contextionary": {"vectorizePropertyName": true}}`)
	AnInt32Err := json.Unmarshal(AnInt32ModuleConfigBytes, &AnInt32ModuleConfig)
	if AnInt32Err != nil {
		panic(AnInt32Err)
	}
	AnInt32Property.ModuleConfig = AnInt32ModuleConfig

	properties = append(properties, AnInt32Property)

	AnInt32WTextProperty := &models.Property{
		Name:     "anInt32Text",
		DataType: []string{"text"},
	}

	var AnInt32TextModuleConfig map[string]interface{}
	AnInt32TextModuleConfigBytes := []byte(`{"text2vec-contextionary": {"vectorizePropertyName": true}}`)
	AnInt32TextErr := json.Unmarshal(AnInt32TextModuleConfigBytes, &AnInt32TextModuleConfig)
	if AnInt32TextErr != nil {
		panic(AnInt32TextErr)
	}
	AnInt32WTextProperty.ModuleConfig = AnInt32TextModuleConfig

	properties = append(properties, AnInt32WTextProperty)
	AnInt64Property := &models.Property{
		Name:     "anInt64",
		DataType: []string{"string"},
	}

	properties = append(properties, AnInt64Property)

	AnInt64WTextProperty := &models.Property{
		Name:     "anInt64Text",
		DataType: []string{"text"},
	}

	properties = append(properties, AnInt64WTextProperty)
	ABoolProperty := &models.Property{
		Name:     "aBool",
		DataType: []string{"boolean"},
	}

	properties = append(properties, ABoolProperty)

	ABoolWTextProperty := &models.Property{
		Name:     "aBoolText",
		DataType: []string{"text"},
	}

	properties = append(properties, ABoolWTextProperty)
	AStringProperty := &models.Property{
		Name:         "aString",
		DataType:     []string{"text"},
		Tokenization: "word",
	}

	properties = append(properties, AStringProperty)

	ABytesProperty := &models.Property{
		Name:     "aBytes",
		DataType: []string{"blob"},
	}

	properties = append(properties, ABytesProperty)

	RepeatedScalarFieldProperty := &models.Property{
		Name:     "repeatedScalarField",
		DataType: []string{"text[]"},
	}

	properties = append(properties, RepeatedScalarFieldProperty)

	RepeatedScalarFieldWTextProperty := &models.Property{
		Name:     "repeatedScalarFieldText",
		DataType: []string{"text"},
	}

	properties = append(properties, RepeatedScalarFieldWTextProperty)
	OptionalScalarFieldProperty := &models.Property{
		Name:     "optionalScalarField",
		DataType: []string{"text"},
	}

	properties = append(properties, OptionalScalarFieldProperty)

	ATimestampProperty := &models.Property{
		Name:     "aTimestamp",
		DataType: []string{"date"},
	}

	properties = append(properties, ATimestampProperty)

	ATimestampWTextProperty := &models.Property{
		Name:     "aTimestampText",
		DataType: []string{"text"},
	}

	properties = append(properties, ATimestampWTextProperty)
	AStructFieldProperty := &models.Property{
		Name:     "aStructField",
		DataType: []string{"text"},
	}

	properties = append(properties, AStructFieldProperty)

	AnEnumProperty := &models.Property{
		Name:     "anEnum",
		DataType: []string{"int"},
	}

	properties = append(properties, AnEnumProperty)

	AnEnumWTextProperty := &models.Property{
		Name:     "anEnumText",
		DataType: []string{"text"},
	}

	properties = append(properties, AnEnumWTextProperty)
	AnOptionalIntProperty := &models.Property{
		Name:     "anOptionalInt",
		DataType: []string{"int"},
	}

	properties = append(properties, AnOptionalIntProperty)

	AnOptionalIntWTextProperty := &models.Property{
		Name:     "anOptionalIntText",
		DataType: []string{"text"},
	}

	properties = append(properties, AnOptionalIntWTextProperty)
	OptionalTimestampProperty := &models.Property{
		Name:     "optionalTimestamp",
		DataType: []string{"date"},
	}

	properties = append(properties, OptionalTimestampProperty)

	OptionalTimestampWTextProperty := &models.Property{
		Name:     "optionalTimestampText",
		DataType: []string{"text"},
	}

	properties = append(properties, OptionalTimestampWTextProperty)
	summaryProperty := &models.Property{
		Name:         "_summary",
		DataType:     []string{"text"},
		Tokenization: "field",
	}
	var summaryModuleConfig map[string]interface{}
	summaryModuleConfigBytes := []byte(`{"text2vec-contextionary": {"skip": true}}`)
	summaryModuleConfigErr := json.Unmarshal(summaryModuleConfigBytes, &summaryModuleConfig)
	if summaryModuleConfigErr != nil {
		panic(summaryModuleConfigErr)
	}
	summaryProperty.ModuleConfig = summaryModuleConfig

	properties = append(properties, summaryProperty)
	return properties
}

func (s ThingWeaviateModel) WeaviateClassSchemaCrossReferenceProperties() []*models.Property {
	properties := []*models.Property{}

	properties = append(properties, &models.Property{
		Name:     "associatedThing",
		DataType: []string{"Thing2"},
	})
	properties = append(properties, &models.Property{
		Name:     "optionalAssociatedThing",
		DataType: []string{"Thing2"},
	})
	properties = append(properties, &models.Property{
		Name:     "repeatedMessages",
		DataType: []string{"Thing2"},
	})

	return properties
}

func (s ThingWeaviateModel) AllWeaviateClassSchemaProperties() []*models.Property {
	return append(s.WeaviateClassSchemaNonCrossReferenceProperties(), s.WeaviateClassSchemaCrossReferenceProperties()...)
}

func (s ThingWeaviateModel) Data() (map[string]interface{}, error) {
	data := map[string]interface{}{}

	data["aDoubleText"] = fmt.Sprintf("%v", s.ADouble)
	data["aDouble"] = s.ADouble

	data["aFloatText"] = fmt.Sprintf("%v", s.AFloat)
	data["aFloat"] = s.AFloat

	data["anInt32Text"] = fmt.Sprintf("%v", s.AnInt32)
	data["anInt32"] = s.AnInt32

	data["anInt64Text"] = fmt.Sprintf("%v", s.AnInt64)
	data["anInt64"] = strconv.FormatInt(s.AnInt64, 10)

	data["aBoolText"] = fmt.Sprintf("%v", s.ABool)
	data["aBool"] = s.ABool

	data["aString"] = s.AString

	data["aBytes"] = s.ABytes

	data["repeatedScalarFieldText"] = fmt.Sprintf("%v", s.RepeatedScalarField)
	data["repeatedScalarField"] = s.RepeatedScalarField

	if s.OptionalScalarField != nil {
		data["optionalScalarField"] = lo.FromPtr(s.OptionalScalarField)
	}

	data["associatedThing"] = []map[string]string{}

	if s.OptionalAssociatedThing != nil {
		data["optionalAssociatedThing"] = []map[string]string{}
	}

	data["repeatedMessages"] = []map[string]string{}

	data["aTimestampText"] = fmt.Sprintf("%v", s.ATimestamp)
	data["aTimestamp"] = s.ATimestamp

	data["aStructField"] = s.AStructField

	data["anEnumText"] = fmt.Sprintf("%v", s.AnEnum)
	data["anEnum"] = s.AnEnum

	data["anOptionalIntText"] = fmt.Sprintf("%v", lo.FromPtr(s.AnOptionalInt))
	if s.AnOptionalInt != nil {
		data["anOptionalInt"] = lo.FromPtr(s.AnOptionalInt)
	}

	data["optionalTimestampText"] = fmt.Sprintf("%v", lo.FromPtr(s.OptionalTimestamp))
	if s.OptionalTimestamp != nil {
		data["optionalTimestamp"] = lo.FromPtr(s.OptionalTimestamp)
	}

	summary, err := s.SummaryData()
	if err != nil {
		return data, err
	}
	data["_summary"] = summary
	data = s.addCrossReferenceData(data)

	return data, nil
}

func (s ThingWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	if s.AssociatedThing != nil {
		if lo.FromPtr(s.AssociatedThing.Id) != "" {
			id := lo.FromPtr(s.AssociatedThing.Id)
			AssociatedThingReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", id)}
			data["associatedThing"] = append(data["associatedThing"].([]map[string]string), AssociatedThingReference)
		}
	}
	if s.OptionalAssociatedThing != nil {
		if lo.FromPtr(s.OptionalAssociatedThing.Id) != "" {
			id := lo.FromPtr(s.OptionalAssociatedThing.Id)
			OptionalAssociatedThingReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", id)}
			data["optionalAssociatedThing"] = append(data["optionalAssociatedThing"].([]map[string]string), OptionalAssociatedThingReference)
		}
	}
	for _, crossReference := range s.RepeatedMessages {
		if crossReference != nil {
			id := lo.FromPtr(crossReference.Id)
			RepeatedMessagesReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", id)}
			data["repeatedMessages"] = append(data["repeatedMessages"].([]map[string]string), RepeatedMessagesReference)
		}
	}
	return data
}

func (s ThingWeaviateModel) exists(ctx context.Context, client *weaviate.Client) (bool, error) {
	return client.Data().Checker().WithID(lo.FromPtr(s.Id)).WithClassName(s.WeaviateClassName()).Do(ctx)
}

func (s ThingWeaviateModel) Upsert(ctx context.Context, client *weaviate.Client, consistencyLevel string) (data *data.ObjectWrapper, err error) {
	data, err = s.Create(ctx, client, consistencyLevel)
	if err != nil && strings.Contains(err.Error(), "already exists") {
		err = s.Update(ctx, client, consistencyLevel)
	}
	return
}

func (s ThingWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (data *data.ObjectWrapper, err error) {
	if s.AssociatedThing != nil {
		_, err = s.AssociatedThing.Upsert(ctx, client, consistencyLevel)
		if err != nil {
			return
		}
	}
	if s.OptionalAssociatedThing != nil {
		_, err = s.OptionalAssociatedThing.Upsert(ctx, client, consistencyLevel)
		if err != nil {
			return
		}
	}
	for _, crossReference := range s.RepeatedMessages {
		if crossReference != nil {
			_, err = crossReference.Upsert(ctx, client, consistencyLevel)
			if err != nil {
				return
			}
		}
	}
	var dataMap map[string]interface{}
	if dataMap, err = s.Data(); err != nil {
		return
	}
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(dataMap).
		WithID(lo.FromPtr(s.Id)).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s ThingWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) (err error) {
	if s.AssociatedThing != nil {
		_, err = s.AssociatedThing.Upsert(ctx, client, consistencyLevel)
		if err != nil {
			return
		}
	}
	if s.OptionalAssociatedThing != nil {
		_, err = s.OptionalAssociatedThing.Upsert(ctx, client, consistencyLevel)
		if err != nil {
			return
		}
	}
	for _, crossReference := range s.RepeatedMessages {
		if crossReference != nil {
			_, err = crossReference.Upsert(ctx, client, consistencyLevel)
			if err != nil {
				return
			}
		}
	}
	var dataMap map[string]interface{}
	if dataMap, err = s.Data(); err != nil {
		return
	}
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(lo.FromPtr(s.Id)).
		WithProperties(dataMap).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s ThingWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(lo.FromPtr(s.Id)).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s ThingWeaviateModel) EnsureFullClass(client *weaviate.Client, continueOnError bool) (err error) {
	if err = s.EnsureClassWithoutCrossReferences(client, continueOnError); err != nil {
		return
	}
	return s.EnsureClassWithCrossReferences(client, continueOnError)
}

func (s ThingWeaviateModel) EnsureClassWithoutCrossReferences(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.NonCrossReferenceWeaviateClassSchema(), continueOnError)
}

func (s ThingWeaviateModel) EnsureClassWithCrossReferences(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.CrossReferenceWeaviateClassSchema(), continueOnError)
}

func (s ThingWeaviateModel) SummaryData() (string, error) {
	return getStringValue(s)
}

type Thing2WeaviateModel struct {

	// @gotags: fake:"{uuid}"
	Id *string `json:"id" fake:"{uuid}"`

	// @gotags: fake:"{name}"
	Name string `json:"name" fake:"{name}"`
}

func (s Thing2WeaviateModel) ToProto() (theProto *Thing2, err error) {
	theProto = &Thing2{}

	if s.Id != nil {
		theProto.Id = s.Id

	}

	theProto.Name = s.Name

	return
}

func (s *Thing2) ToWeaviateModel() (model *Thing2WeaviateModel, err error) {
	model = &Thing2WeaviateModel{}

	model.Id = s.Id

	model.Name = s.Name

	return
}

func (s Thing2WeaviateModel) WeaviateClassName() string {
	return "Thing2"
}

func (s Thing2WeaviateModel) FullWeaviateClassSchema() models.Class {
	class := models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.AllWeaviateClassSchemaProperties(),
	}

	return class
}

func (s Thing2WeaviateModel) CrossReferenceWeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaCrossReferenceProperties(),
	}
}

func (s Thing2WeaviateModel) NonCrossReferenceWeaviateClassSchema() models.Class {
	class := models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaNonCrossReferenceProperties(),
	}

	return class
}

func (s Thing2WeaviateModel) WeaviateClassSchemaNonCrossReferenceProperties() []*models.Property {
	properties := []*models.Property{}

	NameProperty := &models.Property{
		Name:     "name",
		DataType: []string{"text"},
	}

	properties = append(properties, NameProperty)

	return properties
}

func (s Thing2WeaviateModel) WeaviateClassSchemaCrossReferenceProperties() []*models.Property {
	properties := []*models.Property{}

	return properties
}

func (s Thing2WeaviateModel) AllWeaviateClassSchemaProperties() []*models.Property {
	return append(s.WeaviateClassSchemaNonCrossReferenceProperties(), s.WeaviateClassSchemaCrossReferenceProperties()...)
}

func (s Thing2WeaviateModel) Data() (map[string]interface{}, error) {
	data := map[string]interface{}{}

	data["name"] = s.Name

	data = s.addCrossReferenceData(data)

	return data, nil
}

func (s Thing2WeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	return data
}

func (s Thing2WeaviateModel) exists(ctx context.Context, client *weaviate.Client) (bool, error) {
	return client.Data().Checker().WithID(lo.FromPtr(s.Id)).WithClassName(s.WeaviateClassName()).Do(ctx)
}

func (s Thing2WeaviateModel) Upsert(ctx context.Context, client *weaviate.Client, consistencyLevel string) (data *data.ObjectWrapper, err error) {
	data, err = s.Create(ctx, client, consistencyLevel)
	if err != nil && strings.Contains(err.Error(), "already exists") {
		err = s.Update(ctx, client, consistencyLevel)
	}
	return
}

func (s Thing2WeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (data *data.ObjectWrapper, err error) {
	var dataMap map[string]interface{}
	if dataMap, err = s.Data(); err != nil {
		return
	}
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(dataMap).
		WithID(lo.FromPtr(s.Id)).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s Thing2WeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) (err error) {
	var dataMap map[string]interface{}
	if dataMap, err = s.Data(); err != nil {
		return
	}
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(lo.FromPtr(s.Id)).
		WithProperties(dataMap).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s Thing2WeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(lo.FromPtr(s.Id)).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s Thing2WeaviateModel) EnsureFullClass(client *weaviate.Client, continueOnError bool) (err error) {
	if err = s.EnsureClassWithoutCrossReferences(client, continueOnError); err != nil {
		return
	}
	return s.EnsureClassWithCrossReferences(client, continueOnError)
}

func (s Thing2WeaviateModel) EnsureClassWithoutCrossReferences(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.NonCrossReferenceWeaviateClassSchema(), continueOnError)
}

func (s Thing2WeaviateModel) EnsureClassWithCrossReferences(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.CrossReferenceWeaviateClassSchema(), continueOnError)
}

func (s Thing2WeaviateModel) SummaryData() (string, error) {
	return getStringValue(s)
}

func EnsureClasses(client *weaviate.Client, continueOnError bool) (err error) {
	// create classes without cross references first so there are no errors about missing classes
	err = ThingWeaviateModel{}.EnsureClassWithoutCrossReferences(client, continueOnError)
	if !continueOnError && err != nil {
		return
	}
	err = Thing2WeaviateModel{}.EnsureClassWithoutCrossReferences(client, continueOnError)
	if !continueOnError && err != nil {
		return
	}
	// update classes including cross references
	err = ThingWeaviateModel{}.EnsureClassWithCrossReferences(client, continueOnError)
	if !continueOnError && err != nil {
		return
	}
	err = Thing2WeaviateModel{}.EnsureClassWithCrossReferences(client, continueOnError)
	if !continueOnError && err != nil {
		return
	}
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
