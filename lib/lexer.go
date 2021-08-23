package lexer

const jsonQuote = '/'

type jsonString string

func Lex(input string) ([]interface{}, error) {
	str := []rune(input)
	var tokens []interface{}

	for len(str) > 0 {
		var jsonString string
		var err error

	}

	return tokens, nil
}

func lexString(str []rune) jsonString {
	if jsonQuote == str[0] {
		str = str[1:]
	} else {
		return ""
	}

	for _, char := range str {
		if char == jsonQuote {

		}
	}

	return ""
}
