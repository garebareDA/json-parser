package lexer

import (
	"errors"
	"log"
)

const jsonQuote = '"'

type jsonString string
type jsonBool bool

type lexer struct {
	input []rune
	tokens []interface{}
}

func newLexer(input string) lexer {
	return lexer {
		input: []rune(input),
	}
}

func Lex(input string) ([]interface{}, error) {
	lex := newLexer(input)
	for len(lex.input) > 0 {
		jsonStrings, err := lex.lexString()
		if err != nil {
			return nil, err
		}
		if len(jsonStrings) > 0 {
			lex.tokens = append(lex.tokens, jsonStrings)
			continue
		}

		jsonBool, isOk := lex.lexBool()
		if isOk {
			lex.tokens = append(lex.tokens, jsonBool)
			continue
		}

		return nil, errors.New("unexpected character " + string(lex.input[0]))
	}

	return lex.tokens, nil
}

func (l *lexer) lexString () (jsonString, error) {
	if jsonQuote == l.input[0] {
		l.input = l.input[1:]
	} else {
		return "", nil
	}

	var strings []rune
	for i, char := range l.input {
		if char == jsonQuote {
			l.input = l.input[i+1:]
			return jsonString(strings),  nil
		}
		strings = append(strings, char)
	}

	return "", errors.New("expected end-of-string quote")
}

func (l *lexer) lexBool() (jsonBool, bool) {
	trueLen := len("true")
	falseLen := len("false")

	if string(l.input[0:trueLen]) == "true" {
		l.input = l.input[trueLen:]
		log.Println(l.input)
		return jsonBool(true), true
	}

	if string(l.input[0:falseLen]) == "false" {
		l.input = l.input[falseLen:]
		return jsonBool(false), true
	}

	return false, false
}
