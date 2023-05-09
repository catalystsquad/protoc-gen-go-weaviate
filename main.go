package main

import (
	"github.com/catalystsquad/protoc-gen-go-weaviate/modules"
	pgs "github.com/lyft/protoc-gen-star/v2"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		modules.NewWeaviateModule(),
	).RegisterPostProcessor().Render()
}
