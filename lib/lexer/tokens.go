package lexer

const (
	JsonQuote        = '"'
	JsonComma        = ','
	JsonColon        = ':'
	JsonLeftBracket  = '['
	JsonRightBracket = ']'
	JsonLeftBrace    = '{'
	JsonRightBrace   = '}'
)

var jsonWhiteSpace = map[rune]struct{}{
	' ':  {},
	'\t': {},
	'\b': {},
	'\n': {},
	'\r': {},
}

var jsonSyntaxs = map[rune]struct{}{
	JsonColon:        {},
	JsonComma:        {},
	JsonLeftBrace:    {},
	JsonRightBrace:   {},
	JsonLeftBracket:  {},
	JsonRightBracket: {},
}

type JsonString string
type JsonBool bool
type JsonInt int32
type JsonFloat float32
type JsonSyntax rune

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
