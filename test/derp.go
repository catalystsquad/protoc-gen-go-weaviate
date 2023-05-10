package test

//
//import (
//	"context"
//	"fmt"
//	"github.com/weaviate/weaviate-go-client/v4/weaviate"
//	"github.com/weaviate/weaviate-go-client/v4/weaviate/data"
//	"github.com/weaviate/weaviate/entities/models"
//)
//
//type WeaviateThing struct {
//	Id              string
//	Name            string
//	AssociatedThing Thing2
//}
//
//func (s WeaviateThing) WeaviateClassName() string {
//	return "Thing"
//}
//
//func (s WeaviateThing) WeaviateClassSchema() models.Class {
//	return models.Class{
//		Class:      s.WeaviateClassName(),
//		Properties: s.WeaviateClassSchemaProperties(),
//	}
//}
//
//func (s WeaviateThing) WeaviateClassSchemaProperties() []*models.Property {
//	return []*models.Property{
//		{
//			Name:     "name",
//			DataType: []string{"string"},
//		},
//		{
//			Name:     "associated_thing",
//			DataType: []string{"Thing2"},
//		},
//	}
//}
//
//func (s WeaviateThing) Data() map[string]interface{} {
//	return map[string]interface{}{
//		"name": s.Name,
//		"associated_thing": map[string]string{
//			"beacon": fmt.Sprintf("weaviate://localhost/%s", s.AssociatedThing.Id),
//		},
//	}
//}
//
//func (s WeaviateThing) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
//	return client.Data().Creator().
//		WithClassName(s.WeaviateClassName()).
//		WithProperties(s.Data()).
//		WithConsistencyLevel(consistencyLevel).
//		Do(ctx)
//}
//
//func (s WeaviateThing) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
//	return client.Data().Updater().
//		WithClassName(s.WeaviateClassName()).
//		WithID(s.Id).
//		WithProperties(s.Data()).
//		WithConsistencyLevel(consistencyLevel).
//		Do(ctx)
//}
//
//func (s WeaviateThing) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
//	return client.Data().Deleter().
//		WithClassName(s.WeaviateClassName()).
//		WithID(s.Id).
//		WithConsistencyLevel(consistencyLevel).
//		Do(ctx)
//}
//
//type Thing2 struct {
//	Id   string
//	Name string
//}
