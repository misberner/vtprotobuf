// Copyright (c) 2021 PlanetScale Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package generator

import (
	"fmt"
	"github.com/planetscale/vtprotobuf/generator/builtins"
	"github.com/planetscale/vtprotobuf/vtproto"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type GeneratedFile struct {
	*protogen.GeneratedFile
	Ext              *Extensions
	TargetGoPackages map[protogen.GoImportPath]struct{}
}

func (p *GeneratedFile) Ident(path, ident string) string {
	return p.QualifiedGoIdent(protogen.GoImportPath(path).Ident(ident))
}

func (b *GeneratedFile) ShouldPool(message *protogen.Message) bool {
	if message == nil {
		return false
	}
	if b.Ext.Poolable[message.GoIdent] {
		return true
	}
	ext := proto.GetExtension(message.Desc.Options(), vtproto.E_Mempool)
	if mempool, ok := ext.(bool); ok {
		return mempool
	}
	return false
}

func (b *GeneratedFile) Alloc(vname string, message *protogen.Message) {
	if b.ShouldPool(message) {
		b.P(vname, " := ", message.GoIdent, `FromVTPool()`)
	} else {
		b.P(vname, " := new(", message.GoIdent, `)`)
	}
}

func (p *GeneratedFile) FieldGoType(field *protogen.Field) (goType string, pointer bool) {
	if field.Desc.IsWeak() {
		return "struct{}", false
	}

	pointer = field.Desc.HasPresence()
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		goType = "bool"
	case protoreflect.EnumKind:
		goType = p.QualifiedGoIdent(field.Enum.GoIdent)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		goType = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		goType = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		goType = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		goType = "uint64"
	case protoreflect.FloatKind:
		goType = "float32"
	case protoreflect.DoubleKind:
		goType = "float64"
	case protoreflect.StringKind:
		goType = "string"
	case protoreflect.BytesKind:
		goType = "[]byte"
		pointer = false // rely on nullability of slices for presence
	case protoreflect.MessageKind, protoreflect.GroupKind:
		goType = "*" + p.QualifiedGoIdent(field.Message.GoIdent)
		pointer = false // pointer captured as part of the type
	}
	switch {
	case field.Desc.IsList():
		return "[]" + goType, false
	case field.Desc.IsMap():
		keyType, _ := p.FieldGoType(field.Message.Fields[0])
		valType, _ := p.FieldGoType(field.Message.Fields[1])
		return fmt.Sprintf("map[%v]%v", keyType, valType), false
	}
	return goType, pointer
}

// IsGenerationTargetType returns a boolean indicating whether the type identified by the argument
// is a (message or oneof-wrapper) type for which code is generated.
func (p *GeneratedFile) IsGenerationTargetType(typeIdent protogen.GoIdent) bool {
	_, ok := p.TargetGoPackages[typeIdent.GoImportPath]
	return ok
}

// flattenInto recursively flattens the input slice, by recursively expanding any slice elements that are
// of type []interface{}. nil interface elements are removed.
func flattenInto(out, in []interface{}) []interface{} {
	for _, elem := range in {
		switch t := elem.(type) {
		case nil:
			continue
		case []interface{}:
			out = flattenInto(out, t)
		default:
			out = append(out, t)
		}
	}
	return out
}

func (p *GeneratedFile) P(args ...interface{}) {
	p.GeneratedFile.P(flattenInto(nil, args)...)
}

// GetSupportPkgFor returns the import path of the support package containing optimized functions operating on the
// given message type, if any.
func (p *GeneratedFile) GetSupportPkgFor(ty protogen.GoIdent) protogen.GoImportPath {
	if p.IsGenerationTargetType(ty) {
		if p.Ext.SupportGenPrefix != "" {
			return p.Ext.SupportGenPrefix + `/` + ty.GoImportPath
		}
		return ""
	}
	return builtins.LookupSupportPkg(ty)
}

func (p *GeneratedFile) FastCallExpr(methodName, varName string, varTy protogen.GoIdent, args ...interface{}) []interface{} {
	if supportPkg := p.GetSupportPkgFor(varTy); supportPkg != "" {
		return p.X(supportPkg.Ident(methodName+"_"+varTy.GoName), `(`, varName, `, `, args, `)`)
	}
	if p.IsGenerationTargetType(varTy) {
		return p.X(varName, `.`, methodName, `(`, args, `)`)
	}
	return nil
}

func (p *GeneratedFile) FuncHeader(name string, receiver string, receiverType protogen.GoIdent, params, returns interface{}) {
	receiverQType := p.QualifiedGoIdent(receiverType)
	if receiverQType == receiverType.GoName {
		p.P(`func (`, receiver, ` *`, receiverQType, `) `, name, `(`, params, `) (`, returns, `) {`)
	} else {
		p.P(`func `, name, `_`, receiverType.GoName, `(`, receiver, ` *`, receiverQType, `, `, params, `) (`, returns, `) {`)
	}
}

func (p *GeneratedFile) GetUnknownFieldsExpr(x string) interface{} {
	if false {
		return fmt.Sprintf("%s.unknownFields", x)
	}
	return fmt.Sprintf("%s.ProtoReflect().GetUnknown()", x)
}

func (p *GeneratedFile) SetUnknownFieldsStmt(x string, rhsExpr ...interface{}) {
	if !p.IsExternal() {
		p.P(x, `.unknownFields = `, rhsExpr)
	} else {
		p.P(x, `.ProtoReflect().SetUnknown(`, rhsExpr, `)`)
	}
}

func (p *GeneratedFile) X(args ...interface{}) []interface{} {
	return args
}

func (p *GeneratedFile) IsExternal() bool {
	return p.Ext.SupportGenPrefix != ""
}

func (p *GeneratedFile) InterfaceForOneof(oneof *protogen.Oneof) string {
	if p.IsExternal() {
		return "interface{}"
	}
	return fmt.Sprintf("is%s", oneof.GoIdent.GoName)
}
