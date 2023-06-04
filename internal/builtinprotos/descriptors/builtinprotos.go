package descriptors

import (
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

//go:generate ./gen-imports.sh

const (
	builtinProtoPrefix = `google/protobuf/`
)

// GetAll returns a topologically sorted (that is, respecting import dependencies) list of all builtin protobuf files
// (i.e., those that are shipped with the golang protobuf library).
func GetAll() []protoreflect.FileDescriptor {
	var builtinProtoFiles []protoreflect.FileDescriptor
	protoregistry.GlobalFiles.RangeFiles(func(desc protoreflect.FileDescriptor) bool {
		if strings.HasPrefix(desc.Path(), builtinProtoPrefix) {
			builtinProtoFiles = append(builtinProtoFiles, desc)
		}
		return true
	})
	return topSortFiles(builtinProtoFiles)
}
