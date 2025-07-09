package lexer

type Lexer struct {
	Input        string
	Position     int
	ReadPosition int
	Ch           byte
	Line         int
	Column       int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{Input: input, Line: 1, Column: 0}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.ReadPosition >= len(l.Input) {
		l.Ch = 0
	} else {
		l.Ch = l.Input[l.ReadPosition]
	}
	l.Position = l.ReadPosition
	l.ReadPosition++
	if l.Ch == '\n' {
		l.Line++
		l.Column = 0
	} else {
		l.Column++
	}
}

func (l *Lexer) peekChar() byte {
	if l.ReadPosition >= len(l.Input) {
		return 0
	}
	return l.Input[l.ReadPosition]
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespaceAndComments()
	tok := Token{Line: l.Line, Column: l.Column}
	switch l.Ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = l.newToken(EQ, "==")
		} else {
			tok = l.newToken(ASSIGN, string(l.Ch))
		}
	case '+':
		tok = l.newToken(PLUS, string(l.Ch))
	case '-':
		tok = l.newToken(MINUS, string(l.Ch))
	case '*':
		tok = l.newToken(ASTERISK, string(l.Ch))
	case '/':
		if l.peekChar() == '/' {
			l.skipLineComment()
			return l.NextToken()
		} else if l.peekChar() == '*' {
			l.skipBlockComment()
			return l.NextToken()
		} else {
			tok = l.newToken(SLASH, string(l.Ch))
		}
	case '%':
		tok = l.newToken(MODULO, string(l.Ch))
	case '^':
		tok = l.newToken(EXPONENT, string(l.Ch))
	case '<':
		if l.peekChar() == '=' {
			l.readChar()
			tok = l.newToken(LE, "<=")
		} else {
			tok = l.newToken(LT, string(l.Ch))
		}
	case '>':
		if l.peekChar() == '=' {
			l.readChar()
			tok = l.newToken(GE, ">=")
		} else {
			tok = l.newToken(GT, string(l.Ch))
		}
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = l.newToken(NOT_EQ, "!=")
		} else {
			tok = l.newToken(ILLEGAL, string(l.Ch))
		}
	case ',':
		tok = l.newToken(COMMA, string(l.Ch))
	case ':':
		tok = l.newToken(COLON, string(l.Ch))
	case '[':
		tok = l.newToken(LBRACKET, string(l.Ch))
	case ']':
		tok = l.newToken(RBRACKET, string(l.Ch))
	case '.':
		if l.peekChar() == '.' {
			l.readChar()
			if l.peekChar() == '.' {
				l.readChar()
				tok = l.newToken(VARARG, "...")
			} else {
				tok = l.newToken(CONCAT, "..")
			}
		} else {
			tok = l.newToken(DOT, string(l.Ch))
		}
	case '(':
		tok = l.newToken(LPAREN, string(l.Ch))
	case ')':
		tok = l.newToken(RPAREN, string(l.Ch))
	case '{':
		tok = l.newToken(LBRACE, string(l.Ch))
	case '}':
		tok = l.newToken(RBRACE, string(l.Ch))
	case '"':
		tok.Type = STRING
		tok.Literal = l.readString()
		return tok
	case 0:
		tok.Type = EOF
		tok.Literal = ""
	default:
		if isLetter(l.Ch) {
			ident := l.readIdentifier()
			if ident == "as" {
				tok.Type = AS
				tok.Literal = ident
			} else if KEYWORDS[ident] != "" {
				tok.Type = KEYWORDS[ident]
				tok.Literal = ident
			} else {
				tok.Type = IDENT
				tok.Literal = ident
			}
			return tok
		} else if isDigit(l.Ch) {
			tok.Type = NUMBER
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = l.newToken(ILLEGAL, string(l.Ch))
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) newToken(tokenType TokenType, ch string) Token {
	return Token{Type: tokenType, Literal: ch, Line: l.Line, Column: l.Column}
}

func (l *Lexer) skipWhitespaceAndComments() {
	for {
		if l.Ch == ' ' || l.Ch == '\t' || l.Ch == '\r' || l.Ch == '\n' {
			l.readChar()
		} else if l.Ch == '/' && l.peekChar() == '/' {
			l.skipLineComment()
		} else if l.Ch == '/' && l.peekChar() == '*' {
			l.skipBlockComment()
		} else if l.Ch == '/' && l.peekChar() == '/' && l.peekAhead(2) == '/' {
			l.skipDocComment()
		} else {
			break
		}
	}
}

func (l *Lexer) skipLineComment() {
	for l.Ch != '\n' && l.Ch != 0 {
		l.readChar()
	}
}

func (l *Lexer) skipBlockComment() {
	l.readChar()
	l.readChar()
	for {
		if l.Ch == 0 {
			break
		}
		if l.Ch == '*' && l.peekChar() == '/' {
			l.readChar()
			l.readChar()
			break
		}
		l.readChar()
	}
}

func (l *Lexer) skipDocComment() {
	for l.Ch != '\n' && l.Ch != 0 {
		l.readChar()
	}
}

func (l *Lexer) peekAhead(n int) byte {
	pos := l.ReadPosition + n - 1
	if pos >= len(l.Input) {
		return 0
	}
	return l.Input[pos]
}

func (l *Lexer) readIdentifier() string {
	pos := l.Position
	for isLetter(l.Ch) || isDigit(l.Ch) || l.Ch == '_' {
		l.readChar()
	}
	return l.Input[pos:l.Position]
}

func (l *Lexer) readNumber() string {
	pos := l.Position
	for isDigit(l.Ch) {
		l.readChar()
	}
	return l.Input[pos:l.Position]
}

func (l *Lexer) readString() string {
	l.readChar()
	pos := l.Position
	for l.Ch != '"' && l.Ch != 0 {
		l.readChar()
	}
	s := l.Input[pos:l.Position]
	l.readChar()
	return s
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) Tokenize() []Token {
	tokens := []Token{}
	for {
		tok := l.NextToken()
		tokens = append(tokens, tok)
		if tok.Type == EOF {
			break
		}
	}
	return tokens
}
