package main

import (
	"testing"
)

func TestLexerNext(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []struct {
			token Token
			char  rune
		}
	}{
		{
			name:  "single character",
			input: "a",
			expected: []struct {
				token Token
				char  rune
			}{
				{Character, 'a'},
				{EOS, ' '},
			},
		},
		{
			name:  "characters",
			input: "hello",
			expected: []struct {
				token Token
				char  rune
			}{
				{Character, 'h'},
				{Character, 'e'},
				{Character, 'l'},
				{Character, 'l'},
				{Character, 'o'},
				{EOS, ' '},
			},
		},
		{
			name:  "expression",
			input: "ab+c{}[]())%",
			expected: []struct {
				token Token
				char  rune
			}{
				{Character, 'a'},
				{Character, 'b'},
				{PositiveClosure, '+'},
				{Character, 'c'},
				{OpenBrace, '{'},
				{ClosedBrace, '}'},
				{OpenBracket, '['},
				{ClosedBracket, ']'},
				{OpenParenthesis, '('},
				{ClosedParenthesis, ')'},
				{ClosedParenthesis, ')'},
				{EOS, ' '},
			},
		},
		{
			name:  "escaping",
			input: "%a.b%*",
			expected: []struct {
				token Token
				char  rune
			}{
				{Character, 'a'},
				{Concat, '.'},
				{Character, 'b'},
				{Character, '*'},
				{EOS, ' '},
			},
		},
		{
			name:  "multiple escape characters",
			input: "%%%%a",
			expected: []struct {
				token Token
				char  rune
			}{
				{Character, '%'},
				{Character, '%'},
				{Character, 'a'},
				{EOS, ' '},
			},
		},
		{
			name:  "empty",
			input: "",
			expected: []struct {
				token Token
				char  rune
			}{
				{EOS, ' '},
			},
		},
		{
			name:  "space trimming",
			input: "a   .%  ",
			expected: []struct {
				token Token
				char  rune
			}{
				{Character, 'a'},
				{Whitespace, ' '},
				{Concat, '.'},
				{Character, ' '},
				{Whitespace, ' '},
				{EOS, ' '},
			},
		},

		{
			name:  "semi-special characters",
			input: "1,  2%3-",
			expected: []struct {
				token Token
				char  rune
			}{
				{Digit, '1'},
				{Comma, ','},
				{Whitespace, ' '},
				{Digit, '2'},
				{Character, '3'},
				{Minus, '-'},
				{EOS, ' '},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lexer := NewLexer(tt.input)
			for i, expected := range tt.expected {
				token, char := lexer.Next()
				if token != expected.token {
					t.Errorf("Test %s: Expected token at index %d to be %v, got %v", tt.name, i, expected.token, token)
				}
				if char != expected.char {
					t.Errorf("Test %s: Expected char at index %d to be %c, got %c", tt.name, i, expected.char, char)
				}
			}
		})
	}
}
