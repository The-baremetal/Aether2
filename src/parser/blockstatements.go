package parser

import (
	"aether/lib/utils"
	"aether/src/lexer"
	"fmt"
)

func (p *Parser) parseBlock() *Block {
	fmt.Printf("🍕 parseBlock: starting, curToken: %s '%s'\n", p.curToken.Type, p.curToken.Literal)
	if !p.expect(lexer.LBRACE) {
		fmt.Printf("🍕 parseBlock: failed to expect LBRACE\n")
		return nil
	}
	fmt.Printf("🍕 parseBlock: after LBRACE, curToken: %s '%s'\n", p.curToken.Type, p.curToken.Literal)
	block := &Block{Statements: []Statement{}}
	for p.curToken.Type != lexer.RBRACE && p.curToken.Type != lexer.EOF {
		fmt.Printf("🍕 parseBlock: parsing statement, curToken: %s '%s'\n", p.curToken.Type, p.curToken.Literal)
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
			fmt.Printf("🍕 parseBlock: added statement: %T\n", stmt)
		} else {
			fmt.Printf("🍕 parseBlock: got nil statement\n")
			p.addError(utils.ParseError{
				Kind:    utils.InvalidSyntax,
				Message: "nil statement in block",
				Line:    p.curToken.Line,
				Column:  p.curToken.Column,
			})
			// Defensive: advance to avoid infinite loop
			if p.curToken.Type != lexer.RBRACE && p.curToken.Type != lexer.EOF {
				p.nextToken()
			}
		}
	}
	fmt.Printf("🍕 parseBlock: before RBRACE check, curToken: %s '%s'\n", p.curToken.Type, p.curToken.Literal)
	if !p.expect(lexer.RBRACE) {
		fmt.Printf("🍕 parseBlock: failed to expect RBRACE\n")
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expected } to close block",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		// Return the block even if we don't find the closing brace
		// This allows partial parsing to continue
	}
	fmt.Printf("🍕 parseBlock: returning block with %d statements\n", len(block.Statements))
	return block
}

func (p *Parser) parseFunc() *Function {
	if !p.expect(lexer.FUNCTION) {
		return nil
	}
	var name *Identifier
	if p.curToken.Type == lexer.IDENT {
		name = &Identifier{Value: p.curToken.Literal}
		p.nextToken()
	} else {
		name = &Identifier{Value: ""}
	}

	fmt.Printf("🍕 parseFunc: after name, curToken: %s '%s'\n", p.curToken.Type, p.curToken.Literal)

	params := []*Identifier{}
	if p.curToken.Type == lexer.LPAREN {
		p.nextToken()
		for p.curToken.Type != lexer.RPAREN && p.curToken.Type != lexer.EOF {
			// Check if this is a vararg parameter (only at the beginning of parameter name)
			if p.curToken.Type == lexer.VARARG {
				if p.peekToken.Type == lexer.IDENT {
					p.nextToken() // consume VARARG
					param := &Identifier{Value: p.curToken.Literal, IsVararg: true}
					if !p.expect(lexer.IDENT) {
						return nil
					}
					params = append(params, param)
				} else {
					p.addError(utils.ParseError{
						Kind:    utils.InvalidSyntax,
						Message: "invalid use of ... in parameter list",
						Line:    p.curToken.Line,
						Column:  p.curToken.Column,
					})
					return nil
				}
			} else {
				// Regular parameter
				param := &Identifier{Value: p.curToken.Literal}
				if !p.expect(lexer.IDENT) {
					return nil
				}
				params = append(params, param)
			}
			// Handle type annotation
			if p.curToken.Type == lexer.COLON {
				p.expect(lexer.COLON)
				if p.curToken.Type == lexer.IDENT {
					params[len(params)-1].Type = p.curToken.Literal
					p.expect(lexer.IDENT)
				}
			}
			if p.curToken.Type == lexer.COMMA {
				p.expect(lexer.COMMA)
			}
		}
		if !p.expect(lexer.RPAREN) {
			return nil
		}
	}

	fmt.Printf("🍕 parseFunc: before parseBlock, curToken: %s '%s'\n", p.curToken.Type, p.curToken.Literal)
	body := p.parseBlock()
	fmt.Printf("🍕 parseFunc: after parseBlock, body: %+v\n", body)
	return &Function{Name: name, Params: params, Body: body}
}

func (p *Parser) parseMatch() *Match {
	if !p.expect(lexer.MATCH) {
		return nil
	}
	var expr Expression
	switch p.curToken.Type {
	case lexer.IDENT:
		expr = &Identifier{Value: p.curToken.Literal}
		p.nextToken()
	case lexer.NUMBER:
		expr = &Literal{Value: p.curToken.Literal}
		p.nextToken()
	case lexer.STRING:
		expr = &Literal{Value: p.curToken.Literal}
		p.nextToken()
	default:
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expected identifier or literal after match",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		return nil
	}
	// Expect and consume LBRACE
	if !p.expect(lexer.LBRACE) {
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expected { after match expression",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		return nil
	}
	cases := []*Case{}
	for p.curToken.Type == lexer.CASE {
		p.nextToken()
		pat := p.parsePattern()
		if pat == nil {
			p.addError(utils.ParseError{
				Kind:    utils.InvalidSyntax,
				Message: "expected pattern for case",
				Line:    p.curToken.Line,
				Column:  p.curToken.Column,
			})
			return nil
		}
		body := p.parseBlock()
		if body == nil {
			p.addError(utils.ParseError{
				Kind:    utils.InvalidSyntax,
				Message: "expected block for case body",
				Line:    p.curToken.Line,
				Column:  p.curToken.Column,
			})
			return nil
		}
		cases = append(cases, &Case{Pattern: pat, Body: body})
	}
	if !p.expect(lexer.RBRACE) {
		return nil
	}
	// Debug printout
	// fmt.Printf("Match: expr=%#v, cases=%#v\n", expr, cases)
	return &Match{Expr: expr, Cases: cases}
}

func (p *Parser) parsePattern() Expression {
	switch p.curToken.Type {
	case lexer.IDENT:
		// Regular identifier pattern
		pat := &Identifier{Value: p.curToken.Literal}
		p.nextToken()
		return pat

	case lexer.UNDERSCORE:
		// Wildcard pattern
		pat := &Identifier{Value: "_"}
		p.nextToken()
		return pat

	case lexer.NUMBER, lexer.STRING:
		// Literal pattern
		pat := &Literal{Value: p.curToken.Literal}
		p.nextToken()
		return pat

	case lexer.LBRACKET:
		return p.parseArrayPattern()

	case lexer.LBRACE:
		return p.parseStructPattern()

	default:
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "invalid pattern",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		return nil
	}
}

func (p *Parser) parseStructPattern() Expression {
	// For now, just parse as a simple identifier
	// TODO: Implement proper struct pattern parsing
	pat := &Identifier{Value: p.curToken.Literal}
	p.nextToken()
	return pat
}

func (p *Parser) parseArrayPattern() Expression {
	// For now, just parse as a simple identifier
	// TODO: Implement proper array pattern parsing
	pat := &Identifier{Value: p.curToken.Literal}
	p.nextToken()
	return pat
}
