package parser

import (
	"github.com/ckLXHL/gotoy/src/token"
	"github.com/ckLXHL/gotoy/src/expressions"
)

type PrefixParser interface {
	parse(p Parser, t token.Token) expressions.Expression
}

// prefixtype
type Prefix int
/*const (
	PLUS Prefix = iota //+
	MINUS //-
	TILDE //~
	BANG //~
)

type PrefixOperatorParselet struct {

}*/

/*func (operation PrefixOperatorParselet) parse(parser Parser, token Token) ExpressionIf {
	return new(ExpressionIf)
}*/