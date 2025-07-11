package parser

import (
	"aether/lib/utils"
	"aether/src/lexer"
)

func (p *Parser) parseFor() *For {
	if !p.expect(lexer.FOR) {
		return nil
	}
	var index *Identifier
	var value *Identifier
	if p.peekToken.Type == lexer.COMMA {
		// for i, v in ...
		value = &Identifier{Value: p.curToken.Literal}
		if !p.expect(lexer.IDENT) {
			return nil
		}
		if !p.expect(lexer.COMMA) {
			return nil
		}
		index = value
		value = &Identifier{Value: p.curToken.Literal}
		if !p.expect(lexer.IDENT) {
			return nil
		}
	} else {
		// for v in ...
		value = &Identifier{Value: p.curToken.Literal}
		if !p.expect(lexer.IDENT) {
			return nil
		}
	}
	if !p.expect(lexer.IN) {
		return nil
	}
	iterable := p.parseExpression()
	body := p.parseBlock()
	return &For{Index: index, Value: value, Iterable: iterable, Body: body}
}

func (p *Parser) parseIf() *If {
	if !p.expect(lexer.IF) {
		return nil
	}
	cond := p.parseExpression()
	if cond == nil {
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expected expression for if condition",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		return nil
	}
	body := p.parseBlock()
	if body == nil {
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expected block for if body",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		return nil
	}
	var alt *Block
	if p.curToken.Type == lexer.ELSE {
		p.nextToken()
		alt = p.parseBlock()
		if alt == nil {
			p.addError(utils.ParseError{
				Kind:    utils.InvalidSyntax,
				Message: "expected block for if alternative",
				Line:    p.curToken.Line,
				Column:  p.curToken.Column,
			})
			return nil
		}
	}
	return &If{Condition: cond, Consequence: body, Alternative: alt}
}

func (p *Parser) parseWhile() *While {
	if !p.expect(lexer.WHILE) {
		return nil
	}
	cond := p.parseExpression()
	if cond == nil {
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expected expression for while condition",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		return nil
	}
	body := p.parseBlock()
	if body == nil {
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expected block for while body",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		return nil
	}
	return &While{Condition: cond, Body: body}
}

func (p *Parser) parseRepeat() *Repeat {
	if !p.expect(lexer.REPEAT) {
		return nil
	}
	count := p.parseExpression()
	if count == nil {
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expected expression for repeat count",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		return nil
	}
	body := p.parseBlock()
	if body == nil {
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expected block for repeat body",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		return nil
	}
	return &Repeat{Count: count, Body: body}
}
