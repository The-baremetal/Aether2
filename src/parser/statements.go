package parser

import (
	"aether/lib/utils"
	"aether/src/lexer"
	"fmt"
)

// isAssignmentPattern checks if the current token sequence matches
// IDENT (COMMA IDENT)* ASSIGN without consuming tokens.
func (p *Parser) isAssignmentPattern() bool {
	if p.curToken.Type != lexer.IDENT {
		return false
	}

	// Use a local buffer to peek ahead from parser's state

tokens := []lexer.Token{p.curToken, p.peekToken}
	peekIndex := 1
	for {
		if peekIndex+1 >= len(tokens) {
			tokens = append(tokens, p.l.PeekToken(peekIndex))
		}
		nextToken := tokens[peekIndex]
		if nextToken.Type == lexer.COMMA {
			if peekIndex+2 >= len(tokens) {
				tokens = append(tokens, p.l.PeekToken(peekIndex+1))
			}
			identToken := tokens[peekIndex+1]
			if identToken.Type != lexer.IDENT {
				return false
			}
			peekIndex += 2
			continue
		}
		break
	}
	if peekIndex+1 >= len(tokens) {
		tokens = append(tokens, p.l.PeekToken(peekIndex))
	}
	assignToken := tokens[peekIndex]
	result := assignToken.Type == lexer.ASSIGN
	return result
}

func (p *Parser) parseTupleAssignment(names []*Identifier) *Assignment {
	if p.curToken.Type == lexer.ASSIGN {
		p.nextToken()
		var value Expression
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
				elems := make([]Expression, len(names))
				for i := range elems {
					elems[i] = &Literal{Value: nil}
				}
				return &Assignment{Names: names, Value: &Array{Elements: elems}}
			}
			value = arrayNode
		} else {
			elems := []Expression{}
			first := p.parseExpression()
			fmt.Println("First right side expression:", first)
			if first != nil {
				elems = append(elems, first)
			}
			for p.curToken.Type == lexer.COMMA {
				p.nextToken()
				expr := p.parseExpression()
				fmt.Println("Next right side expression:", expr)
				if expr != nil {
					elems = append(elems, expr)
				}
			}
			for len(elems) < len(names) {
				elems = append(elems, &Literal{Value: nil})
			}
			if len(elems) > len(names) {
				elems = elems[:len(names)]
			}
			for i := range elems {
				if elems[i] == nil {
					elems[i] = &Literal{Value: nil}
				}
			}
			if elems == nil {
				elems = make([]Expression, len(names))
				for i := range elems {
					elems[i] = &Literal{Value: nil}
				}
			}
			value = &Array{Elements: elems}
		}
		
		if p.curToken.Type == lexer.COMMA {
			p.nextToken()
		}
		if p.curToken.Type == lexer.EOF {
			p.nextToken()
		}
		
		return &Assignment{Names: names, Value: value}
	}
	
	p.nextToken()
	return nil
}

// parseStatement parses a single statement, which may be an assignment,
// a control structure, or a top-level expression.
func (p *Parser) parseStatement() Statement {
if p.isAssignmentPattern() {
	// starts at the first IDENT
    names := []*Identifier{{Value: p.curToken.Literal}}
    
    for p.peekToken.Type == lexer.COMMA {
        p.nextToken()
        p.nextToken()
        names = append(names, &Identifier{Value: p.curToken.Literal})
    }
    
    p.nextToken()
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
	fmt.Println(p.curToken, p.peekToken)
	
	if len(names) > 1 {
		result := p.parseTupleAssignment(names)
		if result != nil {
			return result
		}
	}
	p.nextToken()
	var value Expression
	if len(names) > 1 {
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "multiple names without comma should be handled by parseTupleAssignment",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		return nil
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
		return nil
	}
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
