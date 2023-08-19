package env

type Parser struct {
	Tokens []Token
	pos    int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{Tokens: tokens, pos: 0}
}

func (p *Parser) Parse() map[string]string {
	var result = make(map[string]string)
	for {
		token := p.read()
		if token.Type == TokenEOF {
			break
		}

		if token.Type == TokenComment {
			continue
		}

		switch token.Type {
		case TokenLiteral:
			assign, value := p.read(), p.read()
			if assign.Type != TokenEqual {
				panic("expected equal sign")
			}

			if value.Type != TokenString && value.Type != TokenLiteral {
				panic("expected string or literal")
			}

			result[token.Value] = value.Value
		}
	}
	return result
}

func (p *Parser) peak() Token {
	return p.Tokens[p.pos]
}

func (p *Parser) read() Token {
	t := p.peak()
	p.pos++
	return t
}
