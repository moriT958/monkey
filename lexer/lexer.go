package lexer

import "monkey/token"

type Lexer struct {
	input   string
	pos     int  // 現在の位置
	readPos int  // 次の位置
	ch      byte // 現在の文字
}

func New(input string) *Lexer {
	l := new(Lexer)
	l.input = input
	l.readChar()
	return l
}

// トークンを取得
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case 0:
		tok = newToken(token.EOF, l.ch)
		tok.Literal = ""
	}

	l.readChar()
	return tok
}

func newToken(tt token.TokenType, l byte) token.Token {
	return token.Token{
		Type:    tt,
		Literal: string(l),
	}
}

// 次の文字を読み込む
func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos++
}
