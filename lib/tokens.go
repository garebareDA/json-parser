package lexer
const jsonQuote = '"'

var numbers map[rune]rune = createNumbers()

type jsonString string
type jsonBool bool
type jsonInt int32
type jsonFloat float32

func createNumbers() map[rune]rune {
	numbers := map[rune]rune{}
	for i := 0; i < 10; i++ {
		numbers[rune(i+48)] = rune(i + 48)
	}
	numbers['.'] = '.'
	numbers['E'] = 'E'
	numbers['e'] = 'e'
	numbers['+'] = '+'
	numbers['-'] = '-'
	return numbers
}