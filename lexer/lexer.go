package lexer

import (
	"monkey/token"
)

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

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peakChar() == '=' {
			chTmp := l.ch
			l.readChar()
			literal := string(chTmp) + string(l.ch)
			tok = token.Token{
				Type:    token.EQ,
				Literal: literal,
			}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peakChar() == '=' {
			chTmp := l.ch
			l.readChar()
			literal := string(chTmp) + string(l.ch)
			tok = token.Token{
				Type:    token.NOT_EQ,
				Literal: literal,
			}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
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
	default:
		// 上記以外の文字だった場合
		if isLetter(l.ch) {
			strLiteral := l.readIndent()
			t := token.Token{
				Type:    token.LookupIdent(strLiteral),
				Literal: strLiteral,
			}
			return t
		} else if isDigit(l.ch) {
			numLiteral := l.readNumber()
			t := token.Token{
				Type:    token.INT,
				Literal: numLiteral,
			}
			return t
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
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

// 次の文字を見る(readCharと違って内部ポインタを動かさない)
func (l *Lexer) peakChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	}
	return l.input[l.readPos]
}

// リテラルとして使える文字であるかどうかを判定
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// 識別子を読み込む
func (l *Lexer) readIndent() string {
	pos := l.pos
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	pos := l.pos
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}
