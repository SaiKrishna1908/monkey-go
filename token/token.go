package token

type TokenType string

/*
A token produced by the Lexer.

example:
Token {Type: token.LET, "let"}
Token {Type: token.IDENT, "foobar"}
Token {Type: token.INT, "10"}
Token {Type: token.ASSIGN, "="}
*/
type Token struct {
	Type    TokenType
	Literal string
}

/*
Different Token Types supported by monkey
*/
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQL = "=="
	LTE = "<="
	GTE = ">="
	NE  = "!="
)

var keywords = map[string]TokenType{
	"let":    LET,
	"fn":     FUNCTION,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
