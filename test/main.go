package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/catalystsquad/app-utils-go/errorutils"
	example "github.com/catalystsquad/protoc-gen-go-weaviate/example"
	client "github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/data/replication"
)

var c = client.New(client.Config{
	Scheme: "http",
	Host:   "localhost:8080",
})

func main() {
	example.Thing2WeaviateModel{}.EnsureClass(c)
	example.ThingWeaviateModel{}.EnsureClass(c)
	// create protos
	thing := example.Thing{}
	associatedThing1 := example.Thing2{}
	associatedThing2 := example.Thing2{}
	associatedThing3 := example.Thing2{}
	err := gofakeit.Struct(&thing)
	errorutils.PanicOnErr(nil, "error generating test data", err)
	err = gofakeit.Struct(&associatedThing1)
	errorutils.PanicOnErr(nil, "error generating test data", err)
	err = gofakeit.Struct(&associatedThing2)
	errorutils.PanicOnErr(nil, "error generating test data", err)
	err = gofakeit.Struct(&associatedThing3)
	errorutils.PanicOnErr(nil, "error generating test data", err)
	thingModel := thing.ToWeaviateModel()
	_, err = thingModel.AssociatedThing.Create(context.Background(), c, replication.ConsistencyLevel.ONE)
	errorutils.PanicOnErr(nil, "error creating thing", err)
	thingModel.ABytes = []byte(gofakeit.HackerPhrase())
	for _, thing := range thingModel.RepeatedMessages {
		_, err = thing.Create(context.Background(), c, replication.ConsistencyLevel.ONE)
		errorutils.PanicOnErr(nil, "error creating thing", err)
	}
	dataBytes, err := json.MarshalIndent(thingModel.Data(), "", "  ")
	errorutils.PanicOnErr(nil, "error marshalling data to json", err)
	fmt.Println(string(dataBytes))
	_, err = thingModel.Create(context.Background(), c, replication.ConsistencyLevel.ONE)
	errorutils.PanicOnErr(nil, "error marshalling create response", err)
	jsonBytes, err := json.Marshal(thingModel)
	errorutils.PanicOnErr(nil, "error marshalling thing", err)
	fmt.Println(string(jsonBytes))
}
