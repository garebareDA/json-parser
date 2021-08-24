package lexer

import (
	"errors"
	"strconv"
	"fmt"
)

type lexer struct {
	input  []rune
	tokens []interface{}
}

func newLexer(input string) *lexer {
	return &lexer{
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

		jsonNull := lex.lexNull()
		if jsonNull {
			lex.tokens = append(lex.tokens, nil)
			continue
		}

		jsonNumber, err := lex.lexNumber()
		if err != nil {
			return nil, err
		}
		if jsonNumber != nil {
			lex.tokens = append(lex.tokens, jsonNumber)
			continue
		}

		jsonSyntax, err := lex.lexSyntax()
		if err != nil {
			return nil, err
		}
		if jsonSyntax != ' ' {
			lex.tokens = append(lex.tokens, jsonSyntax)
			continue
		}
	}

	return lex.tokens, nil
}

func (l *lexer) lexString() (jsonString, error) {
	if jsonQuote == l.input[0] {
		l.input = l.input[1:]
	} else {
		return "", nil
	}

	var strings []rune
	for i, char := range l.input {
		if char == jsonQuote {
			l.input = l.input[i+1:]
			return jsonString(strings), nil
		}
		strings = append(strings, char)
	}

	return "", errors.New("expected end-of-string quote")
}

func (l *lexer) lexBool() (jsonBool, bool) {
	trueLen := len("true")
	falseLen := len("false")

	if len(l.input) >= trueLen && string(l.input[0:trueLen]) == "true" {
		l.input = l.input[trueLen:]
		return jsonBool(true), true
	}

	if len(l.input) >= falseLen && string(l.input[0:falseLen]) == "false" {
		l.input = l.input[falseLen:]
		return jsonBool(false), true
	}

	return false, false
}

func (l *lexer) lexNull() bool {
	nullLen := len("null")
	if len(l.input) >= nullLen && string(l.input[0:nullLen]) == "null" {
		l.input = l.input[nullLen:]
		return true
	}

	return false
}

func (l *lexer) lexNumber() (interface{}, error) {
	jsonNumbers := []rune{}
	isFloat := false
	for _, chr := range l.input {
		_, ok := numbers[chr]
		if ok {
			isFloat = chr == 'e' || chr == '.' || chr == 'E' || isFloat
			jsonNumbers = append(jsonNumbers, chr)
		} else {
			break
		}
	}

	l.input = l.input[len(jsonNumbers):]

	if len(jsonNumbers) < 1 {
		return nil, nil
	}

	if isFloat {
		f, err := strconv.ParseFloat(string(jsonNumbers), 32)
		if err != nil {
			return nil, err
		}
		return jsonFloat(f), nil
	}

	i, err := strconv.ParseInt(string(jsonNumbers), 10, 32)
	if err != nil {
		return nil, err
	}
	return jsonInt(i), nil
}

func (l *lexer) lexSyntax() (jsonSyntax, error) {
	chr := l.input[0]
	_, ok := jsonWhiteSpace[chr]
	if ok {
		l.input = l.input[1:]
		return jsonSyntax(' '), nil
	}

	_, ok = jsonSyntaxs[chr]
	if ok {
		l.input = l.input[1:]
		return jsonSyntax(chr), nil
	}

	return jsonSyntax(' '), fmt.Errorf("unexpected character %s", string(chr))
}
