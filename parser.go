package parser

import(
	"log"
	"github.com/garebareDA/json-parser/lib/lexer"
)

func FromString(str string) error {
	tokens, err := lexer.Lex(str)
	if err != nil {
		return err
	}

	for _, token := range tokens {
		log.Println(token)
	}

	return nil
}

func parse() {
	
}
