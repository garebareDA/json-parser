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
	p.token = p.token[1:]

	for {
		s, ok := p.token[0].(string)
		if ok {
			p.token = p.token[1:]
			if p.token[0] != lexer.JsonSyntax(lexer.JsonColon) {
				return nil, fmt.Errorf("syntax error colon : %s", p.token[0:])
			}

			p.token = p.token[1:]
			t := p.token[0]
			value, err := p.valueParser(t)
			if err != nil {
				return nil, err
			}

			jsonObject[s] = value
			if p.token[0] == lexer.JsonSyntax(lexer.JsonComma) {
				p.token = p.token[1:]
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

func (p *parser) arrayParser() ([]interface{}, error) {
	var jsonArray []interface{}
	if p.token[0] != lexer.JsonSyntax(lexer.JsonLeftBracket) {
		return nil, fmt.Errorf("syntax error left bracket: %s", p.token[0:])
	}

	for {
		p.token = p.token[1:]
		t := p.token[0]
		value, err := p.valueParser(t)
		if err != nil {
			return nil, err
		}

		jsonArray = append(jsonArray, value)
		if p.token[0] == lexer.JsonSyntax(lexer.JsonComma) {
			continue
		}

		if p.token[0] == lexer.JsonSyntax(lexer.JsonRightBracket) {
			return jsonArray, nil
		}

		return nil, fmt.Errorf("syntax error right bracket: %s", p.token[0:])
	}
}

func (p *parser) valueParser(t interface{}) (interface{}, error) {
	js, ok := t.(lexer.JsonSyntax)
	if js == lexer.JsonLeftBrace {
		obj, err := p.objectParse()
		if err != nil {
			return nil, err
		}
		return obj, nil
	}

	if js == lexer.JsonLeftBracket {
		array, err := p.arrayParser()
		if err != nil {
			return nil, err
		}
		p.token = p.token[1:]
		return array, nil
	}

	if !ok {
		p.token = p.token[1:]
		return t, nil
	}

	if t == nil {
		p.token = p.token[1:]
		return nil, nil
	}

	return nil, fmt.Errorf("syntax error invald vlue: %s", p.token[0:])
}
