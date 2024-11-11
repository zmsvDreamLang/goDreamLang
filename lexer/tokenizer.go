package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

var (
	hexChar      = make(map[byte]bool)
	specialChar  = make(map[byte]bool)
	reservedChar = make(map[byte]bool)
)

func init() {
	for i := '0'; i <= '9'; i++ {
		hexChar[byte(i)] = true
	}
	for i := 'a'; i <= 'f'; i++ {
		hexChar[byte(i)] = true
	}
	for i := 'A'; i <= 'F'; i++ {
		hexChar[byte(i)] = true
	}

	specialChars := "+-*/^%<>=!&|()[]{}.,;:\n\"'`"
	for i := 0; i < len(specialChars); i++ {
		specialChar[specialChars[i]] = true
	}

	reservedChars := "#@$"
	for i := 0; i < len(reservedChars); i++ {
		reservedChar[reservedChars[i]] = true
	}
}

type Lexer struct {
	input []byte
	index int
	line  int
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input: []byte(input),
		index: 0,
		line:  1,
	}
}

func (l *Lexer) nextChar() *byte {
	if l.index+1 < len(l.input) {
		return &l.input[l.index+1]
	}
	return nil
}

func (l *Lexer) checkNumber(i int) (string, error) {
	isHex := false
	isBin := false
	isOct := false
	isE := false
	isDec := false
	dot := 0
	var sb strings.Builder
	c := l.input[i]
	if c == '-' {
		sb.WriteByte(c)
		i++
	} else if !unicode.IsDigit(rune(c)) {
		return "", fmt.Errorf("not number at line %d", l.line)
	}
	sb.WriteByte(c)
	i++
	for i < len(l.input) && !(unicode.IsSpace(rune(l.input[i])) || specialChar[l.input[i]]) {
		c = l.input[i]
		if unicode.IsDigit(rune(c)) {
			sb.WriteByte(c)
		} else if c == '.' && dot == 0 {
			sb.WriteByte(c)
			dot++
		} else if c == 'x' && sb.Len() == 1 && sb.String()[0] == '0' {
			sb.WriteByte(c)
			isHex = true
		} else if c == 'b' && sb.Len() == 1 && sb.String()[0] == '0' {
			sb.WriteByte(c)
			isBin = true
		} else if c == 'd' && sb.Len() == 1 && sb.String()[0] == '0' {
			sb.WriteByte(c)
			isDec = true
		} else if c != 'b' && c != 'd' && sb.Len() == 1 && sb.String()[0] == '0' {
			sb.WriteByte(c)
			isOct = true
		} else if !isE && (c == 'e' || c == 'E') {
			sb.WriteByte(c)
			isE = true
		} else if isHex && hexChar[c] {
			sb.WriteByte(c)
		} else if isBin && (c == '0' || c == '1') {
			sb.WriteByte(c)
		} else if isOct && (c >= '0' && c <= '7') {
			sb.WriteByte(c)
		} else if isDec && (c >= '0' && c <= '9') {
			sb.WriteByte(c)
		} else {
			return "", fmt.Errorf("wrong number token at line %d", l.line)
		}
		i++
	}
	return sb.String(), nil
}

func (l *Lexer) checkIdent(i int) (string, error) {
	var sb strings.Builder
	c := l.input[i]
	if reservedChar[c] {
		return "", fmt.Errorf("syntax error at line %d", l.line)
	}
	sb.WriteByte(c)
	i++
	for i < len(l.input) && !(unicode.IsSpace(rune(l.input[i])) || specialChar[l.input[i]]) {
		c = l.input[i]
		sb.WriteByte(c)
		i++
	}
	return sb.String(), nil
}

func (l *Lexer) checkExgesis(i int) string {
	var sb strings.Builder
	isMoreEx := false
	if i+1 < len(l.input) && l.input[i+1] == '*' && l.input[i] == '/' {
		isMoreEx = true
		sb.WriteString("/*")
	} else {
		sb.WriteString("//")
	}
	i += 2
	for i < len(l.input) {
		c := l.input[i]
		if c == '\n' && !isMoreEx {
			break
		} else if isMoreEx && i+1 < len(l.input) && c == '*' && l.input[i+1] == '/' {
			sb.WriteString("*/")
			break
		}
		sb.WriteByte(c)
		i++
	}
	return sb.String()
}
