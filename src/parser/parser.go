package parser

import (
	"aether/lib/utils"
	"aether/src/lexer"
	"fmt"
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
	currentFile string
	isParsingMatch bool
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

func (p *Parser) SetFile(file string) {
	p.currentFile = file
}

func (p *Parser) addError(err utils.ParseError) {
	err.File = p.currentFile
	p.Errors.Add(err)
	p.recoverAfterError()
}

// SyncTokens returns the set of token types that are safe to recover to after a parse error.
func (p *Parser) SyncTokens() []lexer.TokenType {
	return []lexer.TokenType{
		lexer.EOF, lexer.RBRACE, lexer.RETURN, lexer.FUNCTION, lexer.STRUCT, lexer.IF, lexer.WHILE, lexer.REPEAT, lexer.IMPORT, lexer.FOR,
	}
}

func (p *Parser) recoverAfterError() {
	syncTokens := p.SyncTokens()
	for {
		for _, t := range syncTokens {
			if p.curToken.Type == t {
				return
			}
		}
		p.nextToken()
	}
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) expect(t lexer.TokenType) bool {
	if p.curToken.Type != t {
		// Get the code snippet
		snippet := ""
		if p.curToken.Line-1 < len(p.sourceLines) {
			snippet = p.sourceLines[p.curToken.Line-1]
		}

		p.addError(utils.ParseError{
			Kind:    utils.UnexpectedToken,
			Message: utils.FormatTokenError(t.String(), p.curToken.Type.String(), p.curToken.Literal),
			Line:    p.curToken.Line,
			Column:  p.curToken.Column,
			Snippet: snippet,
			Caret:   p.curToken.Column,
			Fix:     fmt.Sprintf("Add %s here", utils.UserFriendlyTokenName(t.String(), "")),
		})
		return false
	}
	p.nextToken()
	return true
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

func ProgramToAST(prog *Program) *ASTNode {
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
