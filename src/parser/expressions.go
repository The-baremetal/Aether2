package parser

import (
	"aether/lib/utils"
	"aether/src/lexer"
)

func parseLiteralForOperator(op lexer.TokenType) string {
	switch op {
	case lexer.PLUS:
		return "+"
	case lexer.MINUS:
		return "-"
	case lexer.ASTERISK:
		return "*"
	case lexer.SLASH:
		return "/"
	case lexer.MODULO:
		return "%"
	case lexer.EXPONENT:
		return "^"
	case lexer.EQ:
		return "=="
	case lexer.NOT_EQ:
		return "!="
	case lexer.LT:
		return "<"
	case lexer.GT:
		return ">"
	case lexer.LE:
		return "<="
	case lexer.GE:
		return ">="
	case lexer.CONCAT:
		return ".."
	default:
		return "?"
	}
}

func (p *Parser) parsePrimary() Expression {
	var expr Expression
	switch p.curToken.Type {
	case lexer.IDENT:
		// Prevent assignment from being parsed as an expression
		if p.peekToken.Type == lexer.ASSIGN {
			p.addError(utils.ParseError{
				Kind:    utils.InvalidSyntax,
				Message: "assignment is not a valid expression",
				Line:    p.curToken.Line,
				Column:  p.curToken.Column,
			})
			return nil
		}
		// Only parse struct instantiation if the next token is LBRACE and we're not in a match context
		if p.peekToken.Type == lexer.LBRACE && !p.isParsingMatch {
			expr = p.parseStructInstantiation()
		} else {
			expr = &Identifier{Value: p.curToken.Literal}
			p.nextToken()
		}
	case lexer.NUMBER, lexer.STRING:
		expr = &Literal{Value: p.curToken.Literal}
		p.nextToken()
	case lexer.LBRACKET:
		expr = p.parseArray()
	case lexer.LBRACE:
		// Try to parse as anonymous struct first, then lambda if that fails
		saveToken := p.curToken
		structExpr := p.parseAnonymousStruct()
		if structExpr != nil {
			expr = structExpr
		} else {
			p.curToken = saveToken // restore if struct parse failed
			expr = p.parseLambda()
		}
	case lexer.FUNCTION:
		expr = p.parseFunc()
	case lexer.VARARG:
		// Handle spread operator in expressions
		expr = p.parseSpread()
	case lexer.LPAREN:
		p.nextToken()
		expr = p.parseExpression()
		if !p.expect(lexer.RPAREN) {
			p.addError(utils.ParseError{Kind: utils.InvalidSyntax, Message: "expected )", Line: p.curToken.Line, Column: p.curToken.Column})
			return nil
		}
	default:
		p.addError(utils.ParseError{Kind: utils.InvalidSyntax, Message: "unexpected token in parsePrimary", Line: p.curToken.Line, Column: p.curToken.Column})
		return nil
	}
	for {
		if p.curToken.Type == lexer.LPAREN {
			expr = p.parseCallExpr(expr)
			continue
		}
		if p.curToken.Type == lexer.DOT {
			p.nextToken()
			if p.curToken.Type != lexer.IDENT {
				p.addError(utils.ParseError{Kind: utils.InvalidSyntax, Message: "expected property name after .", Line: p.curToken.Line, Column: p.curToken.Column})
				return nil
			}
			prop := &Identifier{Value: p.curToken.Literal}
			p.nextToken()
			expr = &PropertyAccess{Object: expr, Property: prop}
			continue
		}
		if p.curToken.Type == lexer.LBRACKET {
			p.nextToken()
			index := p.parseExpression()
			if index == nil {
				p.addError(utils.ParseError{Kind: utils.InvalidSyntax, Message: "expected expression for array index", Line: p.curToken.Line, Column: p.curToken.Column})
				return nil
			}
			if !p.expect(lexer.RBRACKET) {
				p.addError(utils.ParseError{Kind: utils.InvalidSyntax, Message: "expected ] after array index", Line: p.curToken.Line, Column: p.curToken.Column})
				return nil
			}
			expr = &ArrayIndex{Array: expr, Index: index}
			continue
		}
		break
	}
	return expr
}

func (p *Parser) parseCallExpr(fn Expression) Expression {
	if !p.expect(lexer.LPAREN) {
		p.addError(utils.ParseError{Kind: utils.InvalidSyntax, Message: "expected (", Line: p.curToken.Line, Column: p.curToken.Column})
		return nil
	}
	args := []Expression{}
	partial := false
	for {
		if p.curToken.Type == lexer.RPAREN || p.curToken.Type == lexer.EOF {
			break
		}
		var arg Expression
		if p.curToken.Type == lexer.VARARG {
			if p.peekToken.Type == lexer.IDENT {
				arg = &Spread{Name: p.peekToken.Literal}
				p.nextToken()
				p.nextToken()
			} else {
				arg = &Spread{Name: ""}
				p.nextToken()
			}
		} else {
			arg = p.parseExpression()
			if arg == nil {
				p.addError(utils.ParseError{Kind: utils.InvalidSyntax, Message: "nil call argument", Line: p.curToken.Line, Column: p.curToken.Column})
				return nil
			}
		}
		if ident, ok := arg.(*Identifier); ok && ident.Value == "_" {
			partial = true
		}
		args = append(args, arg)
		if p.curToken.Type == lexer.COMMA {
			p.nextToken()
		} else {
			break
		}
	}
	if !p.expect(lexer.RPAREN) {
		p.addError(utils.ParseError{Kind: utils.InvalidSyntax, Message: "expected ) after call args", Line: p.curToken.Line, Column: p.curToken.Column})
		return nil
	}
	if partial {
		return &PartialApplication{Function: fn, Args: args}
	}
	return &Call{Function: fn, Args: args}
}

func (p *Parser) parseSpread() *Spread {
	if !p.expect(lexer.VARARG) {
		return nil
	}
	if p.curToken.Type == lexer.IDENT {
		name := p.curToken.Literal
		p.nextToken()
		return &Spread{Name: name}
	}
	return &Spread{Name: ""}
}

func (p *Parser) parseCall() Expression {
	var fn Expression
	if p.curToken.Type == lexer.IDENT {
		fn = &Identifier{Value: p.curToken.Literal}
		p.nextToken()
	} else {
		fn = p.parseExpression()
	}
	if !p.expect(lexer.LPAREN) {
		// Get the code snippet for function calls
		snippet := ""
		if p.curToken.Line-1 < len(p.sourceLines) {
			snippet = p.sourceLines[p.curToken.Line-1]
		}
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expected ( after function name",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
			Snippet: snippet,
			Caret:   p.curToken.Column,
			Fix:     "Add ( to start function call",
		})
		return nil
	}
	args := []Expression{}
	partial := false
	for {
		if p.curToken.Type == lexer.RPAREN || p.curToken.Type == lexer.EOF {
			break
		}
		if p.curToken.Type == lexer.VARARG {
			if p.peekToken.Type == lexer.IDENT {
				spread := &Spread{Name: p.peekToken.Literal}
				args = append(args, spread)
				p.nextToken() // consume ...
				p.nextToken() // consume ident
			} else {
				spread := &Spread{Name: ""}
				args = append(args, spread)
				p.nextToken() // consume ...
			}
		} else {
			arg := p.parseExpression()
			if arg == nil {
				p.addError(utils.ParseError{
					Kind:    utils.InvalidSyntax,
					Message: "expected expression for call argument",
					Line:    p.curToken.Line,
					Column:  p.curToken.Column,
				})
				return nil
			}
			if ident, ok := arg.(*Identifier); ok && ident.Value == "_" {
				partial = true
			}
			args = append(args, arg)
		}
		if p.curToken.Type == lexer.COMMA {
			p.nextToken()
		} else {
			break
		}
	}
	if !p.expect(lexer.RPAREN) {
		// Get the code snippet for call arguments
		snippet := ""
		if p.curToken.Line-1 < len(p.sourceLines) {
			snippet = p.sourceLines[p.curToken.Line-1]
		}
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expected ) after call arguments",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
			Snippet: snippet,
			Caret:   p.curToken.Column,
			Fix:     "Add ) to close the function call",
		})
		return nil
	}
	if partial {
		return &PartialApplication{Function: fn, Args: args}
	}
	return &Call{Function: fn, Args: args}
}

func (p *Parser) parseArray() Expression {
  elems := []Expression{}
  if !p.expect(lexer.LBRACKET) {
    return nil
  }

  // Handle empty array: [ ]
  if p.curToken.Type == lexer.RBRACKET {
    p.nextToken()
    return &Array{Elements: elems}
  }

  for p.curToken.Type != lexer.RBRACKET && p.curToken.Type != lexer.EOF {
    elem := p.parseExpression()
    if elem == nil {
      p.addError(utils.ParseError{
        Kind:    utils.InvalidSyntax,
        Message: "expected expression for array element",
        Line:    p.curToken.Line,
        Column:  p.curToken.Column,
      })
      // Recovery: skip to next comma or closing bracket
      for p.curToken.Type != lexer.COMMA && p.curToken.Type != lexer.RBRACKET && p.curToken.Type != lexer.EOF {
        p.nextToken()
      }
      if p.curToken.Type == lexer.COMMA {
        p.nextToken()
        continue
      }
      break
    }
    elems = append(elems, elem)
    if p.curToken.Type == lexer.COMMA {
      p.nextToken()
    } else {
      break
    }
  }

  if !p.expect(lexer.RBRACKET) {
    return nil
  }
  return &Array{Elements: elems}
}

func (p *Parser) parseLambda() Expression {
	p.expect(lexer.LBRACE)
	stmts := p.parseStatementList(lexer.RBRACE)
	if !p.expect(lexer.RBRACE) {
		return &Block{Statements: []Statement{}}
	}
	return &Block{Statements: stmts}
}

func (p *Parser) parseBinaryExpr(minPrec int) Expression {
	left := p.parseUnary()
	for {
		prec, isOp := lexer.Precedences[p.curToken.Type]
		if isOp && prec >= minPrec {
			op := p.curToken.Type
			p.nextToken()
			right := p.parseBinaryExpr(prec + 1)
			if left == nil || right == nil {
				p.addError(utils.ParseError{Kind: utils.InvalidSyntax, Message: "nil in binary expression", Line: p.curToken.Line, Column: p.curToken.Column})
				return nil
			}
			left = &Call{Function: &Identifier{Value: parseLiteralForOperator(op)}, Args: []Expression{left, right}}
			continue
		}
		if p.curToken.Type == lexer.LPAREN {
			left = p.parseCallExpr(left)
			continue
		}
		if p.curToken.Type == lexer.DOT {
			p.nextToken()
			prop := &Identifier{Value: p.curToken.Literal}
			p.nextToken()
			left = &PropertyAccess{Object: left, Property: prop}
			continue
		}
		if p.curToken.Type == lexer.LBRACKET {
			p.nextToken()
			index := p.parseExpression()
			if index == nil {
				p.addError(utils.ParseError{
					Kind:    utils.InvalidSyntax,
					Message: "expected expression for array index",
					Line:    p.curToken.Line,
					Column:  p.curToken.Column,
				})
				return nil
			}
			if !p.expect(lexer.RBRACKET) {
				p.addError(utils.ParseError{
					Kind:    utils.InvalidSyntax,
					Message: "expected ] after array index",
					Line:    p.curToken.Line,
					Column:  p.curToken.Column,
				})
				return nil
			}
			left = &ArrayIndex{Array: left, Index: index}
			continue
		}
		break
	}
	return left
}

func (p *Parser) parseExpression() Expression {
	return p.parseBinaryExpr(0)
}

func (p *Parser) parseUnary() Expression {
	if p.curToken.Type == lexer.MINUS || p.curToken.Type == lexer.NOT_EQ {
		op := p.curToken.Type
		p.nextToken()
		right := p.parseUnary()
		if right == nil {
			p.addError(utils.ParseError{
				Kind:    utils.InvalidSyntax,
				Message: "expected expression after unary operator",
				Line:    p.curToken.Line,
				Column:  p.curToken.Column,
			})
			return nil
		}
		operator := "-"
		if op == lexer.NOT_EQ {
			operator = "!"
		}
		return &Call{
			Function: &Identifier{Value: operator},
			Args:     []Expression{right},
		}
	}
	// Handle spread operator in expressions
	if p.curToken.Type == lexer.VARARG {
		return p.parseSpread()
	}
	return p.parsePrimary()
}

func (p *Parser) parsePropertyAccess() *PropertyAccess {
	obj := &Identifier{Value: p.curToken.Literal}
	p.nextToken()
	if !p.expect(lexer.DOT) {
		return nil
	}
	prop := &Identifier{Value: p.curToken.Literal}
	p.nextToken()
	return &PropertyAccess{Object: obj, Property: prop}
}

func (p *Parser) parseStructInstantiation() *StructInstantiation {
	typeName := &Identifier{Value: p.curToken.Literal}
	p.nextToken()
	
	if !p.expect(lexer.LBRACE) {
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expected { after struct type name",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		return nil
	}
	
	fields := make(map[string]Expression)
	
	for p.curToken.Type != lexer.RBRACE && p.curToken.Type != lexer.EOF {
		if p.curToken.Type != lexer.IDENT {
			p.addError(utils.ParseError{
				Kind:    utils.InvalidSyntax,
				Message: "expected field name",
				Line:    p.curToken.Line,
				Column:  p.curToken.Column,
			})
			return nil
		}
		
		fieldName := p.curToken.Literal
		p.nextToken()
		
		if !p.expect(lexer.COLON) {
			p.addError(utils.ParseError{
				Kind:    utils.InvalidSyntax,
				Message: "expected : after field name",
				Line:    p.curToken.Line,
				Column:  p.curToken.Column,
			})
			return nil
		}
		
		fieldValue := p.parseExpression()
		if fieldValue == nil {
			p.addError(utils.ParseError{
				Kind:    utils.InvalidSyntax,
				Message: "expected expression for field value",
				Line:    p.curToken.Line,
				Column:  p.curToken.Column,
			})
			return nil
		}
		
		fields[fieldName] = fieldValue
		
		if p.curToken.Type == lexer.COMMA {
			p.nextToken()
		} else if p.curToken.Type != lexer.RBRACE {
			break
		}
	}
	
	if !p.expect(lexer.RBRACE) {
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expected } to close struct instantiation",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		return nil
	}
	return &StructInstantiation{
		TypeName: typeName,
		Fields:   fields,
	}
}

func (p *Parser) parseAnonymousStruct() *StructInstantiation {
	if !p.expect(lexer.LBRACE) {
		return nil
	}
	fields := make(map[string]Expression)
	// Handle empty struct: { }
	if p.curToken.Type == lexer.RBRACE {
		if !p.expect(lexer.RBRACE) {
			return nil
		}
		return &StructInstantiation{
			TypeName: nil,
			Fields:   fields,
		}
	}
	for p.curToken.Type != lexer.RBRACE && p.curToken.Type != lexer.EOF {
		if p.curToken.Type != lexer.IDENT {
			p.addError(utils.ParseError{
				Kind:    utils.InvalidSyntax,
				Message: "expected field name",
				Line:    p.curToken.Line,
				Column:  p.curToken.Column,
			})
			return nil
		}
		fieldName := p.curToken.Literal
		p.nextToken() // Move to colon!
		if !p.expect(lexer.COLON) {
			p.addError(utils.ParseError{
				Kind:    utils.InvalidSyntax,
				Message: "expected : after field name",
				Line:    p.curToken.Line,
				Column:  p.curToken.Column,
			})
			return nil
		}
		fieldValue := p.parseExpression()
		if fieldValue == nil {
			p.addError(utils.ParseError{
				Kind:    utils.InvalidSyntax,
				Message: "expected expression for field value",
				Line:    p.curToken.Line,
				Column:  p.curToken.Column,
			})
			return nil
		}
		fields[fieldName] = fieldValue
		if p.curToken.Type == lexer.COMMA {
			p.nextToken()
		} else if p.curToken.Type != lexer.RBRACE {
			break
		}
	}
	if !p.expect(lexer.RBRACE) {
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expected } to close struct literal",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		return nil
	}
	return &StructInstantiation{
		TypeName: nil,
		Fields:   fields,
	}
}
