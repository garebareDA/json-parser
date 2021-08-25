package lexer

const (
	jsonQuote        = '"'
	jsonComma        = ','
	jsonColon        = ':'
	jsonLeftBracket  = '['
	jsonRightBracket = ']'
	jsonLeftBrace    = '{'
	jsonRightBrace   = '}'
)

var jsonWhiteSpace = map[rune]struct{}{
	' ':  {},
	'\t': {},
	'\b': {},
	'\n': {},
	'\r': {},
}

var jsonSyntaxs = map[rune]struct{}{
	jsonColon:        {},
	jsonComma:        {},
	jsonLeftBrace:    {},
	jsonRightBrace:   {},
	jsonLeftBracket:  {},
	jsonRightBracket: {},
}

type jsonString string
type jsonBool bool
type jsonInt int32
type jsonFloat float32
type jsonSyntax rune

var numbers map[rune]struct{} = createNumbers()
func createNumbers() map[rune]struct{} {
	numbers := map[rune]struct{}{}

	for i := 0; i < 10; i++ {
		numbers[rune(i+48)] = struct{}{}
	}
	numbers['.'] = struct{}{}
	numbers['E'] = struct{}{}
	numbers['e'] = struct{}{}
	numbers['+'] = struct{}{}
	numbers['-'] = struct{}{}
	return numbers
}
