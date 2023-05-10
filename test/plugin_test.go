package test

import (
	"bytes"
	"github.com/catalystsquad/protoc-gen-go-weaviate/modules"
	"github.com/lyft/protoc-gen-star/v2"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
	"text/template"
)

func TestModule(t *testing.T) {
	req, err := os.Open("code_generator_request.pb.bin")
	if err != nil {
		t.Fatal(err)
	}

	fs := afero.NewMemMapFs()
	res := &bytes.Buffer{}

	pgs.Init(
		pgs.ProtocInput(req),  // use the pre-generated request
		pgs.ProtocOutput(res), // capture CodeGeneratorResponse
		pgs.FileSystem(fs),    // capture any custom files written directly to disk
	).RegisterModule(modules.NewWeaviateModule()).RegisterPostProcessor().Render()
	fileString := string(res.Bytes())
	fileString = fileString[strings.Index(string(res.Bytes()), "package"):]
	err = os.WriteFile("test.weaviate.go", []byte(fileString), 0644)
	require.NoError(t, err)
}

func TestTemplate(t *testing.T) {
	template.New("weaviate").Funcs(map[string]interface{}{})
}
