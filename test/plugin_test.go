package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/catalystsquad/app-utils-go/errorutils"
	. "github.com/catalystsquad/protoc-gen-go-weaviate/example"
	"github.com/google/go-cmp/cmp"
	"github.com/orlangure/gnomock"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	client "github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/data/replication"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"
	"google.golang.org/protobuf/testing/protocmp"
	"net/http"
	"strings"
	"testing"
)

const weaviateScheme = "http"
const weaviateHost = "localhost:8080"

var thingClass = lo.ToPtr(Thing{}).ToWeaviateModel().WeaviateClassName()
var thing2Class = lo.ToPtr(Thing2{}).ToWeaviateModel().WeaviateClassName()
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
}
func (s *PluginSuite) TestPlugin() {
	err := Thing2WeaviateModel{}.EnsureClass(weaviateClient, false)
	require.NoError(s.T(), err)
	err = ThingWeaviateModel{}.EnsureClass(weaviateClient, false)
	require.NoError(s.T(), err)
	// create protos
	thing := &Thing{}
	associatedThing1 := Thing2{}
	associatedThing2 := Thing2{}
	associatedThing3 := Thing2{}
	associatedThing4 := Thing2{}
	// populate protos
	err = gofakeit.Struct(&thing)
	require.NoError(s.T(), err)
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
	// create associated things
	for _, thing2 := range thing.RepeatedMessages {
		_, err = thing2.ToWeaviateModel().Create(context.Background(), weaviateClient, replication.ConsistencyLevel.ONE)
		require.NoError(s.T(), err)
	}
	_, err = thing.AssociatedThing.ToWeaviateModel().Create(context.Background(), weaviateClient, replication.ConsistencyLevel.ONE)
	require.NoError(s.T(), err)
	_, err = thing.OptionalAssociatedThing.ToWeaviateModel().Create(context.Background(), weaviateClient, replication.ConsistencyLevel.ONE)
	require.NoError(s.T(), err)
	// create thing
	_, err = thing.ToWeaviateModel().Create(context.Background(), weaviateClient, replication.ConsistencyLevel.ONE)
	require.NoError(s.T(), err)
	// query for thing
	response := s.queryForThings()
	things, err := ThingWeaviateModelsFromGraphqlResult(response)
	resultThing := things[0]
	resultThingProto := resultThing.ToProto()
	assertProtoEquality(s.T(), thing, resultThingProto)
	// update related object
	updatedThing := resultThingProto.AssociatedThing
	name := gofakeit.Name()
	require.NotEqual(s.T(), updatedThing.Name, name)
	updatedThing.Name = name
	updatedThingModel := updatedThing.ToWeaviateModel()
	err = updatedThingModel.Update(context.Background(), weaviateClient, replication.ConsistencyLevel.ONE)
	require.NoError(s.T(), err)
	// query again
	postUpdateResponse := s.queryForThings()
	postUpdateThings, err := ThingWeaviateModelsFromGraphqlResult(postUpdateResponse)
	require.NoError(s.T(), err)
	postUpdateResultThing := postUpdateThings[0]
	postUpdateResultThingProto := postUpdateResultThing.ToProto()
	// ensure the update is correct
	assertProtoEquality(s.T(), updatedThing, postUpdateResultThingProto.AssociatedThing)
}

func (s *PluginSuite) deleteClass(class string) {
	err := weaviateClient.Schema().ClassDeleter().WithClassName(class).Do(context.Background())
	if err != nil {
		// ignore not found error
		require.Equal(s.T(), true, strings.Contains(err.Error(), "could not find class"))
	}
}

func (s *PluginSuite) queryForThings() *models.GraphQLResponse {
	return s.queryForClass(thingClass, thingFields)
}

func (s *PluginSuite) queryForClass(class string, fields []graphql.Field) *models.GraphQLResponse {
	response, err := weaviateClient.GraphQL().Get().WithFields(fields...).WithClassName(class).Do(context.Background())
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
		var associatedThing []Thing2WeaviateModel
		if err = getCrossReference(associationsMap, "associatedThing", &associatedThing); err != nil {
			return
		}
		if len(associatedThing) > 0 {
			model.AssociatedThing = associatedThing[0]
		}
		var optionalAssociatedThing []Thing2WeaviateModel
		if err = getCrossReference(associationsMap, "optionalAssociatedThing", &optionalAssociatedThing); err != nil {
			return
		}
		if len(optionalAssociatedThing) > 0 {
			model.OptionalAssociatedThing = &optionalAssociatedThing[0]
		}
		var associatedThings []Thing2WeaviateModel
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
