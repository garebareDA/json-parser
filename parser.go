package parser

import (
	"fmt"
	"github.com/garebareDA/json-parser/lib/lexer"
)

type parser struct {
	token []interface{}
	json  []interface{}
}

func newParser(token []interface{}) parser {
	return parser{
		token: token,
	}
}

func FromString(str string) ([]interface{}, error) {
	tokens, err := lexer.Lex(str)
	if err != nil {
		return nil, err
	}

	parser := newParser(tokens)
	err = parser.parseJson()
	if err != nil {
		return nil, err
	}

	return parser.json, nil
}

func (p *parser) parseJson() error {
	for len(p.token) > 0 {
		if p.token[0] != lexer.JsonSyntax(lexer.JsonLeftBrace) {
			return fmt.Errorf("syntax error left brace: %s", p.token[0:])
		}

		json, err := p.objectParse()
		if err != nil {
			return err
		}
		p.json = append(p.json, json)
	}
	return nil
}

func (p *parser) objectParse() (map[string]interface{}, error) {
	jsonObject := make(map[string]interface{})
	if p.token[0] != lexer.JsonSyntax(lexer.JsonLeftBrace) {
		return nil, fmt.Errorf("syntax error left brace: %s", p.token[0:])
	}

	for {
		p.token = p.token[1:]
		s, ok := p.token[0].(string)
		if ok {
			p.token = p.token[1:]
			if p.token[0] != lexer.JsonSyntax(lexer.JsonColon) {
				return nil, fmt.Errorf("syntax error colon : %s", p.token[0:])
			}

			p.token = p.token[1:]
			t := p.token[0]
			js, ok := t.(lexer.JsonSyntax)

			if js == lexer.JsonLeftBrace {
				obj, err := p.objectParse()
				if err != nil {
					return nil, err
				}
				jsonObject[s] = obj
			}

			if js == lexer.JsonLeftBracket {
				
			}

			if !ok {
				jsonObject[s] = t
				p.token = p.token[1:]
			}

			if p.token[0] == lexer.JsonComma {
				continue
			}
		}

		if p.token[0] == lexer.JsonSyntax(lexer.JsonRightBrace) {
			p.token = p.token[1:]
			return jsonObject, nil
		}

		return nil, fmt.Errorf("syntax error right brace : %s", p.token[0])
	}
}
