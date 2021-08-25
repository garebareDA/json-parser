package lexer_test

import (
	"testing"
	"github.com/garebareDA/json-parser/lib/lexer"
	"github.com/stretchr/testify/assert"
)

func TestJsonString(t *testing.T) {
	tokens, err := lexer.Lex("\"test\"")
	if err != nil {
		t.Fatalf("string error: %v", err)
	}

	assert.Equal(t, lexer.JsonString("test"), tokens[0])
}

func TestJsonBoolTrue(t *testing.T) {
	tokens, err := lexer.Lex("true")
	if err != nil {
		t.Fatalf("bool error: %v", err)
	}
	assert.Equal(t, lexer.JsonBool(true), tokens[0])
}

func TestJsonBoolFalse(t *testing.T) {
	tokens, err := lexer.Lex("false")
	if err != nil {
		t.Fatalf("bool error: %v", err)
	}
	assert.Equal(t, lexer.JsonBool(false), tokens[0])
}

func TestJsonNull(t *testing.T) {
	tokens, err := lexer.Lex("null")
	if err != nil {
		t.Fatalf("null error: %v", err)
	}

	assert.Equal(t, nil, tokens[0])
}

func TestJsonNumber(t *testing.T) {
	tokens, err := lexer.Lex("12345")
	if err != nil {
		t.Fatalf("numberFloat error: %v", err)
	}
	assert.Equal(t, tokens[0], lexer.JsonInt(12345))
}

func TestJsonFloat(t *testing.T) {
	token, err := lexer.Lex("123.45")
	if err != nil {
		t.Fatalf("numbersInt error : %v", err)
	}
	assert.Equal(t, token[0], lexer.JsonFloat(123.45))
}

func TestJsonSyntax(t *testing.T) {
	token, err := lexer.Lex("{} [1, 1.23, ]")
	if err != nil {
		t.Fatalf("syntax error : %v", err)
	}
	assert.Equal(t, token[0], lexer.JsonSyntax('{'))
	assert.Equal(t, token[1], lexer.JsonSyntax('}'))
	assert.Equal(t, token[2], lexer.JsonSyntax('['))
	assert.Equal(t, token[3], lexer.JsonInt(1))
	assert.Equal(t, token[4], lexer.JsonSyntax(','))
	assert.Equal(t, token[5], lexer.JsonFloat(1.23))
	assert.Equal(t, token[6], lexer.JsonSyntax(','))
	assert.Equal(t, token[7], lexer.JsonSyntax(']'))
}

func TestJsonAll(t *testing.T) {
	token, err := lexer.Lex(`{"name":"a", "number": 123, "obj":{} }`)
	if err != nil {
		t.Fatalf("syntax all error : %v", err)
	}

	assert.Equal(t, token[0], lexer.JsonSyntax('{'))
	assert.Equal(t, token[1], lexer.JsonString("name"))
	assert.Equal(t, token[2], lexer.JsonSyntax(':'))
	assert.Equal(t, token[3], lexer.JsonString("a"))
	assert.Equal(t, token[4], lexer.JsonSyntax(','))
	assert.Equal(t, token[5], lexer.JsonString("number"))
	assert.Equal(t, token[6], lexer.JsonSyntax(':'))
	assert.Equal(t, token[7], lexer.JsonInt(123))
	assert.Equal(t, token[8], lexer.JsonSyntax(','))
	assert.Equal(t, token[9], lexer.JsonString("obj"))
	assert.Equal(t, token[10], lexer.JsonSyntax(':'))
	assert.Equal(t, token[11], lexer.JsonSyntax('{'))
	assert.Equal(t, token[12], lexer.JsonSyntax('}'))
	assert.Equal(t, token[13], lexer.JsonSyntax('}'))
}
