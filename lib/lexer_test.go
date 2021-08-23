package lexer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonString(t *testing.T) {
	tokens, err := Lex("test");
	if err != nil {
		t.Fatalf("string error: %v", err)
	}

	fmt.Println(tokens[0]);
	assert.Equal(t, "test", tokens[0])
}