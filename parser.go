package parser

import(
	"log"
	"github.com/garebareDA/json-parser/lib/lexer"
)

func Parse(input string) error {
	tokens, err := lexer.Lex(input)
	if err != nil {
		return err
	}

	for _, token := range tokens {
		log.Println(token)
	}

	return nil
}
