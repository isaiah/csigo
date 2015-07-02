// Copied from https://github.com/benbjohnson/sql-parser
package main

import (
	"bufio"
	"bytes"
	"io"
)

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	// If we see a digit then consume as a number.
	if isDash(ch) {
		s.unread()
		return s.scanSeparator()
	} else if isAlphaNum(ch) {
		s.unread()
		return s.scanPrelude()
	}

	// Otherwise read the individual character.
	switch ch {
	case eof:
		return EOF, ""
	}

	return ILLEGAL, string(ch)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

func (s *Scanner) scanSeparator() (tok Token, lit string) {
	if isDash(s.read()) {
		return SEPARATOR, "--"
	}
	s.unread()
	return
}

func isDash(ch rune) bool {
	return ch == '-'
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanPrelude() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isAlpha(ch) && !isNum(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	str := buf.String()
	switch {
	case rev.MatchString(str):
		return REV, str
	case author.MatchString(str):
		return AUTHOR, str
	case date.MatchString(str):
		return DATE, str
	case numstate.MatchString(str):
		return NUMSTAT, str
	case file.MatchString(str):
		return FILE, str
	}

	// Otherwise return as illegal
	return ILLEGAL, str
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == '\n' }

func isAlphaNum(ch rune) bool { return isAlpha(ch) || isNum(ch) }

// isLetter returns true if the rune is a letter.
func isAlpha(ch rune) bool { return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') }

// isDigit returns true if the rune is a digit.
func isNum(ch rune) bool { return (ch >= '0' && ch <= '9') }

// eof represents a marker rune for the end of the reader.
var eof = rune(0)
