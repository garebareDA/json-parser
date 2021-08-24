package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonString(t *testing.T) {
	tokens, err := Lex("\"test\"")
	if err != nil {
		t.Fatalf("string error: %v", err)
	}

	assert.Equal(t, jsonString("test"), tokens[0])
}

func TestJsonBoolTrue(t *testing.T) {
	tokens, err := Lex("true")
	if err != nil {
		t.Fatalf("bool error: %v", err)
	}
	assert.Equal(t, jsonBool(true), tokens[0])
}

func TestJsonBoolFalse(t *testing.T) {
	tokens, err := Lex("false")
	if err != nil {
		t.Fatalf("bool error: %v", err)
	}
	assert.Equal(t, jsonBool(false), tokens[0])
}

func TestJsonNull(t *testing.T) {
	tokens, err := Lex("null")
	if err != nil {
		t.Fatalf("null error: %v", err)
	}

	assert.Equal(t, nil, tokens[0])
}

func TestJsonNumber(t *testing.T) {
	tokens, err := Lex("12345")
	if err != nil {
		t.Fatalf("numberFloat error: %v", err)
	}
	assert.Equal(t, tokens[0], jsonInt(12345))
}

func TestJsonFloat(t *testing.T) {
	token, err := Lex("123.45")
	if err != nil {
		t.Fatalf("numbersInt error : %v", err)
	}
	assert.Equal(t, token[0], jsonFloat(123.45))
}
