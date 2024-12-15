package lexer

import (
	"golang-interpreter/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readNextChar()
	return l
}

func (l *Lexer) readNextChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readNextChar()
			tok = token.Token{Type: token.EQUALS_TO, Literal: string(ch) + string(l.ch)}

		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readNextChar()
			tok = token.Token{Type: token.NOT_EQUALS_TO, Literal: string(ch) + string(l.ch)}

		} else {
			tok = newToken(token.EXCLAMATION, l.ch)
		}
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.FORWARD_SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '>':
		tok = newToken(token.GREATERTHAN, l.ch)
	case '<':
		tok = newToken(token.LESSTHAN, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isNumber(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readNextChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.position
	for isLetter(l.ch) {

		//this will keep reading until the l.ch isn't a letter anymore
		//l.ch keeps changing because l.readNextChar() is called which keeps updating the
		//current read position and ch in turn
		l.readNextChar()
	}
	return l.input[startPosition:l.position]
}

func (l *Lexer) readNumber() string {
	startPosition := l.position
	for isNumber(l.ch) {

		//this will keep reading until the l.ch isn't a number anymore
		//l.ch keeps changing because l.readNextChar() is called which keeps updating the
		//current read position and ch in turn
		l.readNextChar()
	}

	return l.input[startPosition:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readNextChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
