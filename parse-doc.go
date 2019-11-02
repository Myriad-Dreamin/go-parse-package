package parser

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
)


// ParsePackage Parse the package of fileName with parse.Mode
func ParsePackage(fileName string, mode parser.Mode) (*ast.Package, error) {
	fset := token.NewFileSet()
	parsedAst, err := parser.ParseFile(fset, fileName, nil, mode)
	if err != nil {
		return nil, err
	}

	pkg := &ast.Package{
		Name:  "Any",
		Files: make(map[string]*ast.File),
	}
	pkg.Files[fileName] = parsedAst
	return pkg, nil
}

// ParsePackageDir Parse the package of all files in directory with parse.Mode
func ParsePackageDir(fileDir string, mode parser.Mode) (*ast.Package, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, fileDir, nil, mode)
	if err != nil {
		return nil, err
	}
	if len(pkgs) != 1 {
		return nil, fmt.Errorf("invalid parsed package result: %v", pkgs)
	}
	var pkg *ast.Package
	for _, p := range pkgs {
		pkg = p
	}
	return pkg, nil
}

// ParsePackageDoc Parse the package of path with parse.Mode
func ParsePackageDoc(path string) (*doc.Package, error) {
	if s, err := os.Stat(path); err != nil {
		return nil, err
	} else if s.IsDir() {
		return ParsePackageDocDir(path)
	} else {
		return ParsePackageDocFile(path)
	}
}

// ParsePackageDocFile Parse the package of a file in directory with parse.Mode
func ParsePackageDocFile(fileName string) (*doc.Package, error) {
	pkg, err := ParsePackage(fileName, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return doc.New(pkg, "/", doc.AllDecls), nil
}

// ParsePackageDocDir Parse the package of all files in directory with parse.Mode
func ParsePackageDocDir(pkgDir string) (*doc.Package, error) {
	pkg, err := ParsePackageDir(pkgDir, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return doc.New(pkg, "/", doc.AllDecls), nil
}

// ParsePackageName Parse the package of path with parse.Mode
func ParsePackageName(path string) (string, error) {
	if s, err := os.Stat(path); err != nil {
		return "", err
	} else if s.IsDir() {
		return ParsePackageNameDir(path)
	} else {
		return ParsePackageNameFile(path)
	}
}

// ParsePackageNameFile Parse the package name of a file in directory with parse.Mode
func ParsePackageNameFile(fileName string) (string, error) {
	pkg, err := ParsePackage(fileName, parser.PackageClauseOnly)
	if err != nil {
		return "", err
	}
	return pkg.Name, nil
}

// ParsePackageNameDir Parse the package name of all files in directory with parse.Mode
func ParsePackageNameDir(pkgDir string) (string, error) {
	pkg, err := ParsePackageDir(pkgDir, parser.PackageClauseOnly)
	if err != nil {
		return "", err
	}
	return pkg.Name, nil
}
