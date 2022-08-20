package main

import (
	"flag"
	"fmt"
	"github.com/planetscale/vtprotobuf/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
	"io"
	"os"
	"path/filepath"
	"sort"

	_ "github.com/planetscale/vtprotobuf/features/clone"
	_ "github.com/planetscale/vtprotobuf/features/equal"
	_ "github.com/planetscale/vtprotobuf/features/marshal"
	_ "github.com/planetscale/vtprotobuf/features/size"
	_ "github.com/planetscale/vtprotobuf/features/unmarshal"
)

//go:generate ./gen-imports.sh

func topSortRec(out *[]*descriptorpb.FileDescriptorProto, curr protoreflect.FileDescriptor, files map[string]protoreflect.FileDescriptor) {
	delete(files, curr.Path())
	for i := 0; i < curr.Imports().Len(); i++ {
		f := files[curr.Imports().Get(i).Path()]
		if f != nil {
			topSortRec(out, f, files)
		}
	}
	*out = append(*out, protodesc.ToFileDescriptorProto(curr))
}

func topSortFiles(files map[string]protoreflect.FileDescriptor) []*descriptorpb.FileDescriptorProto {
	var result []*descriptorpb.FileDescriptorProto
	for _, desc := range files {
		topSortRec(&result, desc, files)
	}
	return result
}

func registerMessages(w io.Writer, pkg protogen.GoImportPath, msgs []*protogen.Message) {
	for _, msg := range msgs {
		registerMessages(w, pkg, msg.Messages)
		fmt.Fprintf(w, "\tregister(`%s`, `%s`)\n", msg.GoIdent.String(), string(pkg))
	}
}

func main() {
	var origReq pluginpb.CodeGeneratorRequest
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	if err := proto.Unmarshal(data, &origReq); err != nil {
		panic(err)
	}
	req := &pluginpb.CodeGeneratorRequest{}
	req.CompilerVersion = origReq.CompilerVersion
	req.Parameter = origReq.Parameter
	allFiles := make(map[string]protoreflect.FileDescriptor)
	protoregistry.GlobalFiles.RangeFiles(func(desc protoreflect.FileDescriptor) bool {
		allFiles[desc.Path()] = desc
		req.FileToGenerate = append(req.FileToGenerate, desc.Path())
		return true
	})
	sort.Strings(req.FileToGenerate)
	req.ProtoFile = topSortFiles(allFiles)

	var fs flag.FlagSet
	var registryFile, registryPackage string
	fs.StringVar(&registryFile, "registry_file", "", "")
	fs.StringVar(&registryPackage, "registry_package", "", "")
	p, err := protogen.Options{ParamFunc: fs.Set}.New(req)
	if err != nil {
		panic(err)
	}
	if registryFile == "" {
		panic("must specify registry package")
	}
	if registryPackage == "" {
		registryPackage = filepath.Base(filepath.Dir(registryFile))
	}
	p.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	ext := &generator.Extensions{
		SupportGenPrefix: "github.com/planetscale/vtprotobuf/support/types",
	}
	gen, err := generator.NewGenerator(p.Files, []string{"all"}, ext)
	if err != nil {
		panic(err)
	}

	registryF, err := os.Create(registryFile)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(registryF, `//go:build !nobuiltins`)
	fmt.Fprintln(registryF, `// +build !nobuiltins`)
	fmt.Fprintln(registryF)

	fmt.Fprintf(registryF, "package %s\n\nfunc init() {\n", registryPackage)

	for _, file := range p.Files {
		if !file.Generate {
			fmt.Fprintln(os.Stderr, "not generating for ", file.Desc.Path())
			continue
		}

		goImportPath := ext.SupportGenPrefix + `/` + file.GoImportPath
		gf := p.NewGeneratedFile(file.GeneratedFilenamePrefix+"_vtproto.pb.go", goImportPath)
		if !gen.GenerateFile(gf, file) {
			fmt.Fprintln(os.Stderr, "skippy")
			gf.Skip()
		}
		fmt.Fprintln(os.Stderr, "generated for ", file.GeneratedFilenamePrefix)
		registerMessages(registryF, goImportPath, file.Messages)
	}

	fmt.Fprintln(registryF, `}`)
	if err := registryF.Close(); err != nil {
		panic(err)
	}

	resp := p.Response()
	respBytes, err := proto.Marshal(resp)
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.Write(respBytes)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr, "all done")
}
