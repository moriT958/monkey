package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 識別子・リテラル
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"

	// 演算子
	ASSIGN = "="
	PLUS   = "+"

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// キーワードか識別子かを判別するために使う
func LookupIdent(ident string) TokenType {
	if tType, ok := keywords[ident]; ok {
		return tType
	}
	return IDENT
}
