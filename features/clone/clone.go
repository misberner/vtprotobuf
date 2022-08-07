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

func (p *clone) messageShallow(message *protogen.Message) {
	ccTypeName := message.GoIdent
	shallowCloneName := "ShallowCloneVT"
	p.P(`func (m *`, ccTypeName, `) `, shallowCloneName, `() (*`, ccTypeName, `) {`)
	p.P(`if m == nil {`)
	p.P(`return nil`)
	p.P(`}`)
	p.P(`return &`, ccTypeName, `{`)

	oneofs := make(map[string]struct{}, len(message.Oneofs))
	for _, field := range message.Fields {
		fieldname := field.GoName
		oneof := field.Oneof != nil && !field.Oneof.Desc.IsSynthetic()
		if oneof {
			fieldname = field.Oneof.GoName
			if _, ok := oneofs[fieldname]; ok {
				continue
			}
			oneofs[fieldname] = struct{}{}
		}
		p.P(fieldname, `: m.`, fieldname, `,`)
	}
	p.P(`}`)
	p.P(`}`)
}

func (p *clone) message(proto3 bool, message *protogen.Message) {
	for _, nested := range message.Messages {
		p.message(proto3, nested)
	}

	if message.Desc.IsMapEntry() {
		return
	}

	p.once = true

	p.messageShallow(message)
}
