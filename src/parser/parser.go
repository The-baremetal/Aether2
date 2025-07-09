package parser

import (
	"aether/lib/utils"
	"aether/src/lexer"
	"strings"
)

type ParseError struct {
	Message string
	Line    int
	Column  int
}

type Parser struct {
	l           *lexer.Lexer
	curToken    lexer.Token
	peekToken   lexer.Token
	Errors      utils.ParseErrorList
	sourceLines []string
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	p.Errors = utils.ParseErrorList{}
	if l != nil {
		p.sourceLines = strings.Split(l.Input, "\n")
	}
	return p
}

func (p *Parser) addError(err utils.ParseError) {
	p.Errors.Add(err)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) expect(t lexer.TokenType) bool {
	if p.curToken.Type != t {
		p.addError(utils.ParseError{
			Kind:    utils.UnexpectedToken,
			Message: "unexpected token: expected " + t.String() + ", got " + p.curToken.Type.String(),
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		return false
	}
	p.nextToken()
	return true
}

func (p *Parser) parseStatementList(stop lexer.TokenType) []Statement {
	stmts := []Statement{}
	for p.curToken.Type != stop && p.curToken.Type != lexer.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			stmts = append(stmts, stmt)
		}
	}
	return stmts
}

func (p *Parser) parseStatement() Statement {
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
	case lexer.TRY:
		return p.parseTry()
	case lexer.FOR:
		return p.parseFor()
	case lexer.LBRACE:
		return p.parseBlock()
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
	case lexer.IDENT:
		if p.peekToken.Type == lexer.ASSIGN {
			return p.parseAssignment()
		}
		fallthrough
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
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expression is not a valid statement",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		return nil
	}
}

func (p *Parser) parseFunc() *Function {
	if !p.expect(lexer.FUNCTION) {
		return nil
	}
	name := &Identifier{Value: p.curToken.Literal}
	if !p.expect(lexer.IDENT) {
		return nil
	}
	params := []*Identifier{}
	if p.curToken.Type == lexer.LPAREN {
		p.nextToken()
		for p.curToken.Type != lexer.RPAREN && p.curToken.Type != lexer.EOF {
			param := &Identifier{Value: p.curToken.Literal}
			if !p.expect(lexer.IDENT) {
				return nil
			}
			if p.curToken.Type == lexer.COLON {
				p.expect(lexer.COLON)
				if p.curToken.Type == lexer.IDENT {
					param.Type = p.curToken.Literal
					p.expect(lexer.IDENT)
				}
			}
			params = append(params, param)
			if p.curToken.Type == lexer.COMMA {
				p.expect(lexer.COMMA)
			}
		}
		if !p.expect(lexer.RPAREN) {
			return nil
		}
	}
	body := p.parseBlock()
	return &Function{Name: name, Params: params, Body: body}
}

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
	fields := []*Identifier{}
	for p.curToken.Type != lexer.RBRACE && p.curToken.Type != lexer.EOF {
		field := &Identifier{Value: p.curToken.Literal}
		if !p.expect(lexer.IDENT) {
			return nil
		}
		if p.curToken.Type == lexer.COLON {
			p.expect(lexer.COLON)
			if p.curToken.Type == lexer.IDENT {
				field.Type = p.curToken.Literal
				p.expect(lexer.IDENT)
			}
		}
		fields = append(fields, field)
		if p.curToken.Type == lexer.COMMA {
			p.expect(lexer.COMMA)
		}
	}
	if !p.expect(lexer.RBRACE) {
		return nil
	}
	return &StructDef{Name: name, Fields: fields}
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
	name := &Identifier{Value: p.curToken.Literal}
	p.nextToken()
	var as *Identifier
	if p.curToken.Type == lexer.AS {
		p.nextToken()
		as = &Identifier{Value: p.curToken.Literal}
		p.nextToken()
	}
	return &Import{Name: name, As: as}
}

func (p *Parser) parseTry() *Block {
	if !p.expect(lexer.TRY) {
		return nil
	}
	return p.parseBlock()
}

func (p *Parser) parseAssignment() *Assignment {
	name := &Identifier{Value: p.curToken.Literal}
	p.nextToken()
	if !p.expect(lexer.ASSIGN) {
		return nil
	}
	value := p.parseExpression()
	if value == nil {
		p.addError(utils.ParseError{
			Kind:    utils.InvalidSyntax,
			Message: "expected expression for assignment value",
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
		})
		return nil
	}
	return &Assignment{Name: name, Value: value}
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
		return nil
	}
	args := []Expression{}
	partial := false
	if p.curToken.Type != lexer.RPAREN {
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
		for p.curToken.Type == lexer.COMMA {
			p.nextToken()
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
	}
	if !p.expect(lexer.RPAREN) {
		return nil
	}
	if partial {
		return &PartialApplication{Function: fn, Args: args}
	}
	return &Call{Function: fn, Args: args}
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

func (p *Parser) parseBlock() *Block {
	if !p.expect(lexer.LBRACE) {
		return nil
	}
	stmts := p.parseStatementList(lexer.RBRACE)
	if !p.expect(lexer.RBRACE) {
		return nil
	}
	return &Block{Statements: stmts}
}

func (p *Parser) parseExpression() Expression {
	return p.parseBinaryExpr(0)
}

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

func (p *Parser) parseBinaryExpr(minPrec int) Expression {
	left := p.parsePrimary()
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
		break
	}
	return left
}

func (p *Parser) parsePrimary() Expression {
	var expr Expression
	switch p.curToken.Type {
	case lexer.IDENT:
		expr = &Identifier{Value: p.curToken.Literal}
		p.nextToken()
	case lexer.NUMBER, lexer.STRING:
		expr = &Literal{Value: p.curToken.Literal}
		p.nextToken()
	case lexer.LBRACKET:
		expr = p.parseArray()
	case lexer.LBRACE:
		expr = p.parseLambda()
	case lexer.LPAREN:
		p.nextToken()
		expr = p.parseExpression()
		if !p.expect(lexer.RPAREN) {
			p.addError(utils.ParseError{Kind: utils.InvalidSyntax, Message: "expected )", Line: p.curToken.Line, Column: p.curToken.Column})
			return nil
		}
	default:
		p.nextToken()
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
			prop := &Identifier{Value: p.curToken.Literal}
			p.nextToken()
			expr = &PropertyAccess{Object: expr, Property: prop}
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
	if p.curToken.Type != lexer.RPAREN {
		arg := p.parseExpression()
		if arg == nil {
			p.addError(utils.ParseError{Kind: utils.InvalidSyntax, Message: "nil call argument", Line: p.curToken.Line, Column: p.curToken.Column})
			return nil
		}
		if ident, ok := arg.(*Identifier); ok && ident.Value == "_" {
			partial = true
		}
		args = append(args, arg)
		for p.curToken.Type == lexer.COMMA {
			p.nextToken()
			arg := p.parseExpression()
			if arg == nil {
				p.addError(utils.ParseError{Kind: utils.InvalidSyntax, Message: "nil call argument after comma", Line: p.curToken.Line, Column: p.curToken.Column})
				return nil
			}
			if ident, ok := arg.(*Identifier); ok && ident.Value == "_" {
				partial = true
			}
			args = append(args, arg)
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

func (p *Parser) parsePropertyAccessExpr(obj Expression) Expression {
	p.expect(lexer.DOT)
	prop := &Identifier{Value: p.curToken.Literal}
	p.nextToken()
	return &PropertyAccess{Object: obj, Property: prop}
}

func (p *Parser) parseArray() Expression {
	elems := []Expression{}
	p.expect(lexer.LBRACKET)
	if p.curToken.Type != lexer.RBRACKET {
		elem := p.parseExpression()
		if elem == nil {
			p.addError(utils.ParseError{
				Kind:    utils.InvalidSyntax,
				Message: "expected expression for array element",
				Line:    p.curToken.Line,
				Column:  p.curToken.Column,
			})
			return nil
		}
		elems = append(elems, elem)
		for p.curToken.Type == lexer.COMMA {
			p.nextToken()
			elem := p.parseExpression()
			if elem == nil {
				p.addError(utils.ParseError{
					Kind:    utils.InvalidSyntax,
					Message: "expected expression for array element",
					Line:    p.curToken.Line,
					Column:  p.curToken.Column,
				})
				return nil
			}
			elems = append(elems, elem)
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
		// Return an empty block instead of nil to avoid panics
		return &Block{Statements: []Statement{}}
	}
	return &Block{Statements: stmts}
}

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

func (p *Parser) Parse() *Program {
	program := &Program{}
	program.Statements = p.parseStatementList(lexer.EOF)
	return program
}

func (p *Parser) ParseAST() *ASTNode {
	prog := p.Parse()
	return programToASTNode(prog)
}

func programToASTNode(prog *Program) *ASTNode {
	stmts := make([]*ASTNode, 0, len(prog.Statements))
	for _, s := range prog.Statements {
		stmts = append(stmts, statementToASTNode(s))
	}
	return &ASTNode{
		NodeKind: TranslationUnitKind,
		Inner:    stmts,
	}
}

func statementToASTNode(s Statement) *ASTNode {
	switch stmt := s.(type) {
	case *Function:
		return functionToASTNode(stmt)
	case *Assignment:
		return assignmentToASTNode(stmt)
	case *StructDef:
		return structDefToASTNode(stmt)
	case *If:
		return ifToASTNode(stmt)
	case *While:
		return whileToASTNode(stmt)
	case *Repeat:
		return repeatToASTNode(stmt)
	case *Block:
		return blockToASTNode(stmt)
	case *Return:
		return returnToASTNode(stmt)
	case *Import:
		return importToASTNode(stmt)
	case *Call:
		return callToASTNode(stmt)
	case *PropertyAccess:
		return propertyAccessToASTNode(stmt)
	case *Identifier:
		return identifierToASTNode(stmt)
	case *Literal:
		return literalToASTNode(stmt)
	case *For:
		return forToASTNode(stmt)
	default:
		return nil
	}
}

func functionToASTNode(f *Function) *ASTNode {
	params := make([]*ASTNode, 0, len(f.Params))
	for _, p := range f.Params {
		params = append(params, identifierToASTNode(p))
	}
	return &ASTNode{
		NodeKind: FunctionDeclKind,
		Name:     f.Name.Value,
		Params:   params,
		Body:     blockToASTNode(f.Body),
	}
}

func assignmentToASTNode(a *Assignment) *ASTNode {
	return &ASTNode{
		NodeKind: AssignmentKind,
		Name:     a.Name.Value,
		Value:    expressionToASTNode(a.Value),
	}
}

func structDefToASTNode(s *StructDef) *ASTNode {
	fields := make([]*ASTNode, 0, len(s.Fields))
	for _, f := range s.Fields {
		fields = append(fields, identifierToASTNode(f))
	}
	return &ASTNode{
		NodeKind: StructDefKind,
		Name:     s.Name.Value,
		Params:   fields,
	}
}

func ifToASTNode(i *If) *ASTNode {
	return &ASTNode{
		NodeKind: IfKind,
		Left:     expressionToASTNode(i.Condition),
		Body:     blockToASTNode(i.Consequence),
		Right:    blockToASTNode(i.Alternative),
	}
}

func whileToASTNode(w *While) *ASTNode {
	return &ASTNode{
		NodeKind: WhileKind,
		Left:     expressionToASTNode(w.Condition),
		Body:     blockToASTNode(w.Body),
	}
}

func repeatToASTNode(r *Repeat) *ASTNode {
	return &ASTNode{
		NodeKind: RepeatKind,
		Left:     expressionToASTNode(r.Count),
		Body:     blockToASTNode(r.Body),
	}
}

func blockToASTNode(b *Block) *ASTNode {
	stmts := make([]*ASTNode, 0, len(b.Statements))
	for _, s := range b.Statements {
		stmts = append(stmts, statementToASTNode(s))
	}
	return &ASTNode{
		NodeKind: BlockKind,
		Inner:    stmts,
	}
}

func returnToASTNode(r *Return) *ASTNode {
	return &ASTNode{
		NodeKind: ReturnKind,
		Value:    expressionToASTNode(r.Value),
	}
}

func importToASTNode(i *Import) *ASTNode {
	return &ASTNode{
		NodeKind: ImportKind,
		Name:     i.Name.Value,
		Value:    i.As.Value,
	}
}

func callToASTNode(c *Call) *ASTNode {
	args := make([]*ASTNode, 0, len(c.Args))
	for _, a := range c.Args {
		args = append(args, expressionToASTNode(a))
	}
	return &ASTNode{
		NodeKind: CallKind,
		Left:     expressionToASTNode(c.Function),
		Inner:    args,
	}
}

func propertyAccessToASTNode(p *PropertyAccess) *ASTNode {
	return &ASTNode{
		NodeKind: PropertyAccessKind,
		Left:     expressionToASTNode(p.Object),
		Name:     p.Property.Value,
	}
}

func identifierToASTNode(i *Identifier) *ASTNode {
	return &ASTNode{
		NodeKind: IdentifierKind,
		Name:     i.Value,
		Value:    i.Type,
	}
}

func literalToASTNode(l *Literal) *ASTNode {
	return &ASTNode{
		NodeKind: LiteralKind,
		Value:    l.Value,
	}
}

func arrayToASTNode(a *Array) *ASTNode {
	elems := make([]*ASTNode, 0, len(a.Elements))
	for _, e := range a.Elements {
		elems = append(elems, expressionToASTNode(e))
	}
	return &ASTNode{
		NodeKind: ArrayKind,
		Inner:    elems,
	}
}

func forToASTNode(f *For) *ASTNode {
	params := []*ASTNode{}
	if f.Index != nil {
		params = append(params, identifierToASTNode(f.Index))
	}
	if f.Value != nil {
		params = append(params, identifierToASTNode(f.Value))
	}
	return &ASTNode{
		NodeKind: ForKind,
		Params:   params,
		Left:     expressionToASTNode(f.Iterable),
		Body:     blockToASTNode(f.Body),
	}
}
