package main

import (
	"flag"
	"fmt"
	"strings"

	_ "github.com/planetscale/vtprotobuf/features/clone"
	_ "github.com/planetscale/vtprotobuf/features/equal"
	_ "github.com/planetscale/vtprotobuf/features/grpc"
	_ "github.com/planetscale/vtprotobuf/features/marshal"
	_ "github.com/planetscale/vtprotobuf/features/pool"
	_ "github.com/planetscale/vtprotobuf/features/size"
	_ "github.com/planetscale/vtprotobuf/features/unmarshal"
	"github.com/planetscale/vtprotobuf/generator"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

type ObjectSet map[protogen.GoIdent]bool

func (o ObjectSet) String() string {
	return fmt.Sprintf("%#v", o)
}

func (o ObjectSet) Set(s string) error {
	idx := strings.LastIndexByte(s, '.')
	if idx < 0 {
		return fmt.Errorf("invalid object name: %q", s)
	}

	ident := protogen.GoIdent{
		GoImportPath: protogen.GoImportPath(s[0:idx]),
		GoName:       s[idx+1:],
	}
	o[ident] = true
	return nil
}

func main() {
	var allowEmpty bool
	var features string
	poolable := make(ObjectSet)
	var supportGen string

	var f flag.FlagSet
	f.BoolVar(&allowEmpty, "allow-empty", false, "allow generation of empty files")
	f.Var(poolable, "pool", "use memory pooling for this object")
	f.StringVar(&features, "features", "all", "list of features to generate (separated by '+')")
	f.StringVar(&supportGen, "support-gen", "", "if nonempty, generate support code")

	protogen.Options{ParamFunc: f.Set}.Run(func(plugin *protogen.Plugin) error {
		return generateAllFiles(plugin, strings.Split(features, "+"), poolable, allowEmpty, supportGen)
	})
}

var SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

func generateAllFiles(plugin *protogen.Plugin, featureNames []string, poolable ObjectSet, allowEmpty bool, supportGen string) error {
	ext := &generator.Extensions{Poolable: poolable}
	gen, err := generator.NewGenerator(plugin.Files, featureNames, ext)
	if err != nil {
		return err
	}

	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		goImportPath := file.GoImportPath
		if supportGen != "" {
			goImportPath = protogen.GoImportPath(supportGen) + `/` + goImportPath
		}
		gf := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+"_vtproto.pb.go", goImportPath)
		if !gen.GenerateFile(gf, file) && !allowEmpty {
			gf.Skip()
		}
	}

	plugin.SupportedFeatures = SupportedFeatures
	return nil
}
