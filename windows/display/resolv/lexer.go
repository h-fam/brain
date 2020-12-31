package resolv

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS

	// Literals
	IDENT // fields, table_name

	// Misc characters
	COMMENT // #

	// Keywords
	NAMESERVER
)

var eof = rune(0)

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}
func isSpace(ch rune) bool {
	return ch == ' ' || ch == '\t'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
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

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	}

	// Otherwise read the individual character.
	switch ch {
	case eof:
		return EOF, ""
	case '#':
		v := s.lexComment()
		return COMMENT, v
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

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	switch strings.ToUpper(buf.String()) {
	case "NAMESERVER":
		v := s.lexNameserver()
		return NAMESERVER, v
	}

	// Otherwise return as a regular identifier.
	return IDENT, buf.String()
}

func (s *Scanner) lexComment() string {
	var buf bytes.Buffer
	for {
		ch := s.read()
		if ch == eof {
			break
		}
		if ch == '\n' {
			break
		}
		buf.WriteRune(ch)
	}
	return buf.String()
}

// lexNameser will continue all characters until end of line.
func (s *Scanner) lexNameserver() string {
	var buf bytes.Buffer
	for {
		ch := s.read()
		if ch == eof {
			return ""
		}
		if !isSpace(ch) {
			s.unread()
			break
		}
	}
	for {
		ch := s.read()
		if ch == eof {
			break
		}
		if ch == '\n' {
			break
		}
		buf.WriteRune(ch)
	}
	return buf.String()
}
