package lexer_test

import (
	"github.com/garebareDA/json-parser/lib/lexer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonString(t *testing.T) {
	tokens, err := lexer.Lex("\"test\"")
	if err != nil {
		t.Fatalf("string error: %v", err)
	}

	assert.Equal(t, "test", tokens[0])
}

func TestJsonBoolTrue(t *testing.T) {
	tokens, err := lexer.Lex("true")
	if err != nil {
		t.Fatalf("bool error: %v", err)
	}
	assert.Equal(t, true, tokens[0])
}

func TestJsonBoolFalse(t *testing.T) {
	tokens, err := lexer.Lex("false")
	if err != nil {
		t.Fatalf("bool error: %v", err)
	}
	assert.Equal(t, false, tokens[0])
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

	assert.Equal(t, tokens[0], int32(12345))
}

func TestJsonFloat(t *testing.T) {
	token, err := lexer.Lex("123.45")
	if err != nil {
		t.Fatalf("numbersInt error : %v", err)
	}
	assert.Equal(t, token[0], float32(123.45))
}

func TestJsonSyntax(t *testing.T) {
	token, err := lexer.Lex("{} [1, 1.23, ]")
	if err != nil {
		t.Fatalf("syntax error : %v", err)
	}
	assert.Equal(t, token[0], lexer.JsonSyntax('{'))
	assert.Equal(t, token[1], lexer.JsonSyntax('}'))
	assert.Equal(t, token[2], lexer.JsonSyntax('['))
	assert.Equal(t, token[3], int32(1))
	assert.Equal(t, token[4], lexer.JsonSyntax(','))
	assert.Equal(t, token[5], float32(1.23))
	assert.Equal(t, token[6], lexer.JsonSyntax(','))
	assert.Equal(t, token[7], lexer.JsonSyntax(']'))
}

func TestJsonAll(t *testing.T) {
	token, err := lexer.Lex(`{"name":"a", "number": 123, "obj":{} }`)
	if err != nil {
		t.Fatalf("syntax all error : %v", err)
	}

	assert.Equal(t, token[0], lexer.JsonSyntax('{'))
	assert.Equal(t, token[1], "name")
	assert.Equal(t, token[2], lexer.JsonSyntax(':'))
	assert.Equal(t, token[3], "a")
	assert.Equal(t, token[4], lexer.JsonSyntax(','))
	assert.Equal(t, token[5], "number")
	assert.Equal(t, token[6], lexer.JsonSyntax(':'))
	assert.Equal(t, token[7], int32(123))
	assert.Equal(t, token[8], lexer.JsonSyntax(','))
	assert.Equal(t, token[9], "obj")
	assert.Equal(t, token[10], lexer.JsonSyntax(':'))
	assert.Equal(t, token[11], lexer.JsonSyntax('{'))
	assert.Equal(t, token[12], lexer.JsonSyntax('}'))
	assert.Equal(t, token[13], lexer.JsonSyntax('}'))
}
