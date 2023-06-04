package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/planetscale/vtprotobuf/generator"
	"github.com/planetscale/vtprotobuf/internal/builtinprotos/descriptors"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/types/pluginpb"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/planetscale/vtprotobuf/features/clone"
	_ "github.com/planetscale/vtprotobuf/features/equal"
	_ "github.com/planetscale/vtprotobuf/features/marshal"
	_ "github.com/planetscale/vtprotobuf/features/size"
	_ "github.com/planetscale/vtprotobuf/features/unmarshal"
)

func runWithRequest(req *pluginpb.CodeGeneratorRequest) (*pluginpb.CodeGeneratorResponse, error) {
	var fs flag.FlagSet
	var (
		supportPackageRoot string
		registryFile       string
		registryPackage    string
	)
	fs.StringVar(&supportPackageRoot, "support_pkg_root", "",
		"The root of the support package in which to place generated code. This should be the import path "+
			"of the output directory.")
	fs.StringVar(&registryFile, "registry_file", "",
		"The Go file into which to write registry information for the main generator.")
	fs.StringVar(&registryPackage, "registry_package", "",
		"The package of the registry Go file. If unset, this is determined from the containing directory.")

	plugin, err := protogen.Options{ParamFunc: fs.Set}.New(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create generator plugin: %w", err)
	}
	supportPackageRoot = strings.TrimRight(supportPackageRoot, "/")
	if supportPackageRoot == "" {
		return nil, errors.New("-support_pkg_root must be specified")
	}
	if registryFile == "" {
		return nil, errors.New("-registry_file must be specified")
	}
	if registryPackage == "" {
		registryPackage = filepath.Base(filepath.Dir(registryFile))
	}

	plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	ext := &generator.Extensions{
		SupportGenPrefix: protogen.GoImportPath(supportPackageRoot),
	}
	gen, err := generator.NewGenerator(plugin.Files, []string{"all"}, ext)
	if err != nil {
		return nil, fmt.Errorf("error instantiating generator: %w", err)
	}

	for _, file := range plugin.Files {
		if !file.Generate {
			fmt.Fprintln(os.Stderr, "skip ", file.Desc.Path())
			continue
		}
		fmt.Fprintln(os.Stderr, "noskip ", file.Desc.Path())

		goImportPath := ext.SupportGenPrefix + `/` + file.GoImportPath
		gf := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+"_vtproto.pb.go", goImportPath)
		if !gen.GenerateFile(gf, file) {
			gf.Skip()
		}
	}

	return plugin.Response(), nil
}

func runMain() error {
	// Parse a CodeGeneratorRequest from stdin
	inputData, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to read data from stdin: %w", err)
	}
	var origReq pluginpb.CodeGeneratorRequest
	if err := proto.Unmarshal(inputData, &origReq); err != nil {
		return fmt.Errorf("failed to unmarshal CodeGeneratorRequest from stdin: %w", err)
	}

	// Sanity check: because we discard the original generator request, make sure it really is an intentional "dummy"
	// input file.
	if len(origReq.ProtoFile) != 1 || len(origReq.ProtoFile[0].MessageType) > 0 || len(origReq.ProtoFile[0].Service) > 0 {
		return fmt.Errorf("%s should be invoked with a dummy input file containing no definitions", os.Args[0])
	}

	// Assemble a new CodeGeneratorRequest for the builtin protos
	actualReq := &pluginpb.CodeGeneratorRequest{
		Parameter:       origReq.Parameter,
		CompilerVersion: origReq.CompilerVersion,
	}

	builtinFileDescs := descriptors.GetAll()
	for _, desc := range builtinFileDescs {
		fmt.Fprintln(os.Stderr, "builtin file ", desc.Path())
		actualReq.ProtoFile = append(actualReq.ProtoFile, protodesc.ToFileDescriptorProto(desc))
		actualReq.FileToGenerate = append(actualReq.FileToGenerate, desc.Path())
	}

	resp, err := runWithRequest(actualReq)
	if err != nil {
		return fmt.Errorf("error running code generator: %w", err)
	}

	// Write the marshaled response to stdout
	respBytes, err := proto.Marshal(resp)
	if err != nil {
		return fmt.Errorf("error marshaling CodeGeneratorResponse: %w", err)
	}
	_, err = os.Stdout.Write(respBytes)
	if err != nil {
		return fmt.Errorf("error writing response to stdout: %w", err)
	}
	return nil
}

func main() {
	if err := runMain(); err != nil {
		log.Fatalf("Generation for builtin types failed: %v\n", err)
	}
}
