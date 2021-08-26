package parser_test

import (
	"fmt"
	"testing"

	"github.com/garebareDA/json-parser"
)

func TestParser(t *testing.T) {
	json, err := parser.FromString(`{"b" : {"a": 123}, "a":[{"123":123}, 123]}`)

	if err != nil {
		t.Fatalf("perse error: %s", err)
	}
	fmt.Println(json...)
}
