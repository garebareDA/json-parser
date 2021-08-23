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
