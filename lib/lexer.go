package lexer

import (
	"fmt"
)

const jsonQuote = '/'

type jsonString string

func Lex(input string) ([]interface{}, error) {
	str := []rune(input)
	var tokens []interface{}

	for len(str) > 0 {
		jsonStrings, ret, err := lexString(str)
		str = ret;
		if err != nil {
			return nil, err
		}
		if len(jsonStrings) > 0 {
			tokens = append(tokens, jsonStrings)
			continue
		}
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
			str = str[i:]
			return jsonString(strings), str, nil
		}
		strings = append(strings, char)
	}

	return "", str, fmt.Errorf("Expected end-of-string quote")
}
