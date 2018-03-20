package parser

import (
	"github.com/ckLXHL/gotoy/src/token"
)

type Parser struct {
	PrefixParselets map[token.TokenType] PrefixParser
	InfixParselets map[token.TokenType] InfixParser
}

func (p *Parser) registInfix(t token.TokenType, parse InfixParser) {
	p.InfixParselets[t] = parse
}

func (p *Parser) registPrefix(t token.TokenType, parse PrefixParser) {
	p.PrefixParselets[t] = parse
}