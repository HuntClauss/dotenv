package env

const EOF = 0
const (
	TokenInvalid = "INVALID"
	TokenComment = "COMMENT"
	TokenEqual   = "EQUAL"
	TokenLiteral = "LITERAL"
	TokenString  = "STRING"
	TokenEOF     = "EOF"
)

type Token struct {
	Type  string
	Value string
}

func NewToken(t, v string) Token {
	return Token{Type: t, Value: v}
}

type Tokenizer struct {
	content string
	pos     int
}

func NewTokenizer(content string) *Tokenizer {
	return &Tokenizer{content: content, pos: 0}
}

func (t *Tokenizer) ReadAll() []Token {
	var tokens []Token
	for {
		token := t.Next()
		tokens = append(tokens, token)
		if token.Type == TokenEOF {
			break
		}
	}
	return tokens
}

func (t *Tokenizer) Next() Token {
	t.skipWhitespaces()
	if t.pos >= len(t.content) {
		return NewToken(TokenEOF, TokenEOF)
	}

	c := t.PeakRune()

	switch c {
	case '#':
		_ = t.ReadRune()
		return NewToken(TokenComment, t.ReadUntil('\n'))
	case '=':
		_ = t.ReadRune()
		return NewToken(TokenEqual, TokenEqual)
	case '"', '\'':
		_ = t.ReadRune()
		return NewToken(TokenString, t.ReadUntil(c))
	}

	if isAlphanumeric(c) {
		return t.readLiteral()
	}

	return NewToken(TokenInvalid, "")
}

func (t *Tokenizer) skipWhitespaces() {
	for c := t.PeakRune(); isWhitespace(c); c = t.PeakRune() {
		_ = t.ReadRune()
	}
}

func (t *Tokenizer) readLiteral() Token {
	pos := t.pos
	for c := t.PeakRune(); isAlphanumeric(c); c = t.PeakRune() {
		_ = t.ReadRune()
	}
	return NewToken(TokenLiteral, t.content[pos:t.pos])
}

func (t *Tokenizer) ReadRune() rune {
	if t.pos >= len(t.content) {
		return EOF
	}

	c := t.content[t.pos]
	t.pos++
	return rune(c)
}

func (t *Tokenizer) PeakRune() rune {
	if t.pos >= len(t.content) {
		return EOF
	}

	c := t.content[t.pos]
	return rune(c)
}

func (t *Tokenizer) ReadUntil(r rune) string {
	pos := t.pos
	for {
		c := t.ReadRune()
		if c == '\\' {
			_ = t.ReadRune()
			continue
		}
		if c == EOF || c == r {
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
