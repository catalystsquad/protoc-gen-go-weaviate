package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/catalystsquad/app-utils-go/errorutils"
	. "github.com/catalystsquad/protoc-gen-go-weaviate/example"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/orlangure/gnomock"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	client "github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/data/replication"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/filters"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"strings"
	"testing"
	"time"
)

const weaviateScheme = "http"
const weaviateHost = "localhost:8080"

var thingClass = ThingWeaviateModel{}.WeaviateClassName()
var thing2Class = Thing2WeaviateModel{}.WeaviateClassName()
var weaviateGraphqlUrl = fmt.Sprintf("%s://%s/v1/graphql", weaviateScheme, weaviateHost)
var httpClient = &http.Client{}
var weaviateClient = client.New(client.Config{
	Scheme: weaviateScheme,
	Host:   weaviateHost,
})

type PluginSuite struct {
	suite.Suite
}

func TestPluginSuite(t *testing.T) {
	suite.Run(t, new(PluginSuite))
}

func (s *PluginSuite) SetupSuite() {
	s.T().Parallel()
	startWeaviate(s.T())
	err := EnsureClasses(weaviateClient, true)
	require.NoError(s.T(), err)
}
func (s *PluginSuite) TestPlugin() {
	// create protos
	thing := &Thing{}
	associatedThing1 := Thing2{}
	associatedThing2 := Thing2{}
	associatedThing3 := Thing2{}
	associatedThing4 := Thing2{}
	// populate protos
	err := gofakeit.Struct(&thing)
	require.NoError(s.T(), err)
	thing.ATimestamp = timestamppb.New(time.Now())
	thing.OptionalTimestamp = timestamppb.New(time.Now())
	thing.AStructField = &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"name": {
				Kind: &structpb.Value_StringValue{
					StringValue: gofakeit.Name(),
				},
			},
		},
	}
	thing.ABytes = []byte(gofakeit.HackerPhrase())
	err = gofakeit.Struct(&associatedThing1)
	require.NoError(s.T(), err)
	err = gofakeit.Struct(&associatedThing2)
	require.NoError(s.T(), err)
	err = gofakeit.Struct(&associatedThing3)
	require.NoError(s.T(), err)
	err = gofakeit.Struct(&associatedThing4)
	require.NoError(s.T(), err)
	// set associated protos
	thing.AssociatedThing = &associatedThing1
	thing.OptionalAssociatedThing = &associatedThing2
	thing.RepeatedMessages = []*Thing2{&associatedThing3, &associatedThing4}
	require.NoError(s.T(), err)
	optionalAssociatedThingModel, err := thing.OptionalAssociatedThing.ToWeaviateModel()
	require.NoError(s.T(), err)
	_, err = optionalAssociatedThingModel.Create(context.Background(), weaviateClient, replication.ConsistencyLevel.ALL)
	require.NoError(s.T(), err)
	// create thing
	thingModel, err := thing.ToWeaviateModel()
	require.NoError(s.T(), err)
	_, err = thingModel.Create(context.Background(), weaviateClient, replication.ConsistencyLevel.ALL)
	require.NoError(s.T(), err)
	// query for thing
	response := s.queryForThings(nil)
	things, err := ThingWeaviateModelsFromGraphqlResult(response)
	thingsMap := lo.KeyBy(things, func(item ThingWeaviateModel) string {
		return *item.Id
	})
	resultThing := thingsMap[*thing.Id]
	require.NotNil(s.T(), resultThing)
	resultThingProto, err := resultThing.ToProto()
	require.NoError(s.T(), err)
	assertProtoEquality(s.T(), thing, resultThingProto)
	// test summary search by querying the summary field for the name of the associated thing
	query := fmt.Sprintf("*%s*", strings.ToLower(thing.AssociatedThing.Name))
	where := &filters.WhereBuilder{}
	where = where.WithPath([]string{"_summary"}).WithOperator(filters.Like).WithValueText(query)
	summarySearchResponse := s.queryForThings(where)
	summarySearchThings, err := ThingWeaviateModelsFromGraphqlResult(summarySearchResponse)
	summarySearchThingsMap := lo.KeyBy(summarySearchThings, func(item ThingWeaviateModel) string {
		return *item.Id
	})
	summarySearchresultThing := summarySearchThingsMap[*thing.Id]
	require.NotNil(s.T(), resultThing)
	summarySearchresultThingProto, err := summarySearchresultThing.ToProto()
	require.NoError(s.T(), err)
	assertProtoEquality(s.T(), thing, summarySearchresultThingProto)
	// update related object
	updatedThing := resultThingProto.AssociatedThing
	name := gofakeit.Name()
	require.NotEqual(s.T(), updatedThing.Name, name)
	updatedThing.Name = name
	updatedThingModel, err := updatedThing.ToWeaviateModel()
	require.NoError(s.T(), err)
	err = updatedThingModel.Update(context.Background(), weaviateClient, replication.ConsistencyLevel.ALL)
	require.NoError(s.T(), err)
	// query again
	postUpdateResponse := s.queryForThings(nil)
	postUpdateThings, err := ThingWeaviateModelsFromGraphqlResult(postUpdateResponse)
	require.NoError(s.T(), err)
	postUpdateResultThingMap := lo.KeyBy(postUpdateThings, func(item ThingWeaviateModel) string {
		return *item.Id
	})
	postUpdateResultThing := postUpdateResultThingMap[*thing.Id]
	postUpdateResultThingProto, err := postUpdateResultThing.ToProto()
	require.NoError(s.T(), err)
	// ensure the update is correct
	assertProtoEquality(s.T(), updatedThing, postUpdateResultThingProto.AssociatedThing)
}

func (s *PluginSuite) TestUpsert() {
	// create protos
	thing := &Thing{Id: lo.ToPtr(uuid.New().String())}
	associatedThing := &Thing2{Id: lo.ToPtr(uuid.New().String())}
	// populate protos
	err := gofakeit.Struct(&thing)
	require.NoError(s.T(), err)
	err = gofakeit.Struct(&associatedThing)
	require.NoError(s.T(), err)
	thing.AssociatedThing = associatedThing
	thing.ATimestamp = timestamppb.New(time.Now())
	thing.OptionalTimestamp = timestamppb.New(time.Now())
	thing.AStructField = &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"name": {
				Kind: &structpb.Value_StringValue{
					StringValue: gofakeit.Name(),
				},
			},
		},
	}
	thing.ABytes = []byte(gofakeit.HackerPhrase())
	// upsert thing
	thingModel, err := thing.ToWeaviateModel()
	require.NoError(s.T(), err)
	_, err = thingModel.Upsert(context.Background(), weaviateClient, replication.ConsistencyLevel.ALL)
	require.NoError(s.T(), err)
	// query for thing
	response := s.queryForThings(nil)
	things, err := ThingWeaviateModelsFromGraphqlResult(response)
	thingsMap := lo.KeyBy(things, func(item ThingWeaviateModel) string {
		return *item.Id
	})
	resultThing := thingsMap[*thing.Id]
	resultThingProto, err := resultThing.ToProto()
	require.NoError(s.T(), err)
	assertProtoEquality(s.T(), thing, resultThingProto, protocmp.IgnoreFields(&Thing{}, "associated_thing"))
	// update
	thing.AString = gofakeit.HackerPhrase()
	updatedModel, err := thing.ToWeaviateModel()
	require.NoError(s.T(), err)
	_, err = updatedModel.Upsert(context.Background(), weaviateClient, replication.ConsistencyLevel.ALL)
	require.NoError(s.T(), err)
	// query again
	postUpdateResponse := s.queryForThings(nil)
	postUpdateThings, err := ThingWeaviateModelsFromGraphqlResult(postUpdateResponse)
	require.NoError(s.T(), err)
	var postUpdateResultThing *Thing
	for _, postUpdateThing := range postUpdateThings {
		if lo.FromPtr(postUpdateThing.Id) == lo.FromPtr(thing.Id) {
			postUpdateResultThing, err = postUpdateThing.ToProto()
			require.NoError(s.T(), err)
			break
		}
	}
	require.NotNil(s.T(), postUpdateResultThing)
	require.NoError(s.T(), err)
	// ensure the update is correct
	assertProtoEquality(s.T(), thing, postUpdateResultThing, protocmp.IgnoreFields(&Thing{}, "associated_thing"))
}

func (s *PluginSuite) deleteClass(class string) {
	err := weaviateClient.Schema().ClassDeleter().WithClassName(class).Do(context.Background())
	if err != nil {
		// ignore not found error
		require.Equal(s.T(), true, strings.Contains(err.Error(), "could not find class"))
	}
}

func (s *PluginSuite) queryForThings(where *filters.WhereBuilder) *models.GraphQLResponse {
	return s.queryForClass(thingClass, thingFields, where)
}

func (s *PluginSuite) queryForClass(class string, fields []graphql.Field, where *filters.WhereBuilder) *models.GraphQLResponse {
	builder := weaviateClient.GraphQL().Get().WithFields(fields...).WithClassName(class)
	if where != nil {
		builder = builder.WithWhere(where)
	}
	response, err := builder.Do(context.Background())
	require.NoError(s.T(), err)
	require.Len(s.T(), response.Errors, 0)
	return response
}

var thingFields = []graphql.Field{
	{
		Name: "_additional",
		Fields: []graphql.Field{
			{Name: "id"},
		},
	},
	{Name: "aBool"},
	{Name: "aBytes"},
	{Name: "aDouble"},
	{Name: "aFloat"},
	{Name: "aString"},
	{Name: "anInt32"},
	{Name: "anInt64"},
	{Name: "optionalScalarField"},
	{Name: "repeatedScalarField"},
	{
		Name: "associatedThing",
		Fields: []graphql.Field{
			{
				Name: "... on Thing2",
				Fields: []graphql.Field{
					{
						Name:   "_additional",
						Fields: []graphql.Field{{Name: "id"}},
					},
					{Name: "name"}},
			},
		},
	},
	{
		Name: "optionalAssociatedThing",
		Fields: []graphql.Field{
			{
				Name: "... on Thing2",
				Fields: []graphql.Field{
					{
						Name:   "_additional",
						Fields: []graphql.Field{{Name: "id"}},
					},
					{Name: "name"}},
			},
		},
	},
	{
		Name: "repeatedMessages",
		Fields: []graphql.Field{
			{
				Name: "... on Thing2",
				Fields: []graphql.Field{
					{
						Name:   "_additional",
						Fields: []graphql.Field{{Name: "id"}},
					},
					{Name: "name"}},
			},
		},
	},
	{Name: "aStructField"},
	{Name: "anEnum"},
	{Name: "anOptionalInt"},
	{Name: "aTimestamp"},
	{Name: "optionalTimestamp"},
}

func convertType(source, dest interface{}) error {
	bytes, err := json.Marshal(source)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &dest)
	return err
}

func ThingWeaviateModelsFromGraphqlResult(response *models.GraphQLResponse) (models []ThingWeaviateModel, err error) {
	var data []map[string]interface{}
	var dataBytes []byte
	responseObjects := response.Data["Get"].(map[string]interface{})["Thing"]
	dataBytes, err = json.Marshal(responseObjects)
	if err != nil {
		return
	}
	if err = json.Unmarshal(dataBytes, &data); err != nil {
		return
	}
	for _, obj := range data {
		associationsMap := map[string]interface{}{}
		associationsMap["associatedThing"] = obj["associatedThing"]
		delete(obj, "associatedThing")
		associationsMap["optionalAssociatedThing"] = obj["optionalAssociatedThing"]
		delete(obj, "optionalAssociatedThing")
		associationsMap["repeatedMessages"] = obj["repeatedMessages"]
		delete(obj, "repeatedMessages")
		obj["id"] = getIdFromAdditional(obj)
		var model ThingWeaviateModel
		if err = convertType(obj, &model); err != nil {
			return
		}
		var associatedThing []*Thing2WeaviateModel
		if err = getCrossReference(associationsMap, "associatedThing", &associatedThing); err != nil {
			return
		}
		if len(associatedThing) > 0 {
			model.AssociatedThing = associatedThing[0]
		}
		var optionalAssociatedThing []*Thing2WeaviateModel
		if err = getCrossReference(associationsMap, "optionalAssociatedThing", &optionalAssociatedThing); err != nil {
			return
		}
		if len(optionalAssociatedThing) > 0 {
			model.OptionalAssociatedThing = optionalAssociatedThing[0]
		}
		var associatedThings []*Thing2WeaviateModel
		if err = getCrossReference(associationsMap, "repeatedMessages", &associatedThings); err != nil {
			return
		}
		if len(associatedThings) > 0 {
			model.RepeatedMessages = associatedThings
		}
		models = append(models, model)
	}
	return
}

func getCrossReference(obj map[string]interface{}, fieldName string, dest interface{}) (err error) {
	x, ok := obj[fieldName].([]interface{})
	if ok {
		for i, z := range x {
			zMap := z.(map[string]interface{})
			id := getIdFromAdditional(zMap)
			zMap["id"] = id
			x[i] = zMap
		}
		err = convertType(x, dest)
		return
	}
	return
}

func getIdFromAdditional(obj map[string]interface{}) (id string) {
	additional, ok := obj["_additional"]
	if ok {
		id = additional.(map[string]interface{})["id"].(string)
	}
	return
}

func startWeaviate(t *testing.T) {
	container, err := gnomock.StartCustom(
		"semitechnologies/weaviate",
		gnomock.NamedPorts{"default": gnomock.Port{
			Protocol: "tcp",
			Port:     8080,
			HostPort: 8080,
		}},
		gnomock.WithEnv("AUTHENTICATION_ANONYMOUS_ACCESS_ENABLED=true"),
		gnomock.WithEnv("DEFAULT_VECTORIZER_MODULE=none"),
		gnomock.WithEnv("PERSISTENCE_DATA_PATH=/tmp/weaviate"),
	)
	require.NoError(t, err)
	t.Cleanup(func() {
		err := gnomock.Stop(container)
		errorutils.LogOnErr(nil, "error stopping weaviate container", err)
	})
}

func assertProtoEquality(t *testing.T, expected, actual interface{}, options ...cmp.Option) {
	opts := append([]cmp.Option{
		protocmp.Transform(),
	}, options...)
	diff := cmp.Diff(expected, actual, opts...)
	require.Empty(t, diff)
}
