package parser_test

import (
	"testing"

	"github.com/garebareDA/json-parser"
	"github.com/stretchr/testify/assert"
)

func TestParseString(t *testing.T) {
	json, err := parser.FromString(`{"a":"a"}`)
	if err != nil {
		t.Fatalf("perse error: %s", err)
	}

	s, ok := json[0]["a"].(string)
	if !ok {
		t.Fatalf("TestParserString not ok a")
	}

	assert.Equal(t, s, "a")
}

func TestParseInt(t *testing.T) {
	json, err := parser.FromString(`{"a":123}`)
	if err != nil {
		t.Fatalf("perse error: %s", err)
	}

	s, ok := json[0]["a"].(int32)
	if !ok {
		t.Fatalf("TestParserInt not ok a")
	}

	assert.Equal(t, s, int32(123))
}

func TestParseFloat(t *testing.T) {
	json, err := parser.FromString(`{"a":1.23}`)
	if err != nil {
		t.Fatalf("perse error: %s", err)
	}

	s, ok := json[0]["a"].(float32)
	if !ok {
		t.Fatalf("TestParserFloat not ok a")
	}

	assert.Equal(t, s, float32(1.23))
}

func TestParseNull(t *testing.T) {
	json, err := parser.FromString(`{"a":null}`)
	if err != nil {
		t.Fatalf("perse error: %s", err)
	}

	s, ok := json[0]["a"]
	if !ok {
		t.Fatalf("TestParserNull not ok a")
	}

	assert.Equal(t, s, nil)
}

func TestParserObject(t *testing.T) {
	json, err := parser.FromString(`{"a" : {"b": 123, "c":{"d":"aaa"}}, "c":123.1}`)
	if err != nil {
		t.Fatalf("perse error: %s", err)
	}

	jsonObj := json[0]
	aObj := jsonObj["a"]
	a, ok := aObj.(map[string]interface{})
	if !ok {
		t.Fatalf("TestParserObject not ok a")
	}

	i, ok := a["b"].(int32)
	if !ok {
		t.Fatalf("TestParserObject not ok b")
	}

	c, ok := a["c"].(map[string]interface{})
	if !ok {
		t.Fatalf("TestParserObject not ok c")
	}

	d, ok := c["d"].(string)
	if !ok {
		t.Fatalf("TestParserObject not ok d")
	}

	assert.Equal(t, i, int32(123))
	assert.Equal(t, d, "aaa")

	cs, ok := jsonObj["c"].(float32)
	if !ok {
		t.Fatalf("TestParserObject not ok c2")
	}
	assert.Equal(t, cs, float32(123.1))
}

func TestParseArray(t *testing.T) {
	json, err := parser.FromString(`{"a":[1, 1, [1, 1, {"a": 1}]]}`)
	if err != nil {
		t.Fatalf("perse error: %s", err)
	}

	jsonObj := json[0]
	aObj := jsonObj["a"]
	a, ok := aObj.([]interface{})
	if !ok {
		t.Fatalf("TestParserArray not ok a")
	}

	a0, ok := a[0].(int32)
	if !ok {
		t.Fatalf("TestParserArray not ok a0")
	}

	a1, ok := a[1].(int32)
	if !ok {
		t.Fatalf("TestParserArray not ok a1")
	}

	assert.Equal(t, a0, int32(1))
	assert.Equal(t, a1, int32(1))

	aa, ok := a[2].([]interface{})
	if !ok {
		t.Fatalf("TestParserArray not ok aa")
	}

	aa0, ok := aa[0].(int32)
	if !ok {
		t.Fatalf("TestParserArray not ok aa0")
	}

	aa1, ok := aa[1].(int32)
	if !ok {
		t.Fatalf("TestParserArray not ok aa1")
	}

	aa2, ok := aa[2].(map[string]interface{})
	if !ok {
		t.Fatalf("TestParserArray not ok aa2")
	}

	aa2a, ok := aa2["a"].(int32)
	if !ok {
		t.Fatalf("TestParserArray not ok aa2[a]")
	}

	assert.Equal(t, aa0, int32(1))
	assert.Equal(t, aa1, int32(1))
	assert.Equal(t, aa2a, int32(1))
}
