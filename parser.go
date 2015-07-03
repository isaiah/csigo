// Part of the logic copied from https://github.com/benbjohnson/sql-parser
package main

import (
	"fmt"
	"io"
	"strconv"
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
func (p *Parser) Parse() (entries []Entry, err error) {
	for {
		ch := p.s.read()
		if ch == eof {
			break
		}
		p.s.unread()
		entry, err := p.entry()
		if err != nil {
			fmt.Println(err)
			break
		}
		entries = append(entries, *entry)
	}
	return
}

func (p *Parser) entry() (*Entry, error) {
	entry := &Entry{}
	prelude, err := p.prelude()
	if err != nil {
		return nil, err
	}
	entry.Prelude = prelude
	for {
		change, err := p.change()
		if err != nil {
			fmt.Println(err)
			break
		}
		entry.Changes = append(entry.Changes, *change)
	}
	return entry, nil
}

func (p *Parser) prelude() (*Prelude, error) {
	prelude := &Prelude{}

	// First token should be a revision
	tok, lit := p.scanIgnoreSeparator()
	// revision also match all numbers
	if tok != REV && tok != NUMSTAT {
		return nil, fmt.Errorf("found %q, expected REV", lit)
	}
	prelude.Rev = lit

	tok, lit = p.scanIgnoreSeparator()
	if tok != DATE {
		return nil, fmt.Errorf("found %q, expected DATE", lit)
	}
	prelude.Date = lit

	tok, lit = p.scanIgnoreSeparator()
	if tok != AUTHOR {
		return nil, fmt.Errorf("found %q, expected AUTHOR", lit)
	}
	prelude.Author = lit

	// Return the successfully parsed statement.
	return prelude, nil
}

func (p *Parser) change() (*Change, error) {
	change := &Change{}
	tok, lit := p.scanIgnoreWhitespace()
	if tok != NUMSTAT {
		return nil, fmt.Errorf("found %q, expected NUMSTAT", lit)
	}
	change.LocAdded = mustParseInt(lit)
	tok, lit = p.scanIgnoreWhitespace()
	if tok != NUMSTAT {
		return nil, fmt.Errorf("found %q, expected NUMSTAT", lit)
	}
	change.LocDeleted = mustParseInt(lit)
	tok, lit = p.scanIgnoreWhitespace()
	if tok != FILE {
		return nil, fmt.Errorf("found %q, expected FILE", lit)
	}
	change.Entry = lit

	return change, nil
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

func mustParseInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}
