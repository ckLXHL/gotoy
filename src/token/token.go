package token

type TokenType uint

const (
	LEST_PAREN = TokenType(iota)
	RIGHT_PAREN
	ASSIGNMENG
	PLUS
	MINUS
	ASTERISK
	SLASH
	TILDE
	BANG
)


var Punctuator map[TokenType]string = map[TokenType]string{
	LEST_PAREN: "(",
	RIGHT_PAREN: ")",
	ASSIGNMENG: "=",
	PLUS: "+",
	MINUS: "-",
	ASTERISK: "*",
	SLASH: "/",
	TILDE: "~",
	BANG: "!",
}
type Token struct {
	Type TokenType
}

func (t Token) String() string {
	if v, ok := Punctuator[t.Type]; ok {
		return v
	}
	panic("no support tokenType")
}
