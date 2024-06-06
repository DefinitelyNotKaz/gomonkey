package lexer

import (
	"gomonkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.char = 0
	} else {
		lexer.char = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	switch lexer.char {
	case '=':
		tok = newToken(token.ASSIGN, lexer.char)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.char)
	case '(':
		tok = newToken(token.OPEN_PARENTHESIS, lexer.char)
	case ')':
		tok = newToken(token.CLOSE_PARENTHESIS, lexer.char)
	case ',':
		tok = newToken(token.COMMA, lexer.char)
	case '+':
		tok = newToken(token.PLUS, lexer.char)
	case '{':
		tok = newToken(token.OPEN_CURLY, lexer.char)
	case '}':
		tok = newToken(token.CLOSE_CURLY, lexer.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lexer.char) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIndent(tok.Literal)
			return tok
		} else if isDigit(lexer.char) {
			tok.Type = token.INT
			tok.Literal = lexer.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.char)
		}
	}

	lexer.readChar()
	return tok
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.char) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.char) {
		lexer.readChar()
	}
	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.char == ' ' || lexer.char == '\t' || lexer.char == '\n' || lexer.char == '\r' {
		lexer.readChar()
	}
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}
