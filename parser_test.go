package parser_test

import (
	"fmt"
	"testing"

	"github.com/garebareDA/json-parser"
)

func TestParser(t *testing.T) {
	json, err := parser.FromString(`{"name" : 23}`)
	if err != nil {
		t.Fatalf("perse error: %s", err)
	}
	fmt.Println(json...)
}
