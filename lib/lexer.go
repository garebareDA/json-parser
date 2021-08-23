package lexer

import (
	"errors"
)

const jsonQuote = '"'

type jsonString string

func Lex(input string) ([]interface{}, error) {
	str := []rune(input)
	var tokens []interface{}

	for len(str) > 0 {
		jsonStrings, s, err := lexString(str)
		str = s;
		if err != nil {
			return nil, err
		}
		if len(jsonStrings) > 0 {
			tokens = append(tokens, jsonStrings)
			continue
		}

		return nil, errors.New("unexpected character " + string(str))
	}

	return tokens, nil
}

func lexString(str []rune) (jsonString, []rune, error) {
	if jsonQuote == str[0] {
		str = str[1:]
	} else {
		return "", str, nil
	}

	var strings []rune
	for i, char := range str {
		if char == jsonQuote {
			str = str[i+1:]
			return jsonString(strings), str, nil
		}
		strings = append(strings, char)
	}

	return "", str, errors.New("expected end-of-string quote")
}
