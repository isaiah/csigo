// Part of the logic copied from https://github.com/benbjohnson/sql-parser
package main

import (
	"fmt"
	"io"
)

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// Parse parse git2 log
// git log --all -M -C --numstat --date=short --pretty=format:'--%h--%cd--%cn'
func (p *Parser) Parse() (*Prelude, error) {
	prelude := &Prelude{}

	// First token should be a "--"
	if tok, lit := p.scan(); tok != SEPARATOR {
		return nil, fmt.Errorf("found %q, expected SEPARATOR", lit)
	}

	// Read a field.
	tok, lit := p.scan()
	if tok != REV {
		return nil, fmt.Errorf("found %q, expected REV", lit)
	}
	prelude.Rev = lit

	tok, lit = p.scanIgnoreSeparator()
	if tok != DATE {
		return nil, fmt.Errorf("found %q, expected DATE", lit)
	}
	prelude.Rev = lit

	tok, lit = p.scanIgnoreSeparator()
	if tok != AUTHOR {
		return nil, fmt.Errorf("found %q, expected AUTHOR", lit)
	}
	prelude.Rev = lit

	// Return the successfully parsed statement.
	return prelude, nil
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

func (p *Parser) scanIgnoreSeparator() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == SEPARATOR {
		tok, lit = p.scan()
	}
	return
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }
