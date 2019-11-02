package parser

type PackageMapper func(packageName string) (packagePath string)

var getPackagePath PackageMapper

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
