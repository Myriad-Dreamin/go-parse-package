package parser

import (
	"fmt"
	"go/ast"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

// FuncDescription Get description of a function
func FuncDescription(f interface{}) (string, error) {
	fpc := runtime.FuncForPC(reflect.ValueOf(f).Pointer())
	prf := strings.Split(filepath.Base(fpc.Name()), ".")
	var recvName, funcName string
	if len(prf) == 3 {
		recvName = prf[len(prf)-2]
		funcName = prf[len(prf)-1]
	} else {
		funcName = prf[len(prf)-1]
	}
	recvName = stripReceiver(recvName)
	for len(funcName) >= 3 && funcName[len(funcName)-3:] == "-fm" {
		funcName = funcName[:len(funcName)-3]
	}

	filename, _ := fpc.FileLine(0)
	myDoc, err := ParsePackageDocFile(filename)
	if err != nil {
		return "", err
	}
	for _, theFunc := range myDoc.Funcs {
		if theFunc.Name == funcName {
			return theFunc.Doc, nil
		}
		if theFunc.Name == recvName {
			fmt.Println(theFunc.Name)
		}
	}

	for _, obj := range myDoc.Types {
		if obj.Name != recvName {
			continue
		}
		for _, m := range obj.Methods {
			m.Recv = stripReceiver(m.Recv)
			if m.Recv == recvName && m.Name == funcName {
				return m.Doc, nil
			}
		}
		for _, m := range obj.Decl.Specs {
			switch mm := m.(type) {
			case *ast.TypeSpec:
				switch obj := mm.Type.(type) {
				case *ast.InterfaceType:
					for _, m := range obj.Methods.List {
						if len(m.Names) > 0 && m.Names[0].Name == funcName {
							if m.Doc != nil {
								return makeDoc(m.Doc.List), nil
							} else if m.Comment != nil {
								return makeDoc(m.Comment.List), nil
							}
							return "", nil
						}
					}
				}
			}
		}
	}

	//for _, obj := range myDoc.Vars {
	//	fmt.Println(obj.Names)
	//}
	//
	//for _, obj := range myDoc.Consts {
	//	fmt.Println(obj.Names)
	//}
	return "", fmt.Errorf("invalid type: %v", fpc.Name())
}

// InterfaceDescription Get description of an interface
// if v where v in InterfaceDescription(v) is an interface
// you should call it with InterfaceDescription(&v)
func InterfaceDescription(i interface{}) (string, error) {
	t := reflect.TypeOf(i)
	switch t.Kind() {
	case reflect.Ptr:
		return TypeInterfaceDescription(t.Elem())
	default:
		return "", fmt.Errorf("%T is not pointer type", i)
	}
}

// TypeInterfaceDescription Get description of a specified type
func TypeInterfaceDescription(t reflect.Type) (string, error) {
	switch t.Kind() {
	case reflect.Ptr:
		return TypeInterfaceDescription(t.Elem())
	case reflect.Interface, reflect.Struct:
		typeName := t.Name()
		path := getPackagePath(t.PkgPath())
		if len(path) == 0 {
			return "", fmt.Errorf("not registered package: %v", t.PkgPath())
		}
		myDoc, err := ParsePackageDocDir(path)
		if err != nil {
			return "", err
		}
		for _, obj := range myDoc.Types {
			if obj.Name == typeName {
				return obj.Doc, nil
			}
		}
		return "", ErrorNotFound
	default:
		return "", fmt.Errorf("invalid type: %v", t)
	}
}



func makeDoc(list []*ast.Comment) string {
	var d = make([]string, len(list))
	for i := range list {
		text := list[i].Text
		d[i] = strings.TrimPrefix(text, "//")
		if strings.HasPrefix(text, "/*") {
			d[i] = strings.TrimSuffix(strings.TrimPrefix(text, "/*"), "*/")
		}
	}
	return strings.Join(d, "\n")
}

func stripReceiver(recvName string) string {
	var flag = true
	for flag {
		flag = false
		for len(recvName) > 0 && recvName[0] == '*' {
			recvName = recvName[1:]
			flag = true
		}
		for len(recvName) > 1 && recvName[0] == '(' {
			recvName = recvName[1 : len(recvName)-1]
			flag = true
		}
	}
	return recvName
}