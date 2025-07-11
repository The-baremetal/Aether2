package parser

import (
	"aether/lib/utils"
	"aether/src/lexer"
)

// peekTokenN safely peeks n tokens ahead without advancing the parser state.
// n=1 returns the next token, n=2 returns the token after that, etc.
func (p *Parser) peekTokenN(n int) lexer.Token {
	tok := p.curToken
	peek := p.peekToken
	for i := 1; i < n; i++ {
		p.nextToken()
	}
	result := p.peekToken
	// Restore tokens
	p.curToken = tok
	p.peekToken = peek
	return result
}

// isAssignmentPattern checks if the current token sequence matches
// IDENT (COMMA IDENT)* ASSIGN without consuming tokens.
func (p *Parser) isAssignmentPattern() bool {
	origCur := p.curToken
	origPeek := p.peekToken
	// At least one IDENT
	if p.curToken.Type != lexer.IDENT {
		return false
	}
	// Look for (COMMA IDENT)*
	for {
		if p.peekToken.Type == lexer.COMMA {
			p.nextToken() // move to COMMA
			p.nextToken() // move to IDENT
			if p.curToken.Type != lexer.IDENT {
				// Restore tokens
				p.curToken = origCur
				p.peekToken = origPeek
				return false
			}
			continue
		}
		break
	}
	// After the last IDENT, peek for ASSIGN
	if p.peekToken.Type == lexer.ASSIGN {
		// Restore tokens
		p.curToken = origCur
		p.peekToken = origPeek
		return true
	}
	// Restore tokens
	p.curToken = origCur
	p.peekToken = origPeek
	return false
}

// parseStatement parses a single statement, which may be an assignment,
// a control structure, or a top-level expression.
func (p *Parser) parseStatement() Statement {
	// Only treat IDENT (or tuple) as assignment if the pattern matches
	if p.isAssignmentPattern() {
		names := []*Identifier{{Value: p.curToken.Literal}}
		for p.peekToken.Type == lexer.COMMA {
			// Only collect names if the token after COMMA is IDENT and the one after that is ASSIGN
			tok1 := p.peekTokenN(1)
			tok2 := p.peekTokenN(2)
			if tok1.Type == lexer.IDENT && tok2.Type == lexer.ASSIGN {
				p.nextToken() // move to COMMA
				p.nextToken() // move to IDENT
				names = append(names, &Identifier{Value: p.curToken.Literal})
				break // stop after this name, next is ASSIGN
			} else if tok1.Type == lexer.IDENT {
				p.nextToken() // move to COMMA
				p.nextToken() // move to IDENT
				names = append(names, &Identifier{Value: p.curToken.Literal})
			} else {
				break
			}
		}
		p.nextToken() // move to ASSIGN
		return p.parseAssignmentWithNames(names)
	}
	switch p.curToken.Type {
	case lexer.FUNCTION:
		return p.parseFunc()
	case lexer.STRUCT:
		return p.parseStruct()
	case lexer.IF:
		return p.parseIf()
	case lexer.WHILE:
		return p.parseWhile()
	case lexer.REPEAT:
		return p.parseRepeat()
	case lexer.RETURN:
		return p.parseReturn()
	case lexer.IMPORT:
		return p.parseImport()
	case lexer.PACKAGE:
		return p.parsePackage()
	case lexer.FOR:
		return p.parseFor()
	case lexer.LBRACE:
		return p.parseBlock()
	case lexer.MATCH:
		return p.parseMatch()
	case lexer.BREAK:
		p.nextToken()
		return &Break{}
	case lexer.CONTINUE:
		p.nextToken()
		return &Continue{}
	case lexer.C_COMMENT:
		p.nextToken()
		return nil
	case lexer.ILLEGAL:
		if p.curToken.Literal == ";" {
			snippet := ""
			if p.curToken.Line-1 < len(p.sourceLines) {
				snippet = p.sourceLines[p.curToken.Line-1]
			}
			p.addError(utils.ParseError{
				Kind:          utils.UnexpectedSemicolon,
				Message:       "Unexpected semicolon",
				Line:          p.curToken.Line,
				Column:        p.curToken.Column,
				Snippet:       snippet,
				Caret:         p.curToken.Column,
				Fix:           "Remove the semicolon",
				CodemodPrompt: "Do you want to apply the codemod to remove the semicolon? (y/n)",
			})
			p.nextToken()
			return nil
		}
		p.nextToken()
		return nil
	default:
		expr := p.parseExpression()
		if expr == nil {
			p.addError(utils.ParseError{
				Kind:    utils.InvalidSyntax,
				Message: "unexpected or invalid expression",
				Line:    p.curToken.Line,
				Column:  p.curToken.Column,
			})
			p.nextToken()
			return nil
		}
		if stmt, ok := expr.(Statement); ok {
			return stmt
		}
		// Wrap any valid expression as an ExpressionStatement
		return &ExpressionStatement{Expr: expr}
	}
}

func (p *Parser) parseAssignmentWithNames(names []*Identifier) *Assignment {
	if p.curToken.Type != lexer.ASSIGN {
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expected '=' in assignment",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		p.nextToken()
		// Always return a dummy assignment for tuple destructuring
		if len(names) > 1 {
			elems := make([]Expression, len(names))
			for i := range elems {
				elems[i] = &Literal{Value: nil}
			}
			return &Assignment{Names: names, Value: &Array{Elements: elems}}
		}
		return nil
	}
	p.nextToken()
	var value Expression
	if len(names) > 1 {
		var elems []Expression
		if p.curToken.Type == lexer.LBRACKET {
			arr := p.parseArray()
			arrayNode, ok := arr.(*Array)
			if !ok || arrayNode == nil {
				p.addError(utils.ParseError{
					Kind:    utils.InvalidSyntax,
					Message: "invalid array on right-hand side of tuple assignment",
					Line:    p.curToken.Line,
					Column:  p.curToken.Column,
				})
				elems = make([]Expression, len(names))
				for i := range elems {
					elems[i] = &Literal{Value: nil}
				}
				return &Assignment{Names: names, Value: &Array{Elements: elems}}
			}
			elems = arrayNode.Elements
		} else {
			elems = []Expression{}
			first := p.parseExpression()
			if first != nil {
				elems = append(elems, first)
			}
			for p.curToken.Type == lexer.COMMA {
				p.nextToken()
				expr := p.parseExpression()
				if expr != nil {
					elems = append(elems, expr)
				}
			}
		}
		// Pad or truncate to match names
		for len(elems) < len(names) {
			elems = append(elems, &Literal{Value: nil})
		}
		if len(elems) > len(names) {
			elems = elems[:len(names)]
		}
		// Replace any nils with Literal{Value: nil}
		for i := range elems {
			if elems[i] == nil {
				elems[i] = &Literal{Value: nil}
			}
		}
		// Always return a non-nil Array
		if elems == nil {
			elems = make([]Expression, len(names))
			for i := range elems {
				elems[i] = &Literal{Value: nil}
			}
		}
		value = &Array{Elements: elems}
		// Advance past the right-hand side if needed
		if p.curToken.Type == lexer.COMMA {
			p.nextToken()
		}
		if p.curToken.Type == lexer.EOF {
			p.nextToken()
		}
	} else {
		value = p.parseExpression()
		if p.curToken.Type == lexer.COMMA {
			p.nextToken()
		}
		if p.curToken.Type == lexer.EOF {
			p.nextToken()
		}
	}
	if value == nil {
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expected expression for assignment value",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		p.nextToken()
		// Always return a dummy assignment for tuple destructuring
		if len(names) > 1 {
			elems := make([]Expression, len(names))
			for i := range elems {
				elems[i] = &Literal{Value: nil}
			}
			return &Assignment{Names: names, Value: &Array{Elements: elems}}
		}
		return nil
	}
	// Debug printout
	// fmt.Printf("Assignment: names=%v, value=%#v\n", names, value)
	return &Assignment{Names: names, Value: value}
}

func (p *Parser) parseReturn() *Return {
	if !p.expect(lexer.RETURN) {
		return nil
	}
	expr := p.parseExpression()
	if expr == nil {
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expected expression for return value",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		return nil
	}
	return &Return{Value: expr}
}

func (p *Parser) parseImport() *Import {
	if !p.expect(lexer.IMPORT) {
		return nil
	}
	var name *Identifier
	if p.curToken.Type == lexer.STRING {
		name = &Identifier{Value: p.curToken.Literal}
		p.nextToken()
	} else if p.curToken.Type == lexer.IDENT {
		name = &Identifier{Value: p.curToken.Literal}
		p.nextToken()
	} else {
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expected module name (quoted or unquoted) after import",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		return nil
	}
	var as *Identifier
	if p.curToken.Type == lexer.AS {
		p.expect(lexer.AS)
		if p.curToken.Type == lexer.IDENT || p.curToken.Type == lexer.DOT {
			as = &Identifier{Value: p.curToken.Literal}
			p.nextToken()
		} else {
			p.addError(utils.ParseError{
				Kind:    utils.InvalidSyntax,
				Message: "expected identifier or '.' after 'as' in import",
				Line:    p.curToken.Line,
				Column:  p.curToken.Column,
			})
			return nil
		}
	}
	return &Import{Name: name, As: as}
}

func (p *Parser) parsePackage() *Package {
	if !p.expect(lexer.PACKAGE) {
		return nil
	}
	name := &Identifier{Value: p.curToken.Literal}
	if !p.expect(lexer.IDENT) {
		return nil
	}
	return &Package{Name: name}
}

func (p *Parser) parseStatementList(stop lexer.TokenType) []Statement {
	stmts := []Statement{}
	for p.curToken.Type != stop && p.curToken.Type != lexer.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			stmts = append(stmts, stmt)
		}
		// Defensive: advance to avoid infinite loop, but don't skip block closers!
		if stmt == nil && p.curToken.Type != stop && p.curToken.Type != lexer.RBRACE {
			p.nextToken()
		}
	}
	return stmts
}
