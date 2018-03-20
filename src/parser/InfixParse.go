package parser

import (
	"github.com/ckLXHL/gotoy/src/expressions"
	"github.com/ckLXHL/gotoy/src/token"
)

type InfixParser interface {
	parse(p Parser, t token.Token) expressions.Expression
}