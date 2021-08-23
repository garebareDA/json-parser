package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonString(t *testing.T) {
	tokens, err := Lex("\"test\"");
	if err != nil {
		t.Fatalf("string error: %v", err)
	}

	assert.Equal(t, jsonString("test"), tokens[0])
}