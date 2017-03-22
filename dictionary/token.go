package dictionary

type TokenAttr struct {
	frequency float64
	pos       string
}

func (t TokenAttr) Frequency() float64 {
	return t.frequency
}

func (t TokenAttr) Pos() string {
	return t.pos
}

// Token represents a Chinese word with (optional) frequency and POS.
type Token struct {
	text      string
	TokenAttr
}

//Text returns token's text.
func (t Token) Text() string {
	return t.text
}

/*
// Frequency returns token's frequency.
func (t Token) Frequency() float64 {
	return t.frequency
}

// Pos returns token's POS.
func (t Token) Pos() string {
	return t.pos
}*/

func (t Token) Attr() TokenAttr {
	return t.TokenAttr
}

// NewToken creates a new token.
func NewToken(text string, frequency float64, pos string) Token {
	return Token{text: text, TokenAttr: TokenAttr{frequency: frequency, pos: pos}}
}
