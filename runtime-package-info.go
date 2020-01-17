package parser

type PackageMapper func(packageName string) (packagePath string)


var mapper map[string]string
func AddPacakgeMapping(name, path string) {
	if mapper == nil {
		mapper = make(map[string]string)
	}
	mapper[name] = path
}

func defaultMapper(packageName string) (packagePath string) {
	if mapper == nil {
		return ""
	}
	packagePath, _ = mapper[packageName]
	return
}

var getPackagePath PackageMapper = defaultMapper

// SetPackageMapper provide package path mapper of a interface's package
// must set your own package mapper
// e.g.
// import "github.com/Myriad-Dreamin/go-magic-package/instance"
// SetPackageMapper(instance.Get)
func SetPackageMapper(xGetPackagePath PackageMapper) PackageMapper {
	old := getPackagePath
	getPackagePath = xGetPackagePath
	return old
}
