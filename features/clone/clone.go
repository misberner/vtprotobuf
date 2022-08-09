// Copyright (c) 2021 PlanetScale Inc. All rights reserved.
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clone

import (
	"github.com/planetscale/vtprotobuf/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	cloneName = "CloneVT"
)

var (
	protoPkg = protogen.GoImportPath("google.golang.org/protobuf/proto")
)

func init() {
	generator.RegisterFeature("clone", func(gen *generator.GeneratedFile) generator.FeatureGenerator {
		return &clone{GeneratedFile: gen}
	})
}

type clone struct {
	*generator.GeneratedFile
	once bool
}

var _ generator.FeatureGenerator = (*clone)(nil)

func (p *clone) Name() string {
	return "clone"
}

func (p *clone) GenerateFile(file *protogen.File) bool {
	proto3 := file.Desc.Syntax() == protoreflect.Proto3
	for _, message := range file.Messages {
		p.message(proto3, message)
	}

	return p.once
}

func (p *clone) GenerateHelpers() {
}

func (p *clone) oneofFieldClone(lhsBase, rhsBase string, oneof *protogen.Oneof) {
	fieldname := oneof.GoName
	ccInterfaceName := "is" + oneof.GoIdent.GoName
	lhs := lhsBase + "." + fieldname
	rhs := rhsBase + "." + fieldname
	p.P(`if `, rhs, ` != nil {`)
	p.P(lhs, ` = `, rhs, `.(interface{ `, cloneName, `() `, ccInterfaceName, ` }).`, cloneName, `()`)
	p.P(`}`)
}

func (p *clone) fieldCloneSimple(lhs, rhs string, kind protoreflect.Kind, fieldGoType string, localMessage bool) {
	switch {
	case kind == protoreflect.MessageKind, kind == protoreflect.GroupKind:
		if localMessage {
			p.P(lhs, ` = `, rhs, `.`, cloneName, `()`)
		} else {
			p.P(`if vtpb, ok := `, protoPkg.Ident("Message"), `(`, rhs, `).(interface{ `, cloneName, `() `, fieldGoType, ` }); ok {`)
			p.P(lhs, ` = vtpb.`, cloneName, `()`)
			p.P(`} else {`)
			p.P(lhs, ` = `, protoPkg.Ident("Clone"), `(`, rhs, `).(`, fieldGoType, `)`)
			p.P(`}`)
		}
	case kind == protoreflect.BytesKind:
		p.P(`tmpBytes := make([]byte, len(`, rhs, `))`)
		p.P(`copy(tmpBytes, `, rhs, `)`)
		p.P(lhs, ` = tmpBytes`)
	case isScalar(kind):
		p.P(lhs, ` = `, rhs)
	default:
		panic("unexpected")
	}
}

func (p *clone) fieldClone(lhsBase, rhsBase string, allFieldsNullable bool, field *protogen.Field) {
	// At this point, if we encounter a non-synthetic oneof, we assume it to be the representative
	// field for that oneof.
	if field.Oneof != nil && !field.Oneof.Desc.IsSynthetic() {
		p.oneofFieldClone(lhsBase, rhsBase, field.Oneof)
		return
	}

	if !isReference(allFieldsNullable, field) {
		panic("method should not be invoked for non-reference fields")
	}

	fieldname := field.GoName
	lhs := lhsBase + "." + fieldname
	rhs := rhsBase + "." + fieldname

	// At this point, we are only looking at reference types (pointers, maps, slices, interfaces), which can all
	// be nil.
	p.P(`if rhs := `, rhs, `; rhs != nil {`)
	rhs = "rhs"

	goType, _ := p.FieldGoType(field)
	fieldKind := field.Desc.Kind()
	localMessage := false
	if msg := field.Message; msg != nil {
		localMessage = p.IsLocalMessage(msg)
	}

	if field.Desc.Cardinality() == protoreflect.Repeated { // maps and slices
		p.P(`tmpContainer := make(`, goType, `, len(`, rhs, `))`)
		if isScalar(fieldKind) && field.Desc.IsList() {
			// Generated code optimization: instead of iterating over all (key/index, value) pairs,
			// do a single copy(dst, src) invocation for slices whose elements aren't reference types.
			p.P(`copy(tmpContainer, `, rhs, `)`)
		} else {
			if field.Desc.IsMap() {
				valueField := field.Message.Fields[1]
				fieldKind = valueField.Desc.Kind()
				goType, _ = p.FieldGoType(valueField)
				localMessage = false
				if msg := field.Message.Fields[1].Message; msg != nil {
					localMessage = p.IsLocalMessage(msg)
				}
			} else {
				goType = goType[2:] // strip []
			}
			p.P(`for k, v := range `, rhs, ` {`)
			p.fieldCloneSimple("tmpContainer[k]", "v", fieldKind, goType, localMessage)
			p.P(`}`)
		}
		p.P(lhs, ` = tmpContainer`)
	} else if isScalar(fieldKind) {
		p.P(`tmpVal := *`, rhs)
		p.P(lhs, ` = &tmpVal`)
	} else {
		p.fieldCloneSimple(lhs, rhs, fieldKind, goType, localMessage)
	}
	p.P(`}`)
}

func (p *clone) messageDeep(proto3 bool, message *protogen.Message) {
	ccTypeName := message.GoIdent.GoName
	p.P(`func (m *`, ccTypeName, `) `, cloneName, `() *`, ccTypeName, ` {`)
	p.body(!proto3, ccTypeName, message.Fields)
	p.P(`}`)
	p.P()
}

func (p *clone) body(allFieldsNullable bool, ccTypeName string, fields []*protogen.Field) {
	// The method body for a message or a oneof wrapper always starts with a nil check.
	p.P(`if m == nil {`)
	// We use an explicitly typed nil to avoid returning the nil interface in the oneof wrapper
	// case.
	p.P(`return (*`, ccTypeName, `)(nil)`)
	p.P(`}`)

	// Make a first pass over the fields, in which we initialize all non-reference fields via direct
	// struct literal initialization, and extract all other (refernece) fields for a second pass.
	p.P(`r := &`, ccTypeName, `{`)
	var refFields []*protogen.Field
	oneofFields := make(map[string]struct{}, len(fields))

	for _, field := range fields {
		if field.Oneof != nil && !field.Oneof.Desc.IsSynthetic() {
			// Use the first field in a oneof as the representative for that oneof, disregard
			// the other fields in that oneof.
			if _, ok := oneofFields[field.Oneof.GoName]; !ok {
				refFields = append(refFields, field)
				oneofFields[field.Oneof.GoName] = struct{}{}
			}
			continue
		}

		if !isReference(allFieldsNullable, field) {
			p.P(field.GoName, `: m.`, field.GoName, `,`)
			continue
		}
		// Shortcut: for types where we know that an optimized clone method exists, we can call it directly as it is
		// nil-safe.
		if field.Desc.Cardinality() != protoreflect.Repeated && field.Message != nil && p.IsLocalMessage(field.Message) {
			p.P(field.GoName, `: m.`, field.GoName, `.`, cloneName, `(),`)
			continue
		}
		refFields = append(refFields, field)
	}
	p.P(`}`)

	// Generate explicit assignment statements for all reference fields.
	for _, field := range refFields {
		p.fieldClone("r", "m", allFieldsNullable, field)
	}

	p.P(`return r`)
}

func (p *clone) oneofDeep(field *protogen.Field) {
	ccTypeName := field.GoIdent.GoName
	ccInterfaceName := "is" + field.Oneof.GoIdent.GoName
	p.P(`func (m *`, ccTypeName, `) `, cloneName, `() `, ccInterfaceName, ` {`)

	// Create a "fake" field for the single oneof member, pretending it is not a oneof field.
	fieldInOneof := *field
	fieldInOneof.Oneof = nil
	// If we have a scalar field in a oneof, that field is never nullable, even when using proto2
	p.body(false, ccTypeName, []*protogen.Field{&fieldInOneof})
	p.P(`}`)
	p.P()
}

func (p *clone) messageOneofs(message *protogen.Message) {
	for _, field := range message.Fields {
		if field.Oneof == nil || field.Oneof.Desc.IsSynthetic() {
			continue
		}
		p.oneofDeep(field)
	}
}

func (p *clone) message(proto3 bool, message *protogen.Message) {
	for _, nested := range message.Messages {
		p.message(proto3, nested)
	}

	if message.Desc.IsMapEntry() {
		return
	}

	p.once = true

	p.messageDeep(proto3, message)
	p.messageOneofs(message)
}

func isReference(allFieldsNullable bool, field *protogen.Field) bool {
	if allFieldsNullable || field.Oneof != nil || field.Message != nil || field.Desc.Cardinality() == protoreflect.Repeated || field.Desc.Kind() == protoreflect.BytesKind {
		return true
	}
	if !isScalar(field.Desc.Kind()) {
		panic("unexpected non-reference, non-scalar field")
	}
	return false
}

func isScalar(kind protoreflect.Kind) bool {
	switch kind {
	case
		protoreflect.BoolKind,
		protoreflect.StringKind,
		protoreflect.DoubleKind, protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind,
		protoreflect.FloatKind, protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind,
		protoreflect.Int64Kind, protoreflect.Uint64Kind, protoreflect.Sint64Kind,
		protoreflect.Int32Kind, protoreflect.Uint32Kind, protoreflect.Sint32Kind,
		protoreflect.EnumKind:
		return true
	}
	return false
}
