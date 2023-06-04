package builtins

import "google.golang.org/protobuf/compiler/protogen"

var (
	supportPkgs = make(map[string]string)
)

func register(goType string, supportPkg string) {
	supportPkgs[goType] = supportPkg
}

func LookupSupportPkg(goType protogen.GoIdent) protogen.GoImportPath {
	return protogen.GoImportPath(supportPkgs[goType.String()])
}
