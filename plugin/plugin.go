package plugin

import (
	"bytes"
	"fmt"
	weaviate "github.com/catalystsquad/protoc-gen-go-weaviate/options"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"reflect"
	"strings"
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

type Builder struct {
	plugin         *protogen.Plugin
	messages       map[string]struct{}
	currentFile    string
	currentPackage string
	dbEngine       int
	stringEnums    bool
	suppressWarn   bool
}

var g *protogen.GeneratedFile

// I can't find where the constant is for this in protogen, so I'm putting it here
const SUPPORTS_OPTIONAL_FIELDS = 1

var templateFuncs = map[string]any{
	"protoStructName":          getProtoStructName,
	"protoStructTypeFromField": getProtoStructTypeFromField,
	"protoStructType":          getProtoStructType,
	"structName":               getStructName,
	"structField":              getStructField,
	"structFieldName":          getStructFieldName,
	"structFieldType":          getStructFieldType,
	"className":                getClassName,
	"fieldClassName":           getFieldClassName,
	"propertyName":             getPropertyName,
	"propertyDataType":         getPropertyDataType,
	"dataField":                getDataField,
	"fieldIsCrossReference":    isStructType,
	"fieldIsMessage":           isStructType,
	"fieldIsRepeated":          fieldIsRepeated,
	"fieldComments":            getFieldComments,
	"jsonTag":                  getJsonTag,
	"jsonFieldName":            jsonFieldName,
	"fieldIsOptional":          fieldIsOptional,
	"weaviateModelReturnType":  getWeaviateModelReturnType,
	"includeField":             includeField,
	"isTimestamp":              isTimestamp,
	"shouldGenerateMessage":    shouldGenerateMessage,
	"shouldGenerateFile":       shouldGenerateFile,
	"isStructPb":               isStructPb,
}

func New(opts protogen.Options, request *pluginpb.CodeGeneratorRequest) (*Builder, error) {
	plugin, err := opts.New(request)
	if err != nil {
		return nil, err
	}
	plugin.SupportedFeatures = SUPPORTS_OPTIONAL_FIELDS
	builder := &Builder{
		plugin:   plugin,
		messages: make(map[string]struct{}),
	}

	params := parseParameter(request.GetParameter())

	if strings.EqualFold(params["enums"], "string") {
		builder.stringEnums = true
	}

	if _, ok := params["quiet"]; ok {
		builder.suppressWarn = true
	}

	return builder, nil
}

func parseParameter(param string) map[string]string {
	paramMap := make(map[string]string)

	params := strings.Split(param, ",")
	for _, param := range params {
		if strings.Contains(param, "=") {
			kv := strings.Split(param, "=")
			paramMap[kv[0]] = kv[1]
			continue
		}
		paramMap[param] = ""
	}

	return paramMap
}

func (b *Builder) Generate() (response *pluginpb.CodeGeneratorResponse, err error) {
	for _, protoFile := range b.plugin.Files {
		if shouldGenerateFile(protoFile) {
			var tpl *template.Template
			templateFuncs["package"] = func() string { return string(protoFile.GoPackageName) }
			if tpl, err = template.New("weaviate").Funcs(templateFuncs).Parse(WeaviateTemplate); err != nil {
				return
			}
			fileName := protoFile.GeneratedFilenamePrefix + ".pb.weaviate.go"
			g = b.plugin.NewGeneratedFile(fileName, ".")
			var data bytes.Buffer
			templateMap := map[string]any{
				"messages": protoFile.Messages,
			}
			if err = tpl.Execute(&data, templateMap); err != nil {
				return
			}
			if _, err = g.Write(data.Bytes()); err != nil {
				return
			}
		}
	}
	response = b.plugin.Response()
	return
}

func getStructName(m *protogen.Message) string {
	return fmt.Sprintf("%sWeaviateModel", m.Desc.Name())
}

func getProtoStructName(m *protogen.Message) protoreflect.Name {
	return m.Desc.Name()
}

func getProtoStructType(m *protogen.Message) protoreflect.Name {
	return m.Desc.Name()
}

func getProtoStructTypeFromField(field *protogen.Field) protoreflect.Name {
	return getProtoStructType(field.Message)
}

func getStructField(field *protogen.Field) string {
	return fmt.Sprintf("%s %s", getStructFieldName(field), getStructFieldType(field))
}

func getStructFieldName(field *protogen.Field) string {
	return field.GoName
}

func getStructFieldType(field *protogen.Field) (datatype string) {
	if isTimestamp(field) {
		g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "time"})
		g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "google.golang.org/protobuf/types/known/timestamppb"})
		datatype = "*time.Time"
	} else if isStructPb(field) {
		datatype = "string"
	} else if isStructType(field) {
		datatype = getStructFieldStructType(field)
	} else {
		datatype = getStructFieldNonStructType(field)
	}
	if fieldIsRepeated(field) {
		datatype = fmt.Sprintf("[]%s", datatype)
	}
	if isPointer(field) {
		datatype = fmt.Sprintf("*%s", datatype)
	}
	return
}

func isStructType(field *protogen.Field) bool {
	return field.Message != nil && !isTimestamp(field) && !isStructPb(field)
}

func getStructFieldStructType(field *protogen.Field) string {
	return getStructName(field.Message)
}

func getStructFieldNonStructType(field *protogen.Field) string {
	return goTypeMap[field.Desc.Kind()]
}

func fieldIsRepeated(field *protogen.Field) bool {
	return field.Desc.IsList()
}

func isPointer(field *protogen.Field) bool {
	return field.Desc.HasOptionalKeyword()
}

func getClassName(message *protogen.Message) string {
	return message.GoIdent.GoName
}

func getFieldClassName(field *protogen.Field) string {
	return getClassName(field.Message)
}

func getFieldComments(field *protogen.Field) string {
	return field.Comments.Leading.String() + field.Comments.Trailing.String()
}

func getJsonTag(field *protogen.Field) string {
	asString := ""
	if marshallJsonAsString(field) {
		asString = ",string"
	}
	return fmt.Sprintf("`json:\"%s%s\"`", field.Desc.JSONName(), asString)
}

func jsonFieldName(field *protogen.Field) string {
	return field.Desc.JSONName()
}

func marshallJsonAsString(field *protogen.Field) bool {
	return strings.Contains(field.Desc.Kind().GoString(), "64")
}

func fieldIsOptional(field *protogen.Field) bool {
	return field.Desc.HasOptionalKeyword()
}

func getWeaviateModelReturnType(m *protogen.Message) protoreflect.Name {
	return m.Desc.Name()
}

func includeField(f *protogen.Field) bool {
	options := getFieldOptions(f)
	if options != nil {
		return !options.Ignore
	}
	return true // default to include
}

func getPropertyName(field *protogen.Field) string {
	return field.Desc.TextName()
}

func getPropertyDataType(field *protogen.Field) (datatype string) {
	if isStructPb(field) {
		datatype = "text"
		return
	}
	if isStructType(field) {
		datatype = getFieldClassName(field)
		return
	}
	datatype = getNonCrossReferenceType(field)
	if fieldIsRepeated(field) {
		datatype = fmt.Sprintf("%s[]", datatype)
	}
	return
}

func getNonCrossReferenceType(field *protogen.Field) string {
	if isTimestamp(field) {
		return "date"
	}
	return weaviateTypeMap[field.Desc.Kind()]
}

func getDataField(field *protogen.Field) (dataField string) {
	if isStructType(field) {
		return `[]map[string]string{}`
	} else {
		val := fmt.Sprintf(`s.%s`, getStructFieldName(field))
		if isInt64(field) {
			val = fmt.Sprintf("strconv.FormatInt(%s, 10)", val)
		} else if isUint64(field) {
			val = fmt.Sprintf("strconv.FormatUint(%s, 10)", val)
		}
		return val
	}
}

func isInt64(field *protogen.Field) bool {
	kind := field.Desc.Kind()
	return kind == protoreflect.Int64Kind ||
		kind == protoreflect.Sfixed64Kind ||
		kind == protoreflect.Sint64Kind
}

func isUint64(field *protogen.Field) bool {
	return field.Desc.Kind() == protoreflect.Uint64Kind || field.Desc.Kind() == protoreflect.Fixed64Kind
}

var weaviateTypeMap = map[protoreflect.Kind]string{
	protoreflect.StringKind:   "text",
	protoreflect.BoolKind:     "boolean",
	protoreflect.EnumKind:     "int",
	protoreflect.Int32Kind:    "int",
	protoreflect.Sint32Kind:   "int",
	protoreflect.Uint32Kind:   "int",
	protoreflect.Int64Kind:    "string",
	protoreflect.Sint64Kind:   "string",
	protoreflect.Uint64Kind:   "string",
	protoreflect.Sfixed32Kind: "int",
	protoreflect.Fixed32Kind:  "int",
	protoreflect.Sfixed64Kind: "string",
	protoreflect.Fixed64Kind:  "string",
	protoreflect.FloatKind:    "number",
	protoreflect.DoubleKind:   "number",
	protoreflect.BytesKind:    "blob",
}

var goTypeMap = map[protoreflect.Kind]string{
	protoreflect.BoolKind:     "bool",
	protoreflect.EnumKind:     "int",
	protoreflect.Int32Kind:    "int32",
	protoreflect.Sint32Kind:   "int32",
	protoreflect.Uint32Kind:   "uint32",
	protoreflect.Int64Kind:    "int64",
	protoreflect.Sint64Kind:   "int64",
	protoreflect.Uint64Kind:   "uint64",
	protoreflect.Sfixed32Kind: "int32",
	protoreflect.Fixed32Kind:  "uint32",
	protoreflect.FloatKind:    "float32",
	protoreflect.Sfixed64Kind: "int64",
	protoreflect.Fixed64Kind:  "uint64",
	protoreflect.DoubleKind:   "float64",
	protoreflect.StringKind:   "string",
	protoreflect.BytesKind:    "[]byte",
}

func getFieldOptions(field *protogen.Field) *weaviate.WeaviateFieldOptions {
	options := field.Desc.Options().(*descriptorpb.FieldOptions)
	if options == nil {
		return &weaviate.WeaviateFieldOptions{}
	}

	v := proto.GetExtension(options, weaviate.E_WeaviateField)
	if v == nil {
		return nil
	}

	opts, ok := v.(*weaviate.WeaviateFieldOptions)
	if !ok {
		return nil
	}
	return opts
}

func isTimestamp(field *protogen.Field) bool {
	return field.Desc.Message() != nil && field.Desc.Message().FullName() == "google.protobuf.Timestamp"
}

func isStructPb(field *protogen.Field) bool {
	g.QualifiedGoIdent(protogen.GoIdent{GoImportPath: "encoding/json"})
	return field.Desc.Message() != nil && field.Desc.Message().FullName() == "google.protobuf.Struct"
}

func shouldGenerateMessage(message *protogen.Message) bool {
	options := getMessageOptions(message)
	return options.Generate
}

func getMessageOptions(message *protogen.Message) *weaviate.WeaviateMessageOptions {
	options := message.Desc.Options().(*descriptorpb.MessageOptions)
	if options == nil {
		return &weaviate.WeaviateMessageOptions{}
	}

	v := proto.GetExtension(options, weaviate.E_WeaviateOpts)
	if v == nil {
		return nil
	}

	opts, ok := v.(*weaviate.WeaviateMessageOptions)
	if !ok {
		return nil
	}
	return opts
}

func shouldGenerateFile(file *protogen.File) bool {
	options := getFileOptions(file)
	return options != nil && options.Generate
}

func getFileOptions(file *protogen.File) *weaviate.WeaviateFileOptions {
	options := file.Desc.Options().(*descriptorpb.FileOptions)
	if options == nil {
		return &weaviate.WeaviateFileOptions{}
	}
	v := proto.GetExtension(options, weaviate.E_WeaviateFileOpts)
	if reflect.ValueOf(v).IsNil() {
		return nil
	}
	opts, ok := v.(*weaviate.WeaviateFileOptions)
	if !ok {
		return nil
	}
	return opts
}
