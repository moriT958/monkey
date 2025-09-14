package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	cases := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for _, c := range cases {
		tok := l.NextToken()
		t.Run("トークンタイプが正しく認識されている", func(t *testing.T) {
			if tok.Type != c.expectedType {
				t.Errorf("failed to recognize token type: expected %q got %q", c.expectedType, tok.Type)
			}
		})

		t.Run("トークンのリテラルが正しく認識されている", func(t *testing.T) {
			if tok.Literal != c.expectedLiteral {
				t.Errorf("failed to recognize token literal: expected %s got %s", c.expectedLiteral, tok.Literal)
			}
		})
	}
}
