package utils

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"sort"
)

func NormalizedDescriptor(desc protoreflect.MessageDescriptor) *descriptorpb.DescriptorProto {
	descPB := protodesc.ToDescriptorProto(desc)
	normalizedPB := &descriptorpb.DescriptorProto{
		Name:      descPB.Name,
		Field:     descPB.Field,
		OneofDecl: descPB.OneofDecl,
	}
	sort.Slice(normalizedPB.Field, func(i, j int) bool {
		return normalizedPB.Field[i].GetNumber() < normalizedPB.Field[j].GetNumber()
	})

	return normalizedPB
}

func decompress(gzBytes []byte) ([]byte, error) {

}

func CheckDescriptor(msgProtoType proto.Message, gzipped bool, descBytes ...byte) bool {
	normalizedDescPB := NormalizedDescriptor(msgProtoType.ProtoReflect().Descriptor())
	if gzipped {
		var err error
		descBytes, err = decompress(descBytes)
		if err != nil {
			return false
		}
	}
	var givenDescPB descriptorpb.DescriptorProto
	if err := proto.Unmarshal(descBytes, &givenDescPB); err != nil {
		return false
	}
	return proto.Equal(normalizedDescPB, &givenDescPB)
}
