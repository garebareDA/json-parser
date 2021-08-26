package parser_test

import (
	"fmt"
	"testing"

	"github.com/garebareDA/json-parser"
)

func TestParser(t *testing.T) {
	json, err := parser.FromString(`{"name":}`)
	if err != nil {
		t.Fail()
	}

	fmt.Println(json...)
}
