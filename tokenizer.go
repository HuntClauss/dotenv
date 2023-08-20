package dotenv

const eof = 0
const (
	tokenInvalid = "INVALID"
	tokenComment = "COMMENT"
	tokenEqual   = "EQUAL"
	tokenLiteral = "LITERAL"
	tokenString  = "STRING"
	tokenEOF     = "EOF"
)

type token struct {
	Type  string
	Value string
}

func newToken(t, v string) token {
	return token{Type: t, Value: v}
}

type tokenizer struct {
	content string
	pos     int
}

func newTokenizer(content string) *tokenizer {
	return &tokenizer{content: content, pos: 0}
}

func (t *tokenizer) readAll() []token {
	var tokens []token
	for {
		tok := t.next()
		tokens = append(tokens, tok)
		if tok.Type == tokenEOF {
			break
		}
	}
	return tokens
}

func (t *tokenizer) next() token {
	t.skipWhitespaces()
	if t.pos >= len(t.content) {
		return newToken(tokenEOF, tokenEOF)
	}

	c := t.peakRune()

	switch c {
	case '#':
		_ = t.readRune()
		return newToken(tokenComment, t.ReadUntil('\n'))
	case '=':
		_ = t.readRune()
		return newToken(tokenEqual, tokenEqual)
	case '"', '\'':
		_ = t.readRune()
		return newToken(tokenString, t.ReadUntil(c))
	}

	if isAlphanumeric(c) {
		return t.readLiteral()
	}

	return newToken(tokenInvalid, "")
}

func (t *tokenizer) skipWhitespaces() {
	for c := t.peakRune(); isWhitespace(c); c = t.peakRune() {
		_ = t.readRune()
	}
}

func (t *tokenizer) readLiteral() token {
	pos := t.pos
	for c := t.peakRune(); isAlphanumeric(c); c = t.peakRune() {
		_ = t.readRune()
	}
	return newToken(tokenLiteral, t.content[pos:t.pos])
}

func (t *tokenizer) readRune() rune {
	if t.pos >= len(t.content) {
		return eof
	}

	c := t.content[t.pos]
	t.pos++
	return rune(c)
}

func (t *tokenizer) peakRune() rune {
	if t.pos >= len(t.content) {
		return eof
	}

	c := t.content[t.pos]
	return rune(c)
}

func (t *tokenizer) ReadUntil(r rune) string {
	pos := t.pos
	for {
		c := t.readRune()
		if c == '\\' {
			_ = t.readRune()
			continue
		}
		if c == eof || c == r {
			break
		}
	}
	return t.content[pos : t.pos-1]
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

func isAlphanumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_' ||
		(r >= '0' && r <= '9')
}
