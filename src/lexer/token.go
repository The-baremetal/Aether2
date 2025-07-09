package lexer

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

const (
	ILLEGAL   TokenType = "ILLEGAL"
	EOF       TokenType = "EOF"
	IDENT     TokenType = "IDENT"
	NUMBER    TokenType = "NUMBER"
	STRING    TokenType = "STRING"
	ASSIGN    TokenType = "ASSIGN"
	PLUS      TokenType = "PLUS"
	MINUS     TokenType = "MINUS"
	ASTERISK  TokenType = "ASTERISK"
	SLASH     TokenType = "SLASH"
	LT        TokenType = "LT"
	GT        TokenType = "GT"
	EQ        TokenType = "EQ"
	NOT_EQ    TokenType = "NOT_EQ"
	LE        TokenType = "LE"
	GE        TokenType = "GE"
	DIVIDE    TokenType = "DIVIDE"
	MULTIPLY  TokenType = "MULTIPLY"
	MODULO    TokenType = "MODULO"
	EXPONENT  TokenType = "EXPONENT"
	COMMA     TokenType = "COMMA"
	COLON     TokenType = "COLON"
	DOT       TokenType = "DOT"
	CONCAT    TokenType = "CONCAT"
	VARARG    TokenType = "VARARG"
	LPAREN    TokenType = "LPAREN"
	RPAREN    TokenType = "RPAREN"
	LBRACE    TokenType = "LBRACE"
	RBRACE    TokenType = "RBRACE"
	LBRACKET  TokenType = "LBRACKET"
	RBRACKET  TokenType = "RBRACKET"
	FUNCTION  TokenType = "FUNCTION"
	STRUCT    TokenType = "STRUCT"
	IF        TokenType = "IF"
	ELSE      TokenType = "ELSE"
	REPEAT    TokenType = "REPEAT"
	WHILE     TokenType = "WHILE"
	PRINT     TokenType = "PRINT"
	RETURN    TokenType = "RETURN"
	IMPORT    TokenType = "IMPORT"
	TRY       TokenType = "TRY"
	CATCH     TokenType = "CATCH"
	FINALLY   TokenType = "FINALLY"
	SPAWN     TokenType = "SPAWN"
	RECEIVE   TokenType = "RECEIVE"
	SEND      TokenType = "SEND"
	YIELD     TokenType = "YIELD"
	COPY      TokenType = "COPY"
	CASE      TokenType = "CASE"
	MATCH     TokenType = "MATCH"
	AS        TokenType = "AS"
	COMMENT   TokenType = "COMMENT"
	C_COMMENT TokenType = "C_COMMENT"
	KEYWORD   TokenType = "KEYWORD"
	PERCENT   TokenType = "PERCENT"
	CARET     TokenType = "CARET"
	NEQ       TokenType = "NEQ"
	LTE       TokenType = "LTE"
	GTE       TokenType = "GTE"
	IN        TokenType = "IN"
	FOR       TokenType = "FOR"
)

var KEYWORDS = map[string]TokenType{
	"func":    FUNCTION,
	"struct":  STRUCT,
	"if":      IF,
	"else":    ELSE,
	"repeat":  REPEAT,
	"while":   WHILE,
	"return":  RETURN,
	"import":  IMPORT,
	"try":     TRY,
	"catch":   CATCH,
	"finally": FINALLY,
	"spawn":   SPAWN,
	"receive": RECEIVE,
	"send":    SEND,
	"yield":   YIELD,
	"copy":    COPY,
	"case":    CASE,
	"match":   MATCH,
	"in":      IN,
	"for":     FOR,
}

const (
	LOWEST      = 1
	EQUALS      = 2
	LESSGREATER = 3
	SUM         = 4
	PRODUCT     = 5
	PREFIX      = 6
	CALL        = 7
)

var Precedences = map[TokenType]int{
	EQ:       EQUALS,
	NOT_EQ:   EQUALS,
	LT:       LESSGREATER,
	GT:       LESSGREATER,
	LE:       LESSGREATER,
	GE:       LESSGREATER,
	PLUS:     SUM,
	MINUS:    SUM,
	DIVIDE:   PRODUCT,
	MULTIPLY: PRODUCT,
	MODULO:   PRODUCT,
	EXPONENT: PRODUCT,
	CONCAT:   SUM,
}

var tokenTypeToString = map[TokenType]string{
	PERCENT: "%",
	CARET:   "^",
	NEQ:     "!=",
	LTE:     "<=",
	GTE:     ">=",
}

func (t TokenType) String() string {
	if s, ok := tokenTypeToString[t]; ok {
		return s
	}
	return string(t)
}
