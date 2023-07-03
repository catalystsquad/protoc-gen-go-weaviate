package descriptorpb

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/data"
	"github.com/weaviate/weaviate/entities/models"
	"strconv"
)

type FileDescriptorSetWeaviateModel struct {
	File []FileDescriptorProtoWeaviateModel `json:"file"`
}

func (s FileDescriptorSetWeaviateModel) ToProto() *FileDescriptorSet {
	theProto := &FileDescriptorSet{}

	for _, protoField := range s.File {
		msg := protoField.ToProto()
		if theProto.File == nil {
			theProto.File = []*FileDescriptorProto{msg}
		} else {
			theProto.File = append(theProto.File, msg)
		}
	}

	return theProto
}

func (s *FileDescriptorSet) ToWeaviateModel() FileDescriptorSetWeaviateModel {
	model := FileDescriptorSetWeaviateModel{}

	for _, protoField := range s.File {
		msg := protoField.ToWeaviateModel()
		if model.File == nil {
			model.File = []FileDescriptorProtoWeaviateModel{msg}
		} else {
			model.File = append(model.File, msg)
		}
	}

	return model
}

func (s FileDescriptorSetWeaviateModel) WeaviateClassName() string {
	return "FileDescriptorSet"
}

func (s FileDescriptorSetWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s FileDescriptorSetWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "file",
		DataType: []string{"FileDescriptorProto"},
	},
	}
}

func (s FileDescriptorSetWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"file": []map[string]string{},
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s FileDescriptorSetWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	for _, crossReference := range s.File {
		FileReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["file"] = append(data["file"].([]map[string]string), FileReference)
	}
	return data
}

func (s FileDescriptorSetWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s FileDescriptorSetWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s FileDescriptorSetWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s FileDescriptorSetWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type FileDescriptorProtoWeaviateModel struct {

	// file name, relative to root of source tree
	Name *string `json:"name"`

	// e.g. "foo", "foo.bar", etc.
	Package *string `json:"package"`

	// Names of files imported by this file.
	Dependency []string `json:"dependency"`

	// Indexes of the public imported files in the dependency list above.
	PublicDependency []int32 `json:"publicDependency"`

	// Indexes of the weak imported files in the dependency list.
	// For Google-internal migration only. Do not use.
	WeakDependency []int32 `json:"weakDependency"`

	// All top-level definitions in this file.
	MessageType []DescriptorProtoWeaviateModel `json:"messageType"`

	EnumType []EnumDescriptorProtoWeaviateModel `json:"enumType"`

	Service []ServiceDescriptorProtoWeaviateModel `json:"service"`

	Extension []FieldDescriptorProtoWeaviateModel `json:"extension"`

	Options *FileOptionsWeaviateModel `json:"options"`

	// This field contains optional information about the original source code.
	// You may safely remove this entire field without harming runtime
	// functionality of the descriptors -- the information is needed only by
	// development tools.
	SourceCodeInfo *SourceCodeInfoWeaviateModel `json:"sourceCodeInfo"`

	// The syntax of the proto file.
	// The supported values are "proto2", "proto3", and "editions".
	//
	// If `edition` is present, this value must be "editions".
	Syntax *string `json:"syntax"`

	// The edition of the proto file, which is an opaque string.
	Edition *string `json:"edition"`
}

func (s FileDescriptorProtoWeaviateModel) ToProto() *FileDescriptorProto {
	theProto := &FileDescriptorProto{}

	theProto.Name = s.Name

	theProto.Package = s.Package

	theProto.Dependency = s.Dependency

	theProto.PublicDependency = s.PublicDependency

	theProto.WeakDependency = s.WeakDependency

	for _, protoField := range s.MessageType {
		msg := protoField.ToProto()
		if theProto.MessageType == nil {
			theProto.MessageType = []*DescriptorProto{msg}
		} else {
			theProto.MessageType = append(theProto.MessageType, msg)
		}
	}

	for _, protoField := range s.EnumType {
		msg := protoField.ToProto()
		if theProto.EnumType == nil {
			theProto.EnumType = []*EnumDescriptorProto{msg}
		} else {
			theProto.EnumType = append(theProto.EnumType, msg)
		}
	}

	for _, protoField := range s.Service {
		msg := protoField.ToProto()
		if theProto.Service == nil {
			theProto.Service = []*ServiceDescriptorProto{msg}
		} else {
			theProto.Service = append(theProto.Service, msg)
		}
	}

	for _, protoField := range s.Extension {
		msg := protoField.ToProto()
		if theProto.Extension == nil {
			theProto.Extension = []*FieldDescriptorProto{msg}
		} else {
			theProto.Extension = append(theProto.Extension, msg)
		}
	}

	theProto.Options = s.Options.ToProto()

	theProto.SourceCodeInfo = s.SourceCodeInfo.ToProto()

	theProto.Syntax = s.Syntax

	theProto.Edition = s.Edition

	return theProto
}

func (s *FileDescriptorProto) ToWeaviateModel() FileDescriptorProtoWeaviateModel {
	model := FileDescriptorProtoWeaviateModel{}

	model.Name = s.Name

	model.Package = s.Package

	model.Dependency = s.Dependency

	model.PublicDependency = s.PublicDependency

	model.WeakDependency = s.WeakDependency

	for _, protoField := range s.MessageType {
		msg := protoField.ToWeaviateModel()
		if model.MessageType == nil {
			model.MessageType = []DescriptorProtoWeaviateModel{msg}
		} else {
			model.MessageType = append(model.MessageType, msg)
		}
	}

	for _, protoField := range s.EnumType {
		msg := protoField.ToWeaviateModel()
		if model.EnumType == nil {
			model.EnumType = []EnumDescriptorProtoWeaviateModel{msg}
		} else {
			model.EnumType = append(model.EnumType, msg)
		}
	}

	for _, protoField := range s.Service {
		msg := protoField.ToWeaviateModel()
		if model.Service == nil {
			model.Service = []ServiceDescriptorProtoWeaviateModel{msg}
		} else {
			model.Service = append(model.Service, msg)
		}
	}

	for _, protoField := range s.Extension {
		msg := protoField.ToWeaviateModel()
		if model.Extension == nil {
			model.Extension = []FieldDescriptorProtoWeaviateModel{msg}
		} else {
			model.Extension = append(model.Extension, msg)
		}
	}

	if s.Options != nil {
		model.Options = lo.ToPtr(s.Options.ToWeaviateModel())
	}

	if s.SourceCodeInfo != nil {
		model.SourceCodeInfo = lo.ToPtr(s.SourceCodeInfo.ToWeaviateModel())
	}

	model.Syntax = s.Syntax

	model.Edition = s.Edition

	return model
}

func (s FileDescriptorProtoWeaviateModel) WeaviateClassName() string {
	return "FileDescriptorProto"
}

func (s FileDescriptorProtoWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s FileDescriptorProtoWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "name",
		DataType: []string{"text"},
	}, {
		Name:     "package",
		DataType: []string{"text"},
	}, {
		Name:     "dependency",
		DataType: []string{"text[]"},
	}, {
		Name:     "publicDependency",
		DataType: []string{"int[]"},
	}, {
		Name:     "weakDependency",
		DataType: []string{"int[]"},
	}, {
		Name:     "messageType",
		DataType: []string{"DescriptorProto"},
	}, {
		Name:     "enumType",
		DataType: []string{"EnumDescriptorProto"},
	}, {
		Name:     "service",
		DataType: []string{"ServiceDescriptorProto"},
	}, {
		Name:     "extension",
		DataType: []string{"FieldDescriptorProto"},
	}, {
		Name:     "options",
		DataType: []string{"FileOptions"},
	}, {
		Name:     "sourceCodeInfo",
		DataType: []string{"SourceCodeInfo"},
	}, {
		Name:     "syntax",
		DataType: []string{"text"},
	}, {
		Name:     "edition",
		DataType: []string{"text"},
	},
	}
}

func (s FileDescriptorProtoWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"name": s.Name,

		"package": s.Package,

		"dependency": s.Dependency,

		"publicDependency": s.PublicDependency,

		"weakDependency": s.WeakDependency,

		"messageType": []map[string]string{},

		"enumType": []map[string]string{},

		"service": []map[string]string{},

		"extension": []map[string]string{},

		"options": []map[string]string{},

		"sourceCodeInfo": []map[string]string{},

		"syntax": s.Syntax,

		"edition": s.Edition,
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s FileDescriptorProtoWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	for _, crossReference := range s.MessageType {
		MessageTypeReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["messageType"] = append(data["messageType"].([]map[string]string), MessageTypeReference)
	}
	for _, crossReference := range s.EnumType {
		EnumTypeReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["enumType"] = append(data["enumType"].([]map[string]string), EnumTypeReference)
	}
	for _, crossReference := range s.Service {
		ServiceReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["service"] = append(data["service"].([]map[string]string), ServiceReference)
	}
	for _, crossReference := range s.Extension {
		ExtensionReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["extension"] = append(data["extension"].([]map[string]string), ExtensionReference)
	}
	if s.Options != nil {
		if s.Options.Id != "" {
			OptionsReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", s.Options.Id)}
			data["options"] = append(data["options"].([]map[string]string), OptionsReference)
		}
	}
	if s.SourceCodeInfo != nil {
		if s.SourceCodeInfo.Id != "" {
			SourceCodeInfoReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", s.SourceCodeInfo.Id)}
			data["sourceCodeInfo"] = append(data["sourceCodeInfo"].([]map[string]string), SourceCodeInfoReference)
		}
	}
	return data
}

func (s FileDescriptorProtoWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s FileDescriptorProtoWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s FileDescriptorProtoWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s FileDescriptorProtoWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type DescriptorProtoWeaviateModel struct {
	Name *string `json:"name"`

	Field []FieldDescriptorProtoWeaviateModel `json:"field"`

	Extension []FieldDescriptorProtoWeaviateModel `json:"extension"`

	NestedType []DescriptorProtoWeaviateModel `json:"nestedType"`

	EnumType []EnumDescriptorProtoWeaviateModel `json:"enumType"`

	ExtensionRange []ExtensionRangeWeaviateModel `json:"extensionRange"`

	OneofDecl []OneofDescriptorProtoWeaviateModel `json:"oneofDecl"`

	Options *MessageOptionsWeaviateModel `json:"options"`

	ReservedRange []ReservedRangeWeaviateModel `json:"reservedRange"`

	// Reserved field names, which may not be used by fields in the same message.
	// A given name may only be reserved once.
	ReservedName []string `json:"reservedName"`
}

func (s DescriptorProtoWeaviateModel) ToProto() *DescriptorProto {
	theProto := &DescriptorProto{}

	theProto.Name = s.Name

	for _, protoField := range s.Field {
		msg := protoField.ToProto()
		if theProto.Field == nil {
			theProto.Field = []*FieldDescriptorProto{msg}
		} else {
			theProto.Field = append(theProto.Field, msg)
		}
	}

	for _, protoField := range s.Extension {
		msg := protoField.ToProto()
		if theProto.Extension == nil {
			theProto.Extension = []*FieldDescriptorProto{msg}
		} else {
			theProto.Extension = append(theProto.Extension, msg)
		}
	}

	for _, protoField := range s.NestedType {
		msg := protoField.ToProto()
		if theProto.NestedType == nil {
			theProto.NestedType = []*DescriptorProto{msg}
		} else {
			theProto.NestedType = append(theProto.NestedType, msg)
		}
	}

	for _, protoField := range s.EnumType {
		msg := protoField.ToProto()
		if theProto.EnumType == nil {
			theProto.EnumType = []*EnumDescriptorProto{msg}
		} else {
			theProto.EnumType = append(theProto.EnumType, msg)
		}
	}

	for _, protoField := range s.ExtensionRange {
		msg := protoField.ToProto()
		if theProto.ExtensionRange == nil {
			theProto.ExtensionRange = []*ExtensionRange{msg}
		} else {
			theProto.ExtensionRange = append(theProto.ExtensionRange, msg)
		}
	}

	for _, protoField := range s.OneofDecl {
		msg := protoField.ToProto()
		if theProto.OneofDecl == nil {
			theProto.OneofDecl = []*OneofDescriptorProto{msg}
		} else {
			theProto.OneofDecl = append(theProto.OneofDecl, msg)
		}
	}

	theProto.Options = s.Options.ToProto()

	for _, protoField := range s.ReservedRange {
		msg := protoField.ToProto()
		if theProto.ReservedRange == nil {
			theProto.ReservedRange = []*ReservedRange{msg}
		} else {
			theProto.ReservedRange = append(theProto.ReservedRange, msg)
		}
	}

	theProto.ReservedName = s.ReservedName

	return theProto
}

func (s *DescriptorProto) ToWeaviateModel() DescriptorProtoWeaviateModel {
	model := DescriptorProtoWeaviateModel{}

	model.Name = s.Name

	for _, protoField := range s.Field {
		msg := protoField.ToWeaviateModel()
		if model.Field == nil {
			model.Field = []FieldDescriptorProtoWeaviateModel{msg}
		} else {
			model.Field = append(model.Field, msg)
		}
	}

	for _, protoField := range s.Extension {
		msg := protoField.ToWeaviateModel()
		if model.Extension == nil {
			model.Extension = []FieldDescriptorProtoWeaviateModel{msg}
		} else {
			model.Extension = append(model.Extension, msg)
		}
	}

	for _, protoField := range s.NestedType {
		msg := protoField.ToWeaviateModel()
		if model.NestedType == nil {
			model.NestedType = []DescriptorProtoWeaviateModel{msg}
		} else {
			model.NestedType = append(model.NestedType, msg)
		}
	}

	for _, protoField := range s.EnumType {
		msg := protoField.ToWeaviateModel()
		if model.EnumType == nil {
			model.EnumType = []EnumDescriptorProtoWeaviateModel{msg}
		} else {
			model.EnumType = append(model.EnumType, msg)
		}
	}

	for _, protoField := range s.ExtensionRange {
		msg := protoField.ToWeaviateModel()
		if model.ExtensionRange == nil {
			model.ExtensionRange = []ExtensionRangeWeaviateModel{msg}
		} else {
			model.ExtensionRange = append(model.ExtensionRange, msg)
		}
	}

	for _, protoField := range s.OneofDecl {
		msg := protoField.ToWeaviateModel()
		if model.OneofDecl == nil {
			model.OneofDecl = []OneofDescriptorProtoWeaviateModel{msg}
		} else {
			model.OneofDecl = append(model.OneofDecl, msg)
		}
	}

	if s.Options != nil {
		model.Options = lo.ToPtr(s.Options.ToWeaviateModel())
	}

	for _, protoField := range s.ReservedRange {
		msg := protoField.ToWeaviateModel()
		if model.ReservedRange == nil {
			model.ReservedRange = []ReservedRangeWeaviateModel{msg}
		} else {
			model.ReservedRange = append(model.ReservedRange, msg)
		}
	}

	model.ReservedName = s.ReservedName

	return model
}

func (s DescriptorProtoWeaviateModel) WeaviateClassName() string {
	return "DescriptorProto"
}

func (s DescriptorProtoWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s DescriptorProtoWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "name",
		DataType: []string{"text"},
	}, {
		Name:     "field",
		DataType: []string{"FieldDescriptorProto"},
	}, {
		Name:     "extension",
		DataType: []string{"FieldDescriptorProto"},
	}, {
		Name:     "nestedType",
		DataType: []string{"DescriptorProto"},
	}, {
		Name:     "enumType",
		DataType: []string{"EnumDescriptorProto"},
	}, {
		Name:     "extensionRange",
		DataType: []string{"DescriptorProto_ExtensionRange"},
	}, {
		Name:     "oneofDecl",
		DataType: []string{"OneofDescriptorProto"},
	}, {
		Name:     "options",
		DataType: []string{"MessageOptions"},
	}, {
		Name:     "reservedRange",
		DataType: []string{"DescriptorProto_ReservedRange"},
	}, {
		Name:     "reservedName",
		DataType: []string{"text[]"},
	},
	}
}

func (s DescriptorProtoWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"name": s.Name,

		"field": []map[string]string{},

		"extension": []map[string]string{},

		"nestedType": []map[string]string{},

		"enumType": []map[string]string{},

		"extensionRange": []map[string]string{},

		"oneofDecl": []map[string]string{},

		"options": []map[string]string{},

		"reservedRange": []map[string]string{},

		"reservedName": s.ReservedName,
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s DescriptorProtoWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	for _, crossReference := range s.Field {
		FieldReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["field"] = append(data["field"].([]map[string]string), FieldReference)
	}
	for _, crossReference := range s.Extension {
		ExtensionReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["extension"] = append(data["extension"].([]map[string]string), ExtensionReference)
	}
	for _, crossReference := range s.NestedType {
		NestedTypeReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["nestedType"] = append(data["nestedType"].([]map[string]string), NestedTypeReference)
	}
	for _, crossReference := range s.EnumType {
		EnumTypeReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["enumType"] = append(data["enumType"].([]map[string]string), EnumTypeReference)
	}
	for _, crossReference := range s.ExtensionRange {
		ExtensionRangeReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["extensionRange"] = append(data["extensionRange"].([]map[string]string), ExtensionRangeReference)
	}
	for _, crossReference := range s.OneofDecl {
		OneofDeclReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["oneofDecl"] = append(data["oneofDecl"].([]map[string]string), OneofDeclReference)
	}
	if s.Options != nil {
		if s.Options.Id != "" {
			OptionsReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", s.Options.Id)}
			data["options"] = append(data["options"].([]map[string]string), OptionsReference)
		}
	}
	for _, crossReference := range s.ReservedRange {
		ReservedRangeReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["reservedRange"] = append(data["reservedRange"].([]map[string]string), ReservedRangeReference)
	}
	return data
}

func (s DescriptorProtoWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s DescriptorProtoWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s DescriptorProtoWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s DescriptorProtoWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type ExtensionRangeOptionsWeaviateModel struct {

	// The parser stores options it doesn't recognize here. See above.
	UninterpretedOption []UninterpretedOptionWeaviateModel `json:"uninterpretedOption"`
}

func (s ExtensionRangeOptionsWeaviateModel) ToProto() *ExtensionRangeOptions {
	theProto := &ExtensionRangeOptions{}

	for _, protoField := range s.UninterpretedOption {
		msg := protoField.ToProto()
		if theProto.UninterpretedOption == nil {
			theProto.UninterpretedOption = []*UninterpretedOption{msg}
		} else {
			theProto.UninterpretedOption = append(theProto.UninterpretedOption, msg)
		}
	}

	return theProto
}

func (s *ExtensionRangeOptions) ToWeaviateModel() ExtensionRangeOptionsWeaviateModel {
	model := ExtensionRangeOptionsWeaviateModel{}

	for _, protoField := range s.UninterpretedOption {
		msg := protoField.ToWeaviateModel()
		if model.UninterpretedOption == nil {
			model.UninterpretedOption = []UninterpretedOptionWeaviateModel{msg}
		} else {
			model.UninterpretedOption = append(model.UninterpretedOption, msg)
		}
	}

	return model
}

func (s ExtensionRangeOptionsWeaviateModel) WeaviateClassName() string {
	return "ExtensionRangeOptions"
}

func (s ExtensionRangeOptionsWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s ExtensionRangeOptionsWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "uninterpretedOption",
		DataType: []string{"UninterpretedOption"},
	},
	}
}

func (s ExtensionRangeOptionsWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"uninterpretedOption": []map[string]string{},
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s ExtensionRangeOptionsWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	for _, crossReference := range s.UninterpretedOption {
		UninterpretedOptionReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["uninterpretedOption"] = append(data["uninterpretedOption"].([]map[string]string), UninterpretedOptionReference)
	}
	return data
}

func (s ExtensionRangeOptionsWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s ExtensionRangeOptionsWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s ExtensionRangeOptionsWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s ExtensionRangeOptionsWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type FieldDescriptorProtoWeaviateModel struct {
	Name *string `json:"name"`

	Number *int32 `json:"number"`

	Label *int `json:"label"`

	// If type_name is set, this need not be set.  If both this and type_name
	// are set, this must be one of TYPE_ENUM, TYPE_MESSAGE or TYPE_GROUP.
	Type *int `json:"type"`

	// For message and enum types, this is the name of the type.  If the name
	// starts with a '.', it is fully-qualified.  Otherwise, C++-like scoping
	// rules are used to find the type (i.e. first the nested types within this
	// message are searched, then within the parent, on up to the root
	// namespace).
	TypeName *string `json:"typeName"`

	// For extensions, this is the name of the type being extended.  It is
	// resolved in the same manner as type_name.
	Extendee *string `json:"extendee"`

	// For numeric types, contains the original text representation of the value.
	// For booleans, "true" or "false".
	// For strings, contains the default text contents (not escaped in any way).
	// For bytes, contains the C escaped value.  All bytes >= 128 are escaped.
	DefaultValue *string `json:"defaultValue"`

	// If set, gives the index of a oneof in the containing type's oneof_decl
	// list.  This field is a member of that oneof.
	OneofIndex *int32 `json:"oneofIndex"`

	// JSON name of this field. The value is set by protocol compiler. If the
	// user has set a "json_name" option on this field, that option's value
	// will be used. Otherwise, it's deduced from the field's name by converting
	// it to camelCase.
	JsonName *string `json:"jsonName"`

	Options *FieldOptionsWeaviateModel `json:"options"`

	// If true, this is a proto3 "optional". When a proto3 field is optional, it
	// tracks presence regardless of field type.
	//
	// When proto3_optional is true, this field must be belong to a oneof to
	// signal to old proto3 clients that presence is tracked for this field. This
	// oneof is known as a "synthetic" oneof, and this field must be its sole
	// member (each proto3 optional field gets its own synthetic oneof). Synthetic
	// oneofs exist in the descriptor only, and do not generate any API. Synthetic
	// oneofs must be ordered after all "real" oneofs.
	//
	// For message fields, proto3_optional doesn't create any semantic change,
	// since non-repeated message fields always track presence. However it still
	// indicates the semantic detail of whether the user wrote "optional" or not.
	// This can be useful for round-tripping the .proto file. For consistency we
	// give message fields a synthetic oneof also, even though it is not required
	// to track presence. This is especially important because the parser can't
	// tell if a field is a message or an enum, so it must always create a
	// synthetic oneof.
	//
	// Proto2 optional fields do not set this flag, because they already indicate
	// optional with `LABEL_OPTIONAL`.
	Proto3Optional *bool `json:"proto3Optional"`
}

func (s FieldDescriptorProtoWeaviateModel) ToProto() *FieldDescriptorProto {
	theProto := &FieldDescriptorProto{}

	theProto.Name = s.Name

	theProto.Number = s.Number

	theProto.Label = s.Label

	theProto.Type = s.Type

	theProto.TypeName = s.TypeName

	theProto.Extendee = s.Extendee

	theProto.DefaultValue = s.DefaultValue

	theProto.OneofIndex = s.OneofIndex

	theProto.JsonName = s.JsonName

	theProto.Options = s.Options.ToProto()

	theProto.Proto3Optional = s.Proto3Optional

	return theProto
}

func (s *FieldDescriptorProto) ToWeaviateModel() FieldDescriptorProtoWeaviateModel {
	model := FieldDescriptorProtoWeaviateModel{}

	model.Name = s.Name

	model.Number = s.Number

	model.Label = s.Label

	model.Type = s.Type

	model.TypeName = s.TypeName

	model.Extendee = s.Extendee

	model.DefaultValue = s.DefaultValue

	model.OneofIndex = s.OneofIndex

	model.JsonName = s.JsonName

	if s.Options != nil {
		model.Options = lo.ToPtr(s.Options.ToWeaviateModel())
	}

	model.Proto3Optional = s.Proto3Optional

	return model
}

func (s FieldDescriptorProtoWeaviateModel) WeaviateClassName() string {
	return "FieldDescriptorProto"
}

func (s FieldDescriptorProtoWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s FieldDescriptorProtoWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "name",
		DataType: []string{"text"},
	}, {
		Name:     "number",
		DataType: []string{"int"},
	}, {
		Name:     "label",
		DataType: []string{"int"},
	}, {
		Name:     "type",
		DataType: []string{"int"},
	}, {
		Name:     "typeName",
		DataType: []string{"text"},
	}, {
		Name:     "extendee",
		DataType: []string{"text"},
	}, {
		Name:     "defaultValue",
		DataType: []string{"text"},
	}, {
		Name:     "oneofIndex",
		DataType: []string{"int"},
	}, {
		Name:     "jsonName",
		DataType: []string{"text"},
	}, {
		Name:     "options",
		DataType: []string{"FieldOptions"},
	}, {
		Name:     "proto3Optional",
		DataType: []string{"boolean"},
	},
	}
}

func (s FieldDescriptorProtoWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"name": s.Name,

		"number": s.Number,

		"label": s.Label,

		"type": s.Type,

		"typeName": s.TypeName,

		"extendee": s.Extendee,

		"defaultValue": s.DefaultValue,

		"oneofIndex": s.OneofIndex,

		"jsonName": s.JsonName,

		"options": []map[string]string{},

		"proto3Optional": s.Proto3Optional,
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s FieldDescriptorProtoWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	if s.Options != nil {
		if s.Options.Id != "" {
			OptionsReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", s.Options.Id)}
			data["options"] = append(data["options"].([]map[string]string), OptionsReference)
		}
	}
	return data
}

func (s FieldDescriptorProtoWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s FieldDescriptorProtoWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s FieldDescriptorProtoWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s FieldDescriptorProtoWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type OneofDescriptorProtoWeaviateModel struct {
	Name *string `json:"name"`

	Options *OneofOptionsWeaviateModel `json:"options"`
}

func (s OneofDescriptorProtoWeaviateModel) ToProto() *OneofDescriptorProto {
	theProto := &OneofDescriptorProto{}

	theProto.Name = s.Name

	theProto.Options = s.Options.ToProto()

	return theProto
}

func (s *OneofDescriptorProto) ToWeaviateModel() OneofDescriptorProtoWeaviateModel {
	model := OneofDescriptorProtoWeaviateModel{}

	model.Name = s.Name

	if s.Options != nil {
		model.Options = lo.ToPtr(s.Options.ToWeaviateModel())
	}

	return model
}

func (s OneofDescriptorProtoWeaviateModel) WeaviateClassName() string {
	return "OneofDescriptorProto"
}

func (s OneofDescriptorProtoWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s OneofDescriptorProtoWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "name",
		DataType: []string{"text"},
	}, {
		Name:     "options",
		DataType: []string{"OneofOptions"},
	},
	}
}

func (s OneofDescriptorProtoWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"name": s.Name,

		"options": []map[string]string{},
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s OneofDescriptorProtoWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	if s.Options != nil {
		if s.Options.Id != "" {
			OptionsReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", s.Options.Id)}
			data["options"] = append(data["options"].([]map[string]string), OptionsReference)
		}
	}
	return data
}

func (s OneofDescriptorProtoWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s OneofDescriptorProtoWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s OneofDescriptorProtoWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s OneofDescriptorProtoWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type EnumDescriptorProtoWeaviateModel struct {
	Name *string `json:"name"`

	Value []EnumValueDescriptorProtoWeaviateModel `json:"value"`

	Options *EnumOptionsWeaviateModel `json:"options"`

	// Range of reserved numeric values. Reserved numeric values may not be used
	// by enum values in the same enum declaration. Reserved ranges may not
	// overlap.
	ReservedRange []EnumReservedRangeWeaviateModel `json:"reservedRange"`

	// Reserved enum value names, which may not be reused. A given name may only
	// be reserved once.
	ReservedName []string `json:"reservedName"`
}

func (s EnumDescriptorProtoWeaviateModel) ToProto() *EnumDescriptorProto {
	theProto := &EnumDescriptorProto{}

	theProto.Name = s.Name

	for _, protoField := range s.Value {
		msg := protoField.ToProto()
		if theProto.Value == nil {
			theProto.Value = []*EnumValueDescriptorProto{msg}
		} else {
			theProto.Value = append(theProto.Value, msg)
		}
	}

	theProto.Options = s.Options.ToProto()

	for _, protoField := range s.ReservedRange {
		msg := protoField.ToProto()
		if theProto.ReservedRange == nil {
			theProto.ReservedRange = []*EnumReservedRange{msg}
		} else {
			theProto.ReservedRange = append(theProto.ReservedRange, msg)
		}
	}

	theProto.ReservedName = s.ReservedName

	return theProto
}

func (s *EnumDescriptorProto) ToWeaviateModel() EnumDescriptorProtoWeaviateModel {
	model := EnumDescriptorProtoWeaviateModel{}

	model.Name = s.Name

	for _, protoField := range s.Value {
		msg := protoField.ToWeaviateModel()
		if model.Value == nil {
			model.Value = []EnumValueDescriptorProtoWeaviateModel{msg}
		} else {
			model.Value = append(model.Value, msg)
		}
	}

	if s.Options != nil {
		model.Options = lo.ToPtr(s.Options.ToWeaviateModel())
	}

	for _, protoField := range s.ReservedRange {
		msg := protoField.ToWeaviateModel()
		if model.ReservedRange == nil {
			model.ReservedRange = []EnumReservedRangeWeaviateModel{msg}
		} else {
			model.ReservedRange = append(model.ReservedRange, msg)
		}
	}

	model.ReservedName = s.ReservedName

	return model
}

func (s EnumDescriptorProtoWeaviateModel) WeaviateClassName() string {
	return "EnumDescriptorProto"
}

func (s EnumDescriptorProtoWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s EnumDescriptorProtoWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "name",
		DataType: []string{"text"},
	}, {
		Name:     "value",
		DataType: []string{"EnumValueDescriptorProto"},
	}, {
		Name:     "options",
		DataType: []string{"EnumOptions"},
	}, {
		Name:     "reservedRange",
		DataType: []string{"EnumDescriptorProto_EnumReservedRange"},
	}, {
		Name:     "reservedName",
		DataType: []string{"text[]"},
	},
	}
}

func (s EnumDescriptorProtoWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"name": s.Name,

		"value": []map[string]string{},

		"options": []map[string]string{},

		"reservedRange": []map[string]string{},

		"reservedName": s.ReservedName,
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s EnumDescriptorProtoWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	for _, crossReference := range s.Value {
		ValueReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["value"] = append(data["value"].([]map[string]string), ValueReference)
	}
	if s.Options != nil {
		if s.Options.Id != "" {
			OptionsReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", s.Options.Id)}
			data["options"] = append(data["options"].([]map[string]string), OptionsReference)
		}
	}
	for _, crossReference := range s.ReservedRange {
		ReservedRangeReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["reservedRange"] = append(data["reservedRange"].([]map[string]string), ReservedRangeReference)
	}
	return data
}

func (s EnumDescriptorProtoWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s EnumDescriptorProtoWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s EnumDescriptorProtoWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s EnumDescriptorProtoWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type EnumValueDescriptorProtoWeaviateModel struct {
	Name *string `json:"name"`

	Number *int32 `json:"number"`

	Options *EnumValueOptionsWeaviateModel `json:"options"`
}

func (s EnumValueDescriptorProtoWeaviateModel) ToProto() *EnumValueDescriptorProto {
	theProto := &EnumValueDescriptorProto{}

	theProto.Name = s.Name

	theProto.Number = s.Number

	theProto.Options = s.Options.ToProto()

	return theProto
}

func (s *EnumValueDescriptorProto) ToWeaviateModel() EnumValueDescriptorProtoWeaviateModel {
	model := EnumValueDescriptorProtoWeaviateModel{}

	model.Name = s.Name

	model.Number = s.Number

	if s.Options != nil {
		model.Options = lo.ToPtr(s.Options.ToWeaviateModel())
	}

	return model
}

func (s EnumValueDescriptorProtoWeaviateModel) WeaviateClassName() string {
	return "EnumValueDescriptorProto"
}

func (s EnumValueDescriptorProtoWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s EnumValueDescriptorProtoWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "name",
		DataType: []string{"text"},
	}, {
		Name:     "number",
		DataType: []string{"int"},
	}, {
		Name:     "options",
		DataType: []string{"EnumValueOptions"},
	},
	}
}

func (s EnumValueDescriptorProtoWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"name": s.Name,

		"number": s.Number,

		"options": []map[string]string{},
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s EnumValueDescriptorProtoWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	if s.Options != nil {
		if s.Options.Id != "" {
			OptionsReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", s.Options.Id)}
			data["options"] = append(data["options"].([]map[string]string), OptionsReference)
		}
	}
	return data
}

func (s EnumValueDescriptorProtoWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s EnumValueDescriptorProtoWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s EnumValueDescriptorProtoWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s EnumValueDescriptorProtoWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type ServiceDescriptorProtoWeaviateModel struct {
	Name *string `json:"name"`

	Method []MethodDescriptorProtoWeaviateModel `json:"method"`

	Options *ServiceOptionsWeaviateModel `json:"options"`
}

func (s ServiceDescriptorProtoWeaviateModel) ToProto() *ServiceDescriptorProto {
	theProto := &ServiceDescriptorProto{}

	theProto.Name = s.Name

	for _, protoField := range s.Method {
		msg := protoField.ToProto()
		if theProto.Method == nil {
			theProto.Method = []*MethodDescriptorProto{msg}
		} else {
			theProto.Method = append(theProto.Method, msg)
		}
	}

	theProto.Options = s.Options.ToProto()

	return theProto
}

func (s *ServiceDescriptorProto) ToWeaviateModel() ServiceDescriptorProtoWeaviateModel {
	model := ServiceDescriptorProtoWeaviateModel{}

	model.Name = s.Name

	for _, protoField := range s.Method {
		msg := protoField.ToWeaviateModel()
		if model.Method == nil {
			model.Method = []MethodDescriptorProtoWeaviateModel{msg}
		} else {
			model.Method = append(model.Method, msg)
		}
	}

	if s.Options != nil {
		model.Options = lo.ToPtr(s.Options.ToWeaviateModel())
	}

	return model
}

func (s ServiceDescriptorProtoWeaviateModel) WeaviateClassName() string {
	return "ServiceDescriptorProto"
}

func (s ServiceDescriptorProtoWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s ServiceDescriptorProtoWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "name",
		DataType: []string{"text"},
	}, {
		Name:     "method",
		DataType: []string{"MethodDescriptorProto"},
	}, {
		Name:     "options",
		DataType: []string{"ServiceOptions"},
	},
	}
}

func (s ServiceDescriptorProtoWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"name": s.Name,

		"method": []map[string]string{},

		"options": []map[string]string{},
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s ServiceDescriptorProtoWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	for _, crossReference := range s.Method {
		MethodReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["method"] = append(data["method"].([]map[string]string), MethodReference)
	}
	if s.Options != nil {
		if s.Options.Id != "" {
			OptionsReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", s.Options.Id)}
			data["options"] = append(data["options"].([]map[string]string), OptionsReference)
		}
	}
	return data
}

func (s ServiceDescriptorProtoWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s ServiceDescriptorProtoWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s ServiceDescriptorProtoWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s ServiceDescriptorProtoWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type MethodDescriptorProtoWeaviateModel struct {
	Name *string `json:"name"`

	// Input and output type names.  These are resolved in the same way as
	// FieldDescriptorProto.type_name, but must refer to a message type.
	InputType *string `json:"inputType"`

	OutputType *string `json:"outputType"`

	Options *MethodOptionsWeaviateModel `json:"options"`

	// Identifies if client streams multiple client messages
	ClientStreaming *bool `json:"clientStreaming"`

	// Identifies if server streams multiple server messages
	ServerStreaming *bool `json:"serverStreaming"`
}

func (s MethodDescriptorProtoWeaviateModel) ToProto() *MethodDescriptorProto {
	theProto := &MethodDescriptorProto{}

	theProto.Name = s.Name

	theProto.InputType = s.InputType

	theProto.OutputType = s.OutputType

	theProto.Options = s.Options.ToProto()

	theProto.ClientStreaming = s.ClientStreaming

	theProto.ServerStreaming = s.ServerStreaming

	return theProto
}

func (s *MethodDescriptorProto) ToWeaviateModel() MethodDescriptorProtoWeaviateModel {
	model := MethodDescriptorProtoWeaviateModel{}

	model.Name = s.Name

	model.InputType = s.InputType

	model.OutputType = s.OutputType

	if s.Options != nil {
		model.Options = lo.ToPtr(s.Options.ToWeaviateModel())
	}

	model.ClientStreaming = s.ClientStreaming

	model.ServerStreaming = s.ServerStreaming

	return model
}

func (s MethodDescriptorProtoWeaviateModel) WeaviateClassName() string {
	return "MethodDescriptorProto"
}

func (s MethodDescriptorProtoWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s MethodDescriptorProtoWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "name",
		DataType: []string{"text"},
	}, {
		Name:     "inputType",
		DataType: []string{"text"},
	}, {
		Name:     "outputType",
		DataType: []string{"text"},
	}, {
		Name:     "options",
		DataType: []string{"MethodOptions"},
	}, {
		Name:     "clientStreaming",
		DataType: []string{"boolean"},
	}, {
		Name:     "serverStreaming",
		DataType: []string{"boolean"},
	},
	}
}

func (s MethodDescriptorProtoWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"name": s.Name,

		"inputType": s.InputType,

		"outputType": s.OutputType,

		"options": []map[string]string{},

		"clientStreaming": s.ClientStreaming,

		"serverStreaming": s.ServerStreaming,
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s MethodDescriptorProtoWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	if s.Options != nil {
		if s.Options.Id != "" {
			OptionsReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", s.Options.Id)}
			data["options"] = append(data["options"].([]map[string]string), OptionsReference)
		}
	}
	return data
}

func (s MethodDescriptorProtoWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s MethodDescriptorProtoWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s MethodDescriptorProtoWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s MethodDescriptorProtoWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type FileOptionsWeaviateModel struct {

	// Sets the Java package where classes generated from this .proto will be
	// placed.  By default, the proto package is used, but this is often
	// inappropriate because proto packages do not normally start with backwards
	// domain names.
	JavaPackage *string `json:"javaPackage"`

	// Controls the name of the wrapper Java class generated for the .proto file.
	// That class will always contain the .proto file's getDescriptor() method as
	// well as any top-level extensions defined in the .proto file.
	// If java_multiple_files is disabled, then all the other classes from the
	// .proto file will be nested inside the single wrapper outer class.
	JavaOuterClassname *string `json:"javaOuterClassname"`

	// If enabled, then the Java code generator will generate a separate .java
	// file for each top-level message, enum, and service defined in the .proto
	// file.  Thus, these types will *not* be nested inside the wrapper class
	// named by java_outer_classname.  However, the wrapper class will still be
	// generated to contain the file's getDescriptor() method as well as any
	// top-level extensions defined in the file.
	JavaMultipleFiles *bool `json:"javaMultipleFiles"`

	// This option does nothing.
	JavaGenerateEqualsAndHash *bool `json:"javaGenerateEqualsAndHash"`

	// If set true, then the Java2 code generator will generate code that
	// throws an exception whenever an attempt is made to assign a non-UTF-8
	// byte sequence to a string field.
	// Message reflection will do the same.
	// However, an extension field still accepts non-UTF-8 byte sequences.
	// This option has no effect on when used with the lite runtime.
	JavaStringCheckUtf8 *bool `json:"javaStringCheckUtf8"`

	OptimizeFor *int `json:"optimizeFor"`

	// Sets the Go package where structs generated from this .proto will be
	// placed. If omitted, the Go package will be derived from the following:
	//   - The basename of the package import path, if provided.
	//   - Otherwise, the package statement in the .proto file, if present.
	//   - Otherwise, the basename of the .proto file, without extension.
	GoPackage *string `json:"goPackage"`

	// Should generic services be generated in each language?  "Generic" services
	// are not specific to any particular RPC system.  They are generated by the
	// main code generators in each language (without additional plugins).
	// Generic services were the only kind of service generation supported by
	// early versions of google.protobuf.
	//
	// Generic services are now considered deprecated in favor of using plugins
	// that generate code specific to your particular RPC system.  Therefore,
	// these default to false.  Old code which depends on generic services should
	// explicitly set them to true.
	CcGenericServices *bool `json:"ccGenericServices"`

	JavaGenericServices *bool `json:"javaGenericServices"`

	PyGenericServices *bool `json:"pyGenericServices"`

	PhpGenericServices *bool `json:"phpGenericServices"`

	// Is this file deprecated?
	// Depending on the target platform, this can emit Deprecated annotations
	// for everything in the file, or it will be completely ignored; in the very
	// least, this is a formalization for deprecating files.
	Deprecated *bool `json:"deprecated"`

	// Enables the use of arenas for the proto messages in this file. This applies
	// only to generated classes for C++.
	CcEnableArenas *bool `json:"ccEnableArenas"`

	// Sets the objective c class prefix which is prepended to all objective c
	// generated classes from this .proto. There is no default.
	ObjcClassPrefix *string `json:"objcClassPrefix"`

	// Namespace for generated classes; defaults to the package.
	CsharpNamespace *string `json:"csharpNamespace"`

	// By default Swift generators will take the proto package and CamelCase it
	// replacing '.' with underscore and use that to prefix the types/symbols
	// defined. When this options is provided, they will use this value instead
	// to prefix the types/symbols defined.
	SwiftPrefix *string `json:"swiftPrefix"`

	// Sets the php class prefix which is prepended to all php generated classes
	// from this .proto. Default is empty.
	PhpClassPrefix *string `json:"phpClassPrefix"`

	// Use this option to change the namespace of php generated classes. Default
	// is empty. When this option is empty, the package name will be used for
	// determining the namespace.
	PhpNamespace *string `json:"phpNamespace"`

	// Use this option to change the namespace of php generated metadata classes.
	// Default is empty. When this option is empty, the proto file name will be
	// used for determining the namespace.
	PhpMetadataNamespace *string `json:"phpMetadataNamespace"`

	// Use this option to change the package of ruby generated classes. Default
	// is empty. When this option is not set, the package name will be used for
	// determining the ruby package.
	RubyPackage *string `json:"rubyPackage"`

	// The parser stores options it doesn't recognize here.
	// See the documentation for the "Options" section above.
	UninterpretedOption []UninterpretedOptionWeaviateModel `json:"uninterpretedOption"`
}

func (s FileOptionsWeaviateModel) ToProto() *FileOptions {
	theProto := &FileOptions{}

	theProto.JavaPackage = s.JavaPackage

	theProto.JavaOuterClassname = s.JavaOuterClassname

	theProto.JavaMultipleFiles = s.JavaMultipleFiles

	theProto.JavaGenerateEqualsAndHash = s.JavaGenerateEqualsAndHash

	theProto.JavaStringCheckUtf8 = s.JavaStringCheckUtf8

	theProto.OptimizeFor = s.OptimizeFor

	theProto.GoPackage = s.GoPackage

	theProto.CcGenericServices = s.CcGenericServices

	theProto.JavaGenericServices = s.JavaGenericServices

	theProto.PyGenericServices = s.PyGenericServices

	theProto.PhpGenericServices = s.PhpGenericServices

	theProto.Deprecated = s.Deprecated

	theProto.CcEnableArenas = s.CcEnableArenas

	theProto.ObjcClassPrefix = s.ObjcClassPrefix

	theProto.CsharpNamespace = s.CsharpNamespace

	theProto.SwiftPrefix = s.SwiftPrefix

	theProto.PhpClassPrefix = s.PhpClassPrefix

	theProto.PhpNamespace = s.PhpNamespace

	theProto.PhpMetadataNamespace = s.PhpMetadataNamespace

	theProto.RubyPackage = s.RubyPackage

	for _, protoField := range s.UninterpretedOption {
		msg := protoField.ToProto()
		if theProto.UninterpretedOption == nil {
			theProto.UninterpretedOption = []*UninterpretedOption{msg}
		} else {
			theProto.UninterpretedOption = append(theProto.UninterpretedOption, msg)
		}
	}

	return theProto
}

func (s *FileOptions) ToWeaviateModel() FileOptionsWeaviateModel {
	model := FileOptionsWeaviateModel{}

	model.JavaPackage = s.JavaPackage

	model.JavaOuterClassname = s.JavaOuterClassname

	model.JavaMultipleFiles = s.JavaMultipleFiles

	model.JavaGenerateEqualsAndHash = s.JavaGenerateEqualsAndHash

	model.JavaStringCheckUtf8 = s.JavaStringCheckUtf8

	model.OptimizeFor = s.OptimizeFor

	model.GoPackage = s.GoPackage

	model.CcGenericServices = s.CcGenericServices

	model.JavaGenericServices = s.JavaGenericServices

	model.PyGenericServices = s.PyGenericServices

	model.PhpGenericServices = s.PhpGenericServices

	model.Deprecated = s.Deprecated

	model.CcEnableArenas = s.CcEnableArenas

	model.ObjcClassPrefix = s.ObjcClassPrefix

	model.CsharpNamespace = s.CsharpNamespace

	model.SwiftPrefix = s.SwiftPrefix

	model.PhpClassPrefix = s.PhpClassPrefix

	model.PhpNamespace = s.PhpNamespace

	model.PhpMetadataNamespace = s.PhpMetadataNamespace

	model.RubyPackage = s.RubyPackage

	for _, protoField := range s.UninterpretedOption {
		msg := protoField.ToWeaviateModel()
		if model.UninterpretedOption == nil {
			model.UninterpretedOption = []UninterpretedOptionWeaviateModel{msg}
		} else {
			model.UninterpretedOption = append(model.UninterpretedOption, msg)
		}
	}

	return model
}

func (s FileOptionsWeaviateModel) WeaviateClassName() string {
	return "FileOptions"
}

func (s FileOptionsWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s FileOptionsWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "javaPackage",
		DataType: []string{"text"},
	}, {
		Name:     "javaOuterClassname",
		DataType: []string{"text"},
	}, {
		Name:     "javaMultipleFiles",
		DataType: []string{"boolean"},
	}, {
		Name:     "javaGenerateEqualsAndHash",
		DataType: []string{"boolean"},
	}, {
		Name:     "javaStringCheckUtf8",
		DataType: []string{"boolean"},
	}, {
		Name:     "optimizeFor",
		DataType: []string{"int"},
	}, {
		Name:     "goPackage",
		DataType: []string{"text"},
	}, {
		Name:     "ccGenericServices",
		DataType: []string{"boolean"},
	}, {
		Name:     "javaGenericServices",
		DataType: []string{"boolean"},
	}, {
		Name:     "pyGenericServices",
		DataType: []string{"boolean"},
	}, {
		Name:     "phpGenericServices",
		DataType: []string{"boolean"},
	}, {
		Name:     "deprecated",
		DataType: []string{"boolean"},
	}, {
		Name:     "ccEnableArenas",
		DataType: []string{"boolean"},
	}, {
		Name:     "objcClassPrefix",
		DataType: []string{"text"},
	}, {
		Name:     "csharpNamespace",
		DataType: []string{"text"},
	}, {
		Name:     "swiftPrefix",
		DataType: []string{"text"},
	}, {
		Name:     "phpClassPrefix",
		DataType: []string{"text"},
	}, {
		Name:     "phpNamespace",
		DataType: []string{"text"},
	}, {
		Name:     "phpMetadataNamespace",
		DataType: []string{"text"},
	}, {
		Name:     "rubyPackage",
		DataType: []string{"text"},
	}, {
		Name:     "uninterpretedOption",
		DataType: []string{"UninterpretedOption"},
	},
	}
}

func (s FileOptionsWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"javaPackage": s.JavaPackage,

		"javaOuterClassname": s.JavaOuterClassname,

		"javaMultipleFiles": s.JavaMultipleFiles,

		"javaGenerateEqualsAndHash": s.JavaGenerateEqualsAndHash,

		"javaStringCheckUtf8": s.JavaStringCheckUtf8,

		"optimizeFor": s.OptimizeFor,

		"goPackage": s.GoPackage,

		"ccGenericServices": s.CcGenericServices,

		"javaGenericServices": s.JavaGenericServices,

		"pyGenericServices": s.PyGenericServices,

		"phpGenericServices": s.PhpGenericServices,

		"deprecated": s.Deprecated,

		"ccEnableArenas": s.CcEnableArenas,

		"objcClassPrefix": s.ObjcClassPrefix,

		"csharpNamespace": s.CsharpNamespace,

		"swiftPrefix": s.SwiftPrefix,

		"phpClassPrefix": s.PhpClassPrefix,

		"phpNamespace": s.PhpNamespace,

		"phpMetadataNamespace": s.PhpMetadataNamespace,

		"rubyPackage": s.RubyPackage,

		"uninterpretedOption": []map[string]string{},
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s FileOptionsWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	for _, crossReference := range s.UninterpretedOption {
		UninterpretedOptionReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["uninterpretedOption"] = append(data["uninterpretedOption"].([]map[string]string), UninterpretedOptionReference)
	}
	return data
}

func (s FileOptionsWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s FileOptionsWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s FileOptionsWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s FileOptionsWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type MessageOptionsWeaviateModel struct {

	// Set true to use the old proto1 MessageSet wire format for extensions.
	// This is provided for backwards-compatibility with the MessageSet wire
	// format.  You should not use this for any other reason:  It's less
	// efficient, has fewer features, and is more complicated.
	//
	// The message must be defined exactly as follows:
	//   message Foo {
	//     option message_set_wire_format = true;
	//     extensions 4 to max;
	//   }
	// Note that the message cannot have any defined fields; MessageSets only
	// have extensions.
	//
	// All extensions of your type must be singular messages; e.g. they cannot
	// be int32s, enums, or repeated messages.
	//
	// Because this is an option, the above two restrictions are not enforced by
	// the protocol compiler.
	MessageSetWireFormat *bool `json:"messageSetWireFormat"`

	// Disables the generation of the standard "descriptor()" accessor, which can
	// conflict with a field of the same name.  This is meant to make migration
	// from proto1 easier; new code should avoid fields named "descriptor".
	NoStandardDescriptorAccessor *bool `json:"noStandardDescriptorAccessor"`

	// Is this message deprecated?
	// Depending on the target platform, this can emit Deprecated annotations
	// for the message, or it will be completely ignored; in the very least,
	// this is a formalization for deprecating messages.
	Deprecated *bool `json:"deprecated"`

	// NOTE: Do not set the option in .proto files. Always use the maps syntax
	// instead. The option should only be implicitly set by the proto compiler
	// parser.
	//
	// Whether the message is an automatically generated map entry type for the
	// maps field.
	//
	// For maps fields:
	//     map<KeyType, ValueType> map_field = 1;
	// The parsed descriptor looks like:
	//     message MapFieldEntry {
	//         option map_entry = true;
	//         optional KeyType key = 1;
	//         optional ValueType value = 2;
	//     }
	//     repeated MapFieldEntry map_field = 1;
	//
	// Implementations may choose not to generate the map_entry=true message, but
	// use a native map in the target language to hold the keys and values.
	// The reflection APIs in such implementations still need to work as
	// if the field is a repeated message field.
	MapEntry *bool `json:"mapEntry"`

	// Enable the legacy handling of JSON field name conflicts.  This lowercases
	// and strips underscored from the fields before comparison in proto3 only.
	// The new behavior takes `json_name` into account and applies to proto2 as
	// well.
	//
	// This should only be used as a temporary measure against broken builds due
	// to the change in behavior for JSON field name conflicts.
	//
	// TODO(b/261750190) This is legacy behavior we plan to remove once downstream
	// teams have had time to migrate.
	DeprecatedLegacyJsonFieldConflicts *bool `json:"deprecatedLegacyJsonFieldConflicts"`

	// The parser stores options it doesn't recognize here. See above.
	UninterpretedOption []UninterpretedOptionWeaviateModel `json:"uninterpretedOption"`
}

func (s MessageOptionsWeaviateModel) ToProto() *MessageOptions {
	theProto := &MessageOptions{}

	theProto.MessageSetWireFormat = s.MessageSetWireFormat

	theProto.NoStandardDescriptorAccessor = s.NoStandardDescriptorAccessor

	theProto.Deprecated = s.Deprecated

	theProto.MapEntry = s.MapEntry

	theProto.DeprecatedLegacyJsonFieldConflicts = s.DeprecatedLegacyJsonFieldConflicts

	for _, protoField := range s.UninterpretedOption {
		msg := protoField.ToProto()
		if theProto.UninterpretedOption == nil {
			theProto.UninterpretedOption = []*UninterpretedOption{msg}
		} else {
			theProto.UninterpretedOption = append(theProto.UninterpretedOption, msg)
		}
	}

	return theProto
}

func (s *MessageOptions) ToWeaviateModel() MessageOptionsWeaviateModel {
	model := MessageOptionsWeaviateModel{}

	model.MessageSetWireFormat = s.MessageSetWireFormat

	model.NoStandardDescriptorAccessor = s.NoStandardDescriptorAccessor

	model.Deprecated = s.Deprecated

	model.MapEntry = s.MapEntry

	model.DeprecatedLegacyJsonFieldConflicts = s.DeprecatedLegacyJsonFieldConflicts

	for _, protoField := range s.UninterpretedOption {
		msg := protoField.ToWeaviateModel()
		if model.UninterpretedOption == nil {
			model.UninterpretedOption = []UninterpretedOptionWeaviateModel{msg}
		} else {
			model.UninterpretedOption = append(model.UninterpretedOption, msg)
		}
	}

	return model
}

func (s MessageOptionsWeaviateModel) WeaviateClassName() string {
	return "MessageOptions"
}

func (s MessageOptionsWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s MessageOptionsWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "messageSetWireFormat",
		DataType: []string{"boolean"},
	}, {
		Name:     "noStandardDescriptorAccessor",
		DataType: []string{"boolean"},
	}, {
		Name:     "deprecated",
		DataType: []string{"boolean"},
	}, {
		Name:     "mapEntry",
		DataType: []string{"boolean"},
	}, {
		Name:     "deprecatedLegacyJsonFieldConflicts",
		DataType: []string{"boolean"},
	}, {
		Name:     "uninterpretedOption",
		DataType: []string{"UninterpretedOption"},
	},
	}
}

func (s MessageOptionsWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"messageSetWireFormat": s.MessageSetWireFormat,

		"noStandardDescriptorAccessor": s.NoStandardDescriptorAccessor,

		"deprecated": s.Deprecated,

		"mapEntry": s.MapEntry,

		"deprecatedLegacyJsonFieldConflicts": s.DeprecatedLegacyJsonFieldConflicts,

		"uninterpretedOption": []map[string]string{},
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s MessageOptionsWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	for _, crossReference := range s.UninterpretedOption {
		UninterpretedOptionReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["uninterpretedOption"] = append(data["uninterpretedOption"].([]map[string]string), UninterpretedOptionReference)
	}
	return data
}

func (s MessageOptionsWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s MessageOptionsWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s MessageOptionsWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s MessageOptionsWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type FieldOptionsWeaviateModel struct {

	// The ctype option instructs the C++ code generator to use a different
	// representation of the field than it normally would.  See the specific
	// options below.  This option is not yet implemented in the open source
	// release -- sorry, we'll try to include it in a future version!
	Ctype *int `json:"ctype"`

	// The packed option can be enabled for repeated primitive fields to enable
	// a more efficient representation on the wire. Rather than repeatedly
	// writing the tag and type for each element, the entire array is encoded as
	// a single length-delimited blob. In proto3, only explicit setting it to
	// false will avoid using packed encoding.
	Packed *bool `json:"packed"`

	// The jstype option determines the JavaScript type used for values of the
	// field.  The option is permitted only for 64 bit integral and fixed types
	// (int64, uint64, sint64, fixed64, sfixed64).  A field with jstype JS_STRING
	// is represented as JavaScript string, which avoids loss of precision that
	// can happen when a large value is converted to a floating point JavaScript.
	// Specifying JS_NUMBER for the jstype causes the generated JavaScript code to
	// use the JavaScript "number" type.  The behavior of the default option
	// JS_NORMAL is implementation dependent.
	//
	// This option is an enum to permit additional types to be added, e.g.
	// goog.math.Integer.
	Jstype *int `json:"jstype"`

	// Should this field be parsed lazily?  Lazy applies only to message-type
	// fields.  It means that when the outer message is initially parsed, the
	// inner message's contents will not be parsed but instead stored in encoded
	// form.  The inner message will actually be parsed when it is first accessed.
	//
	// This is only a hint.  Implementations are free to choose whether to use
	// eager or lazy parsing regardless of the value of this option.  However,
	// setting this option true suggests that the protocol author believes that
	// using lazy parsing on this field is worth the additional bookkeeping
	// overhead typically needed to implement it.
	//
	// This option does not affect the public interface of any generated code;
	// all method signatures remain the same.  Furthermore, thread-safety of the
	// interface is not affected by this option; const methods remain safe to
	// call from multiple threads concurrently, while non-const methods continue
	// to require exclusive access.
	//
	// Note that implementations may choose not to check required fields within
	// a lazy sub-message.  That is, calling IsInitialized() on the outer message
	// may return true even if the inner message has missing required fields.
	// This is necessary because otherwise the inner message would have to be
	// parsed in order to perform the check, defeating the purpose of lazy
	// parsing.  An implementation which chooses not to check required fields
	// must be consistent about it.  That is, for any particular sub-message, the
	// implementation must either *always* check its required fields, or *never*
	// check its required fields, regardless of whether or not the message has
	// been parsed.
	//
	// As of May 2022, lazy verifies the contents of the byte stream during
	// parsing.  An invalid byte stream will cause the overall parsing to fail.
	Lazy *bool `json:"lazy"`

	// unverified_lazy does no correctness checks on the byte stream. This should
	// only be used where lazy with verification is prohibitive for performance
	// reasons.
	UnverifiedLazy *bool `json:"unverifiedLazy"`

	// Is this field deprecated?
	// Depending on the target platform, this can emit Deprecated annotations
	// for accessors, or it will be completely ignored; in the very least, this
	// is a formalization for deprecating fields.
	Deprecated *bool `json:"deprecated"`

	// For Google-internal migration only. Do not use.
	Weak *bool `json:"weak"`

	// Indicate that the field value should not be printed out when using debug
	// formats, e.g. when the field contains sensitive credentials.
	DebugRedact *bool `json:"debugRedact"`

	Retention *int `json:"retention"`

	Target *int `json:"target"`

	// The parser stores options it doesn't recognize here. See above.
	UninterpretedOption []UninterpretedOptionWeaviateModel `json:"uninterpretedOption"`
}

func (s FieldOptionsWeaviateModel) ToProto() *FieldOptions {
	theProto := &FieldOptions{}

	theProto.Ctype = s.Ctype

	theProto.Packed = s.Packed

	theProto.Jstype = s.Jstype

	theProto.Lazy = s.Lazy

	theProto.UnverifiedLazy = s.UnverifiedLazy

	theProto.Deprecated = s.Deprecated

	theProto.Weak = s.Weak

	theProto.DebugRedact = s.DebugRedact

	theProto.Retention = s.Retention

	theProto.Target = s.Target

	for _, protoField := range s.UninterpretedOption {
		msg := protoField.ToProto()
		if theProto.UninterpretedOption == nil {
			theProto.UninterpretedOption = []*UninterpretedOption{msg}
		} else {
			theProto.UninterpretedOption = append(theProto.UninterpretedOption, msg)
		}
	}

	return theProto
}

func (s *FieldOptions) ToWeaviateModel() FieldOptionsWeaviateModel {
	model := FieldOptionsWeaviateModel{}

	model.Ctype = s.Ctype

	model.Packed = s.Packed

	model.Jstype = s.Jstype

	model.Lazy = s.Lazy

	model.UnverifiedLazy = s.UnverifiedLazy

	model.Deprecated = s.Deprecated

	model.Weak = s.Weak

	model.DebugRedact = s.DebugRedact

	model.Retention = s.Retention

	model.Target = s.Target

	for _, protoField := range s.UninterpretedOption {
		msg := protoField.ToWeaviateModel()
		if model.UninterpretedOption == nil {
			model.UninterpretedOption = []UninterpretedOptionWeaviateModel{msg}
		} else {
			model.UninterpretedOption = append(model.UninterpretedOption, msg)
		}
	}

	return model
}

func (s FieldOptionsWeaviateModel) WeaviateClassName() string {
	return "FieldOptions"
}

func (s FieldOptionsWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s FieldOptionsWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "ctype",
		DataType: []string{"int"},
	}, {
		Name:     "packed",
		DataType: []string{"boolean"},
	}, {
		Name:     "jstype",
		DataType: []string{"int"},
	}, {
		Name:     "lazy",
		DataType: []string{"boolean"},
	}, {
		Name:     "unverifiedLazy",
		DataType: []string{"boolean"},
	}, {
		Name:     "deprecated",
		DataType: []string{"boolean"},
	}, {
		Name:     "weak",
		DataType: []string{"boolean"},
	}, {
		Name:     "debugRedact",
		DataType: []string{"boolean"},
	}, {
		Name:     "retention",
		DataType: []string{"int"},
	}, {
		Name:     "target",
		DataType: []string{"int"},
	}, {
		Name:     "uninterpretedOption",
		DataType: []string{"UninterpretedOption"},
	},
	}
}

func (s FieldOptionsWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"ctype": s.Ctype,

		"packed": s.Packed,

		"jstype": s.Jstype,

		"lazy": s.Lazy,

		"unverifiedLazy": s.UnverifiedLazy,

		"deprecated": s.Deprecated,

		"weak": s.Weak,

		"debugRedact": s.DebugRedact,

		"retention": s.Retention,

		"target": s.Target,

		"uninterpretedOption": []map[string]string{},
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s FieldOptionsWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	for _, crossReference := range s.UninterpretedOption {
		UninterpretedOptionReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["uninterpretedOption"] = append(data["uninterpretedOption"].([]map[string]string), UninterpretedOptionReference)
	}
	return data
}

func (s FieldOptionsWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s FieldOptionsWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s FieldOptionsWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s FieldOptionsWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type OneofOptionsWeaviateModel struct {

	// The parser stores options it doesn't recognize here. See above.
	UninterpretedOption []UninterpretedOptionWeaviateModel `json:"uninterpretedOption"`
}

func (s OneofOptionsWeaviateModel) ToProto() *OneofOptions {
	theProto := &OneofOptions{}

	for _, protoField := range s.UninterpretedOption {
		msg := protoField.ToProto()
		if theProto.UninterpretedOption == nil {
			theProto.UninterpretedOption = []*UninterpretedOption{msg}
		} else {
			theProto.UninterpretedOption = append(theProto.UninterpretedOption, msg)
		}
	}

	return theProto
}

func (s *OneofOptions) ToWeaviateModel() OneofOptionsWeaviateModel {
	model := OneofOptionsWeaviateModel{}

	for _, protoField := range s.UninterpretedOption {
		msg := protoField.ToWeaviateModel()
		if model.UninterpretedOption == nil {
			model.UninterpretedOption = []UninterpretedOptionWeaviateModel{msg}
		} else {
			model.UninterpretedOption = append(model.UninterpretedOption, msg)
		}
	}

	return model
}

func (s OneofOptionsWeaviateModel) WeaviateClassName() string {
	return "OneofOptions"
}

func (s OneofOptionsWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s OneofOptionsWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "uninterpretedOption",
		DataType: []string{"UninterpretedOption"},
	},
	}
}

func (s OneofOptionsWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"uninterpretedOption": []map[string]string{},
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s OneofOptionsWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	for _, crossReference := range s.UninterpretedOption {
		UninterpretedOptionReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["uninterpretedOption"] = append(data["uninterpretedOption"].([]map[string]string), UninterpretedOptionReference)
	}
	return data
}

func (s OneofOptionsWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s OneofOptionsWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s OneofOptionsWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s OneofOptionsWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type EnumOptionsWeaviateModel struct {

	// Set this option to true to allow mapping different tag names to the same
	// value.
	AllowAlias *bool `json:"allowAlias"`

	// Is this enum deprecated?
	// Depending on the target platform, this can emit Deprecated annotations
	// for the enum, or it will be completely ignored; in the very least, this
	// is a formalization for deprecating enums.
	Deprecated *bool `json:"deprecated"`

	// Enable the legacy handling of JSON field name conflicts.  This lowercases
	// and strips underscored from the fields before comparison in proto3 only.
	// The new behavior takes `json_name` into account and applies to proto2 as
	// well.
	// TODO(b/261750190) Remove this legacy behavior once downstream teams have
	// had time to migrate.
	DeprecatedLegacyJsonFieldConflicts *bool `json:"deprecatedLegacyJsonFieldConflicts"`

	// The parser stores options it doesn't recognize here. See above.
	UninterpretedOption []UninterpretedOptionWeaviateModel `json:"uninterpretedOption"`
}

func (s EnumOptionsWeaviateModel) ToProto() *EnumOptions {
	theProto := &EnumOptions{}

	theProto.AllowAlias = s.AllowAlias

	theProto.Deprecated = s.Deprecated

	theProto.DeprecatedLegacyJsonFieldConflicts = s.DeprecatedLegacyJsonFieldConflicts

	for _, protoField := range s.UninterpretedOption {
		msg := protoField.ToProto()
		if theProto.UninterpretedOption == nil {
			theProto.UninterpretedOption = []*UninterpretedOption{msg}
		} else {
			theProto.UninterpretedOption = append(theProto.UninterpretedOption, msg)
		}
	}

	return theProto
}

func (s *EnumOptions) ToWeaviateModel() EnumOptionsWeaviateModel {
	model := EnumOptionsWeaviateModel{}

	model.AllowAlias = s.AllowAlias

	model.Deprecated = s.Deprecated

	model.DeprecatedLegacyJsonFieldConflicts = s.DeprecatedLegacyJsonFieldConflicts

	for _, protoField := range s.UninterpretedOption {
		msg := protoField.ToWeaviateModel()
		if model.UninterpretedOption == nil {
			model.UninterpretedOption = []UninterpretedOptionWeaviateModel{msg}
		} else {
			model.UninterpretedOption = append(model.UninterpretedOption, msg)
		}
	}

	return model
}

func (s EnumOptionsWeaviateModel) WeaviateClassName() string {
	return "EnumOptions"
}

func (s EnumOptionsWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s EnumOptionsWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "allowAlias",
		DataType: []string{"boolean"},
	}, {
		Name:     "deprecated",
		DataType: []string{"boolean"},
	}, {
		Name:     "deprecatedLegacyJsonFieldConflicts",
		DataType: []string{"boolean"},
	}, {
		Name:     "uninterpretedOption",
		DataType: []string{"UninterpretedOption"},
	},
	}
}

func (s EnumOptionsWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"allowAlias": s.AllowAlias,

		"deprecated": s.Deprecated,

		"deprecatedLegacyJsonFieldConflicts": s.DeprecatedLegacyJsonFieldConflicts,

		"uninterpretedOption": []map[string]string{},
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s EnumOptionsWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	for _, crossReference := range s.UninterpretedOption {
		UninterpretedOptionReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["uninterpretedOption"] = append(data["uninterpretedOption"].([]map[string]string), UninterpretedOptionReference)
	}
	return data
}

func (s EnumOptionsWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s EnumOptionsWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s EnumOptionsWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s EnumOptionsWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type EnumValueOptionsWeaviateModel struct {

	// Is this enum value deprecated?
	// Depending on the target platform, this can emit Deprecated annotations
	// for the enum value, or it will be completely ignored; in the very least,
	// this is a formalization for deprecating enum values.
	Deprecated *bool `json:"deprecated"`

	// The parser stores options it doesn't recognize here. See above.
	UninterpretedOption []UninterpretedOptionWeaviateModel `json:"uninterpretedOption"`
}

func (s EnumValueOptionsWeaviateModel) ToProto() *EnumValueOptions {
	theProto := &EnumValueOptions{}

	theProto.Deprecated = s.Deprecated

	for _, protoField := range s.UninterpretedOption {
		msg := protoField.ToProto()
		if theProto.UninterpretedOption == nil {
			theProto.UninterpretedOption = []*UninterpretedOption{msg}
		} else {
			theProto.UninterpretedOption = append(theProto.UninterpretedOption, msg)
		}
	}

	return theProto
}

func (s *EnumValueOptions) ToWeaviateModel() EnumValueOptionsWeaviateModel {
	model := EnumValueOptionsWeaviateModel{}

	model.Deprecated = s.Deprecated

	for _, protoField := range s.UninterpretedOption {
		msg := protoField.ToWeaviateModel()
		if model.UninterpretedOption == nil {
			model.UninterpretedOption = []UninterpretedOptionWeaviateModel{msg}
		} else {
			model.UninterpretedOption = append(model.UninterpretedOption, msg)
		}
	}

	return model
}

func (s EnumValueOptionsWeaviateModel) WeaviateClassName() string {
	return "EnumValueOptions"
}

func (s EnumValueOptionsWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s EnumValueOptionsWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "deprecated",
		DataType: []string{"boolean"},
	}, {
		Name:     "uninterpretedOption",
		DataType: []string{"UninterpretedOption"},
	},
	}
}

func (s EnumValueOptionsWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"deprecated": s.Deprecated,

		"uninterpretedOption": []map[string]string{},
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s EnumValueOptionsWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	for _, crossReference := range s.UninterpretedOption {
		UninterpretedOptionReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["uninterpretedOption"] = append(data["uninterpretedOption"].([]map[string]string), UninterpretedOptionReference)
	}
	return data
}

func (s EnumValueOptionsWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s EnumValueOptionsWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s EnumValueOptionsWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s EnumValueOptionsWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type ServiceOptionsWeaviateModel struct {

	// Is this service deprecated?
	// Depending on the target platform, this can emit Deprecated annotations
	// for the service, or it will be completely ignored; in the very least,
	// this is a formalization for deprecating services.
	Deprecated *bool `json:"deprecated"`

	// The parser stores options it doesn't recognize here. See above.
	UninterpretedOption []UninterpretedOptionWeaviateModel `json:"uninterpretedOption"`
}

func (s ServiceOptionsWeaviateModel) ToProto() *ServiceOptions {
	theProto := &ServiceOptions{}

	theProto.Deprecated = s.Deprecated

	for _, protoField := range s.UninterpretedOption {
		msg := protoField.ToProto()
		if theProto.UninterpretedOption == nil {
			theProto.UninterpretedOption = []*UninterpretedOption{msg}
		} else {
			theProto.UninterpretedOption = append(theProto.UninterpretedOption, msg)
		}
	}

	return theProto
}

func (s *ServiceOptions) ToWeaviateModel() ServiceOptionsWeaviateModel {
	model := ServiceOptionsWeaviateModel{}

	model.Deprecated = s.Deprecated

	for _, protoField := range s.UninterpretedOption {
		msg := protoField.ToWeaviateModel()
		if model.UninterpretedOption == nil {
			model.UninterpretedOption = []UninterpretedOptionWeaviateModel{msg}
		} else {
			model.UninterpretedOption = append(model.UninterpretedOption, msg)
		}
	}

	return model
}

func (s ServiceOptionsWeaviateModel) WeaviateClassName() string {
	return "ServiceOptions"
}

func (s ServiceOptionsWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s ServiceOptionsWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "deprecated",
		DataType: []string{"boolean"},
	}, {
		Name:     "uninterpretedOption",
		DataType: []string{"UninterpretedOption"},
	},
	}
}

func (s ServiceOptionsWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"deprecated": s.Deprecated,

		"uninterpretedOption": []map[string]string{},
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s ServiceOptionsWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	for _, crossReference := range s.UninterpretedOption {
		UninterpretedOptionReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["uninterpretedOption"] = append(data["uninterpretedOption"].([]map[string]string), UninterpretedOptionReference)
	}
	return data
}

func (s ServiceOptionsWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s ServiceOptionsWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s ServiceOptionsWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s ServiceOptionsWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type MethodOptionsWeaviateModel struct {

	// Is this method deprecated?
	// Depending on the target platform, this can emit Deprecated annotations
	// for the method, or it will be completely ignored; in the very least,
	// this is a formalization for deprecating methods.
	Deprecated *bool `json:"deprecated"`

	IdempotencyLevel *int `json:"idempotencyLevel"`

	// The parser stores options it doesn't recognize here. See above.
	UninterpretedOption []UninterpretedOptionWeaviateModel `json:"uninterpretedOption"`
}

func (s MethodOptionsWeaviateModel) ToProto() *MethodOptions {
	theProto := &MethodOptions{}

	theProto.Deprecated = s.Deprecated

	theProto.IdempotencyLevel = s.IdempotencyLevel

	for _, protoField := range s.UninterpretedOption {
		msg := protoField.ToProto()
		if theProto.UninterpretedOption == nil {
			theProto.UninterpretedOption = []*UninterpretedOption{msg}
		} else {
			theProto.UninterpretedOption = append(theProto.UninterpretedOption, msg)
		}
	}

	return theProto
}

func (s *MethodOptions) ToWeaviateModel() MethodOptionsWeaviateModel {
	model := MethodOptionsWeaviateModel{}

	model.Deprecated = s.Deprecated

	model.IdempotencyLevel = s.IdempotencyLevel

	for _, protoField := range s.UninterpretedOption {
		msg := protoField.ToWeaviateModel()
		if model.UninterpretedOption == nil {
			model.UninterpretedOption = []UninterpretedOptionWeaviateModel{msg}
		} else {
			model.UninterpretedOption = append(model.UninterpretedOption, msg)
		}
	}

	return model
}

func (s MethodOptionsWeaviateModel) WeaviateClassName() string {
	return "MethodOptions"
}

func (s MethodOptionsWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s MethodOptionsWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "deprecated",
		DataType: []string{"boolean"},
	}, {
		Name:     "idempotencyLevel",
		DataType: []string{"int"},
	}, {
		Name:     "uninterpretedOption",
		DataType: []string{"UninterpretedOption"},
	},
	}
}

func (s MethodOptionsWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"deprecated": s.Deprecated,

		"idempotencyLevel": s.IdempotencyLevel,

		"uninterpretedOption": []map[string]string{},
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s MethodOptionsWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	for _, crossReference := range s.UninterpretedOption {
		UninterpretedOptionReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["uninterpretedOption"] = append(data["uninterpretedOption"].([]map[string]string), UninterpretedOptionReference)
	}
	return data
}

func (s MethodOptionsWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s MethodOptionsWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s MethodOptionsWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s MethodOptionsWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type UninterpretedOptionWeaviateModel struct {
	Name []NamePartWeaviateModel `json:"name"`

	// The value of the uninterpreted option, in whatever type the tokenizer
	// identified it as during parsing. Exactly one of these should be set.
	IdentifierValue *string `json:"identifierValue"`

	PositiveIntValue *uint64 `json:"positiveIntValue,string"`

	NegativeIntValue *int64 `json:"negativeIntValue,string"`

	DoubleValue *float64 `json:"doubleValue"`

	StringValue *[]byte `json:"stringValue"`

	AggregateValue *string `json:"aggregateValue"`
}

func (s UninterpretedOptionWeaviateModel) ToProto() *UninterpretedOption {
	theProto := &UninterpretedOption{}

	for _, protoField := range s.Name {
		msg := protoField.ToProto()
		if theProto.Name == nil {
			theProto.Name = []*NamePart{msg}
		} else {
			theProto.Name = append(theProto.Name, msg)
		}
	}

	theProto.IdentifierValue = s.IdentifierValue

	theProto.PositiveIntValue = s.PositiveIntValue

	theProto.NegativeIntValue = s.NegativeIntValue

	theProto.DoubleValue = s.DoubleValue

	theProto.StringValue = s.StringValue

	theProto.AggregateValue = s.AggregateValue

	return theProto
}

func (s *UninterpretedOption) ToWeaviateModel() UninterpretedOptionWeaviateModel {
	model := UninterpretedOptionWeaviateModel{}

	for _, protoField := range s.Name {
		msg := protoField.ToWeaviateModel()
		if model.Name == nil {
			model.Name = []NamePartWeaviateModel{msg}
		} else {
			model.Name = append(model.Name, msg)
		}
	}

	model.IdentifierValue = s.IdentifierValue

	model.PositiveIntValue = s.PositiveIntValue

	model.NegativeIntValue = s.NegativeIntValue

	model.DoubleValue = s.DoubleValue

	model.StringValue = s.StringValue

	model.AggregateValue = s.AggregateValue

	return model
}

func (s UninterpretedOptionWeaviateModel) WeaviateClassName() string {
	return "UninterpretedOption"
}

func (s UninterpretedOptionWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s UninterpretedOptionWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "name",
		DataType: []string{"UninterpretedOption_NamePart"},
	}, {
		Name:     "identifierValue",
		DataType: []string{"text"},
	}, {
		Name:     "positiveIntValue",
		DataType: []string{"string"},
	}, {
		Name:     "negativeIntValue",
		DataType: []string{"string"},
	}, {
		Name:     "doubleValue",
		DataType: []string{"number"},
	}, {
		Name:     "stringValue",
		DataType: []string{"blob"},
	}, {
		Name:     "aggregateValue",
		DataType: []string{"text"},
	},
	}
}

func (s UninterpretedOptionWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"name": []map[string]string{},

		"identifierValue": s.IdentifierValue,

		"positiveIntValue": strconv.FormatUint(s.PositiveIntValue, 10),

		"negativeIntValue": strconv.FormatInt(s.NegativeIntValue, 10),

		"doubleValue": s.DoubleValue,

		"stringValue": s.StringValue,

		"aggregateValue": s.AggregateValue,
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s UninterpretedOptionWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	for _, crossReference := range s.Name {
		NameReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["name"] = append(data["name"].([]map[string]string), NameReference)
	}
	return data
}

func (s UninterpretedOptionWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s UninterpretedOptionWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s UninterpretedOptionWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s UninterpretedOptionWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type SourceCodeInfoWeaviateModel struct {

	// A Location identifies a piece of source code in a .proto file which
	// corresponds to a particular definition.  This information is intended
	// to be useful to IDEs, code indexers, documentation generators, and similar
	// tools.
	//
	// For example, say we have a file like:
	//   message Foo {
	//     optional string foo = 1;
	//   }
	// Let's look at just the field definition:
	//   optional string foo = 1;
	//   ^       ^^     ^^  ^  ^^^
	//   a       bc     de  f  ghi
	// We have the following locations:
	//   span   path               represents
	//   [a,i)  [ 4, 0, 2, 0 ]     The whole field definition.
	//   [a,b)  [ 4, 0, 2, 0, 4 ]  The label (optional).
	//   [c,d)  [ 4, 0, 2, 0, 5 ]  The type (string).
	//   [e,f)  [ 4, 0, 2, 0, 1 ]  The name (foo).
	//   [g,h)  [ 4, 0, 2, 0, 3 ]  The number (1).
	//
	// Notes:
	// - A location may refer to a repeated field itself (i.e. not to any
	//   particular index within it).  This is used whenever a set of elements are
	//   logically enclosed in a single code segment.  For example, an entire
	//   extend block (possibly containing multiple extension definitions) will
	//   have an outer location whose path refers to the "extensions" repeated
	//   field without an index.
	// - Multiple locations may have the same path.  This happens when a single
	//   logical declaration is spread out across multiple places.  The most
	//   obvious example is the "extend" block again -- there may be multiple
	//   extend blocks in the same scope, each of which will have the same path.
	// - A location's span is not always a subset of its parent's span.  For
	//   example, the "extendee" of an extension declaration appears at the
	//   beginning of the "extend" block and is shared by all extensions within
	//   the block.
	// - Just because a location's span is a subset of some other location's span
	//   does not mean that it is a descendant.  For example, a "group" defines
	//   both a type and a field in a single declaration.  Thus, the locations
	//   corresponding to the type and field and their components will overlap.
	// - Code which tries to interpret locations should probably be designed to
	//   ignore those that it doesn't understand, as more types of locations could
	//   be recorded in the future.
	Location []LocationWeaviateModel `json:"location"`
}

func (s SourceCodeInfoWeaviateModel) ToProto() *SourceCodeInfo {
	theProto := &SourceCodeInfo{}

	for _, protoField := range s.Location {
		msg := protoField.ToProto()
		if theProto.Location == nil {
			theProto.Location = []*Location{msg}
		} else {
			theProto.Location = append(theProto.Location, msg)
		}
	}

	return theProto
}

func (s *SourceCodeInfo) ToWeaviateModel() SourceCodeInfoWeaviateModel {
	model := SourceCodeInfoWeaviateModel{}

	for _, protoField := range s.Location {
		msg := protoField.ToWeaviateModel()
		if model.Location == nil {
			model.Location = []LocationWeaviateModel{msg}
		} else {
			model.Location = append(model.Location, msg)
		}
	}

	return model
}

func (s SourceCodeInfoWeaviateModel) WeaviateClassName() string {
	return "SourceCodeInfo"
}

func (s SourceCodeInfoWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s SourceCodeInfoWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "location",
		DataType: []string{"SourceCodeInfo_Location"},
	},
	}
}

func (s SourceCodeInfoWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"location": []map[string]string{},
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s SourceCodeInfoWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	for _, crossReference := range s.Location {
		LocationReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["location"] = append(data["location"].([]map[string]string), LocationReference)
	}
	return data
}

func (s SourceCodeInfoWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s SourceCodeInfoWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s SourceCodeInfoWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s SourceCodeInfoWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

type GeneratedCodeInfoWeaviateModel struct {

	// An Annotation connects some span of text in generated code to an element
	// of its generating .proto file.
	Annotation []AnnotationWeaviateModel `json:"annotation"`
}

func (s GeneratedCodeInfoWeaviateModel) ToProto() *GeneratedCodeInfo {
	theProto := &GeneratedCodeInfo{}

	for _, protoField := range s.Annotation {
		msg := protoField.ToProto()
		if theProto.Annotation == nil {
			theProto.Annotation = []*Annotation{msg}
		} else {
			theProto.Annotation = append(theProto.Annotation, msg)
		}
	}

	return theProto
}

func (s *GeneratedCodeInfo) ToWeaviateModel() GeneratedCodeInfoWeaviateModel {
	model := GeneratedCodeInfoWeaviateModel{}

	for _, protoField := range s.Annotation {
		msg := protoField.ToWeaviateModel()
		if model.Annotation == nil {
			model.Annotation = []AnnotationWeaviateModel{msg}
		} else {
			model.Annotation = append(model.Annotation, msg)
		}
	}

	return model
}

func (s GeneratedCodeInfoWeaviateModel) WeaviateClassName() string {
	return "GeneratedCodeInfo"
}

func (s GeneratedCodeInfoWeaviateModel) WeaviateClassSchema() models.Class {
	return models.Class{
		Class:      s.WeaviateClassName(),
		Properties: s.WeaviateClassSchemaProperties(),
	}
}

func (s GeneratedCodeInfoWeaviateModel) WeaviateClassSchemaProperties() []*models.Property {
	return []*models.Property{{
		Name:     "annotation",
		DataType: []string{"GeneratedCodeInfo_Annotation"},
	},
	}
}

func (s GeneratedCodeInfoWeaviateModel) Data() map[string]interface{} {
	data := map[string]interface{}{

		"annotation": []map[string]string{},
	}

	data = s.addCrossReferenceData(data)

	return data
}

func (s GeneratedCodeInfoWeaviateModel) addCrossReferenceData(data map[string]interface{}) map[string]interface{} {
	for _, crossReference := range s.Annotation {
		AnnotationReference := map[string]string{"beacon": fmt.Sprintf("weaviate://localhost/%s", crossReference.Id)}
		data["annotation"] = append(data["annotation"].([]map[string]string), AnnotationReference)
	}
	return data
}

func (s GeneratedCodeInfoWeaviateModel) Create(ctx context.Context, client *weaviate.Client, consistencyLevel string) (*data.ObjectWrapper, error) {
	return client.Data().Creator().
		WithClassName(s.WeaviateClassName()).
		WithProperties(s.Data()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s GeneratedCodeInfoWeaviateModel) Update(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Updater().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithProperties(s.Data()).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s GeneratedCodeInfoWeaviateModel) Delete(ctx context.Context, client *weaviate.Client, consistencyLevel string) error {
	return client.Data().Deleter().
		WithClassName(s.WeaviateClassName()).
		WithID(s.Id).
		WithConsistencyLevel(consistencyLevel).
		Do(ctx)
}

func (s GeneratedCodeInfoWeaviateModel) EnsureClass(client *weaviate.Client, continueOnError bool) error {
	return ensureClass(client, s.WeaviateClassSchema(), continueOnError)
}

func ensureClass(client *weaviate.Client, class models.Class, continueOnError bool) (err error) {
	var exists bool
	exists, err = classExists(client, class.Class)
	if err != nil {
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
	if err != nil || fetchedClass == nil {
		return
	}
	for _, property := range class.Properties {
		// continue on error, weaviate doesn't support updating property data types so we don't try to do that on startup
		// because it requires reindexing and is non trivial
		if containsProperty(fetchedClass.Properties, property) {
			continue
		}
		err = createProperty(client, class.Class, property)
		if err != nil {
			if !continueOnError {
				return
			}
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
