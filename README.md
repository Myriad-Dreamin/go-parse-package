# go-parse-package
 get the package info of function/variable at runtime

## Installation

```bash
go get github.com/Myriad-Dreamin/go-parse-package
```
or just import it in your package

## Functions

### runtime parser

```go
func FuncDescription(f interface{}) (string, error)
```
`FuncDescription` Get description of a function

```go
func InterfaceDescription(i interface{}) (string, error)
```
`InterfaceDescription` Get description of an interface
if v where v in `InterfaceDescription(v)` is an interface you should call it with `InterfaceDescription(&v)`

```go
func TypeInterfaceDescription(t reflect.Type) (string, error)
```
`TypeInterfaceDescription` Get description of a specified type

### parsing tools

```go
func ParsePackage(fileName string, mode parser.Mode) (*ast.Package, error)
```
`ParsePackage` Parse the package of a file with `parse.Mode`

```go
func ParsePackageDir(fileDir string, mode parser.Mode) (*ast.Package, error)
```
`ParsePackageDir` Parse the package of all files in directory with `parse.Mode`

```go
func ParsePackageDoc(path string) (*doc.Package, error)
```
`ParsePackageDoc` Parse the package by path

```go
func ParsePackageDocDir(pkgDir string) (*doc.Package, error)
```
`ParsePackageDocDir` Parse the package of all files in directory

```go
func ParsePackageDocFile(fileName string) (*doc.Package, error)
```
`ParsePackageDocFile` Parse the package of a file

```go
func ParsePackageName(path string) (string, error)
```
`ParsePackageName` Parse the package by path

```go
func ParsePackageNameDir(pkgDir string) (string, error)
```
`ParsePackageNameDir` Parse the package name of all files in directory

```go
func ParsePackageNameFile(fileName string) (string, error)
```
`ParsePackageNameFile` Parse the package name of a file


## Types

```go
type PackageMapper func(packageName string) (packagePath string)
```

```go
func SetPackageMapper(xGetPackagePath PackageMapper) PackageMapper
```
`SetPackageMapper` provide package path mapper of a interface's package
you must set your own package mapper before calling function `TypeInterfaceDescription`
for example

```go
import "github.com/Myriad-Dreamin/go-magic-package/instance"
... //(in your function)
    SetPackageMapper(instance.Get)
    TypeInterfaceDescription(reflect.Typeof(&MyInterface))
```