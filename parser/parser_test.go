package parser

import (
	"gomonkey/ast"
	"gomonkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `let x = 5;
						let y = 10;
						let foo = 69420;`

	lexer := lexer.New(input)
	parser := New(lexer)

	program := parser.ParseProgram()
	checkParserErorrs(t, parser)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Satements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentfier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i, tokenType := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, tokenType.expectedIdentfier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "let" {
		t.Errorf("statement.TokenLiteral not 'let'. got=%q", statement.TokenLiteral())
		return false
	}

	letStatement, ok := statement.(*ast.LetStatement)
	if !ok {
		t.Errorf("statement not *ast.LetStatement. got=%T", statement)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("letStatement.Name.Value not '%s'. got=%s", name, letStatement.Name)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("statement.Name not '%s'. got=%s", name, letStatement.Name)
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	input := `return 5;
						return 10;
						return 69420;`

	lexer := lexer.New(input)
	parser := New(lexer)

	program := parser.ParseProgram()
	checkParserErorrs(t, parser)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Satements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, statement := range program.Statements {
		returnStatement, ok := statement.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("statement not *ast.returnStatement. got=%T", statement)
			continue
		}
		if returnStatement.TokenLiteral() != "return" {
			t.Errorf("returnStatement.TokenLiteral not 'return', got %q",
				returnStatement.TokenLiteral())
		}
	}
}

func checkParserErorrs(t *testing.T, parser *Parser) {
	errors := parser.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
