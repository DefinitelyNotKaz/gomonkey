package lexer

import (
	"gomonkey/token"
	"testing"
)

func TestSimpleSymbols(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.OPEN_PARENTHESIS, "("},
		{token.CLOSE_PARENTHESIS, ")"},
		{token.OPEN_CURLY, "{"},
		{token.CLOSE_CURLY, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	lexer := New(input)

	for i, tokenType := range tests {
		token := lexer.NextToken()

		if token.Type != tokenType.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tokenType.expectedType, token.Type)
		}

		if token.Literal != tokenType.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tokenType.expectedLiteral, token.Literal)
		}

	}
}

func TestBasicSyntaxTokens(t *testing.T) {
	input := `let five = 5; let ten = 10;

let add = fn(x, y) {

x + y; };

let result = add(five, ten); `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.OPEN_PARENTHESIS, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.CLOSE_PARENTHESIS, ")"},
		{token.OPEN_CURLY, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.CLOSE_CURLY, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.OPEN_PARENTHESIS, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.CLOSE_PARENTHESIS, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, tokenType := range tests {
		token := lexer.NextToken()

		if token.Type != tokenType.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tokenType.expectedType, token.Type)
		}

		if token.Literal != tokenType.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tokenType.expectedLiteral, token.Literal)
		}

	}
}

func TestExpandedSymbols(t *testing.T) {
	input := `!-/*5;
						5 < 10 > 5;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
	}

	lexer := New(input)

	for i, tokenType := range tests {
		token := lexer.NextToken()

		if token.Type != tokenType.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tokenType.expectedType, token.Type)
		}

		if token.Literal != tokenType.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tokenType.expectedLiteral, token.Literal)
		}

	}
}

func TestKeywords(t *testing.T) {
	input := `if (5 < 10) {
							return true
						} else {
							return false
						}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IF, "if"},
		{token.OPEN_PARENTHESIS, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.CLOSE_PARENTHESIS, ")"},
		{token.OPEN_CURLY, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.CLOSE_CURLY, "}"},
		{token.ELSE, "else"},
		{token.OPEN_CURLY, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.CLOSE_CURLY, "}"},
	}

	lexer := New(input)

	for i, tokenType := range tests {
		token := lexer.NextToken()

		if token.Type != tokenType.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tokenType.expectedType, token.Type)
		}

		if token.Literal != tokenType.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tokenType.expectedLiteral, token.Literal)
		}

	}

}

func TestDoubleSymbols(t *testing.T) {
	input := `10 == 10
						10 != 9`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INT, "10"},
		{token.EQUAL, "=="},
		{token.INT, "10"},
		{token.INT, "10"},
		{token.NOT_EQUAL, "!="},
		{token.INT, "9"},
	}

	lexer := New(input)

	for i, tokenType := range tests {
		token := lexer.NextToken()

		if token.Type != tokenType.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tokenType.expectedType, token.Type)
		}

		if token.Literal != tokenType.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tokenType.expectedLiteral, token.Literal)
		}

	}

}
