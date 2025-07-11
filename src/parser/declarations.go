package parser

import (
	"aether/src/lexer"
)

func (p *Parser) parseStruct() *StructDef {
	if !p.expect(lexer.STRUCT) {
		return nil
	}
	name := &Identifier{Value: p.curToken.Literal}
	if !p.expect(lexer.IDENT) {
		return nil
	}
	if !p.expect(lexer.LBRACE) {
		return nil
	}
	fields := []*Field{}
	for p.curToken.Type != lexer.RBRACE && p.curToken.Type != lexer.EOF {
		fieldName := &Identifier{Value: p.curToken.Literal}
		if !p.expect(lexer.IDENT) {
			return nil
		}
		var fieldType string
		if p.curToken.Type == lexer.COLON {
			p.expect(lexer.COLON)
			if p.curToken.Type == lexer.IDENT {
				fieldType = p.curToken.Literal
				p.expect(lexer.IDENT)
			}
		}
		fields = append(fields, &Field{Name: fieldName, Type: fieldType})
		if p.curToken.Type == lexer.COMMA {
			p.expect(lexer.COMMA)
		}
	}
	if !p.expect(lexer.RBRACE) {
		return nil
	}
	return &StructDef{Name: name, Fields: fields}
}
