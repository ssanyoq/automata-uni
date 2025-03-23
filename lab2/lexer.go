package main

import "unicode/utf8"

type Lexer struct {
	input []rune

	pos int
}

func ToRunes(str string) []rune {
	result := []rune{}
	bytes := []byte(str)
	for i := 0; i < len(bytes); {
		r, size := utf8.DecodeRune(bytes[i:])
		result = append(result, r)
		i += size
	}
	return result
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input: ToRunes(input),
		pos:   0,
	}
}

func (l *Lexer) skipWhitespaces() {
	for l.input[l.pos] == ' ' {
		l.pos++
		if len(l.input) == l.pos {
			break
		}
	}
}

func (l *Lexer) Next() (Token, rune) {
	if l.pos == len(l.input) {
		return EOS, ' '
	}
	char := l.input[l.pos]
	t := GetToken(char)
	// Managing escape characters
	switch t {
	case Escape:
		l.pos++
		if l.pos == len(l.input) {
			return EOS, ' '
		}
		t = Character
		char = l.input[l.pos]
		l.pos++
	case Whitespace:
		l.skipWhitespaces()
	default:
		l.pos++
	}
	return t, char
}
