package dotenv

type Parser struct {
	Tokens []token
	pos    int
}

func newParser(tokens []token) *Parser {
	return &Parser{Tokens: tokens, pos: 0}
}

func (p *Parser) parse() map[string]string {
	var result = make(map[string]string)
	for {
		tok := p.read()
		if tok.Type == tokenEOF {
			break
		}

		if tok.Type == tokenComment {
			continue
		}

		switch tok.Type {
		case tokenLiteral:
			assign, value := p.read(), p.read()
			if assign.Type != tokenEqual {
				panic("expected equal sign")
			}

			if value.Type != tokenString && value.Type != tokenLiteral {
				panic("expected string or literal")
			}

			result[tok.Value] = value.Value
		}
	}
	return result
}

func (p *Parser) peak() token {
	return p.Tokens[p.pos]
}

func (p *Parser) read() token {
	t := p.peak()
	p.pos++
	return t
}
