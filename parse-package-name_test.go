package parser

import (
	"fmt"
	"testing"
)

func TestParsePackageName(t *testing.T) {
	fmt.Println(ParsePackageName("./parse-doc-at-runtime.go"))
	fmt.Println(ParsePackageName("./"))
}