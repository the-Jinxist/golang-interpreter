package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	//Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456
	//Operators
	ASSIGN = "="
	PLUS   = "+"
	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	//Conditionals
	EXCLAMATION   = "!"
	MINUS         = "-"
	FORWARD_SLASH = "/"
	ASTERISK      = "*"
	GREATERTHAN   = ">"
	LESSTHAN      = "<"
	EQUALS_TO     = "=="
	NOT_EQUALS_TO = "!="

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
