package sql

import "strings"

type Parser struct {
	buf string
	pos int
}

func NewParser(s string) Parser {
	return Parser{buf: s, pos: 0}
}

func isSpace(ch byte) bool {
	switch ch {
	case '\t', '\n', '\v', '\f', '\r', ' ':
		return true
	}
	return false
}
func isAlpha(ch byte) bool {
	return 'a' <= (ch|32) && (ch|32) <= 'z'
}
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
func isNameStart(ch byte) bool {
	return isAlpha(ch) || ch == '_'
}
func isNameContinue(ch byte) bool {
	return isAlpha(ch) || isDigit(ch) || ch == '_'
}
func isSeparator(ch byte) bool {
	return ch < 128 && !isNameContinue(ch)
}

func (p *Parser) skipSpaces() {
	for p.pos < len(p.buf) && isSpace(p.buf[p.pos]) {
		p.pos += 1
	}
}

func (p *Parser) tryKeyword(kw string) bool {
	p.skipSpaces()

	if p.pos+len(kw) > len(p.buf) {
		return false
	}
	if strings.EqualFold(p.buf[p.pos:p.pos+len(kw)], kw) && isSeparator(p.buf[p.pos+len(kw)]) {
		p.pos += len(kw)
		return true
	}

	return false
}

func (p *Parser) tryName() (string, bool) {
	p.skipSpaces()

	if !isNameStart(p.buf[p.pos]) {
		return "", false
	}

	start, pos := p.pos, p.pos
	for ; pos < len(p.buf) && isNameContinue(p.buf[pos]); pos++ {
	}

	p.pos = pos
	return p.buf[start:pos], true
}

func (p *Parser) isEnd() bool {
	p.skipSpaces()
	return p.pos >= len(p.buf)
}
