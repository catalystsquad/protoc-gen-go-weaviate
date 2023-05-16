build-options:
	buf generate --template proto/options/buf.gen.yaml --path proto/options
build-test:
	go install
	protoc --go-weaviate_out=example example/*.proto
	go install github.com/favadi/protoc-go-inject-tag@latest
	protoc-go-inject-tag -input example/example.example/example.pb.weaviate.go