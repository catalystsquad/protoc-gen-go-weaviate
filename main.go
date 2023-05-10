package main

import (
	"github.com/catalystsquad/protoc-gen-go-weaviate/modules"
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"github.com/samber/lo"
)

func main() {
	pgs.Init(
		pgs.SupportedFeatures(lo.ToPtr(uint64(1))),
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		modules.NewWeaviateModule(),
	).RegisterPostProcessor(pgsgo.GoFmt()).Render()
}
