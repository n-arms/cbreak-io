// Provide a new Char type to be used to talk about single characters

package cbrk

import (
)

var (
  Up = EscapeChar{sequence: "\x1b[A"}
  Down = EscapeChar{sequence: "\x1b[B"}
  Right = EscapeChar{sequence: "\x1b[C"}
  Left = EscapeChar{sequence: "\x1b[D"}
  Clear = EscapeChar{sequence: "\x1b[2J"}
  ClearLine = EscapeChar{sequence: "\x1b[2K"}
  Red = EscapeChar{sequence: "\x1b[1;31m"}
  Blue = EscapeChar{sequence: "\x1b[1;34m"}
  Cyan = EscapeChar{sequence: "\x1b[1;36m"}
  Green = EscapeChar{sequence: "\x1b[0;32m"}
  Reset = EscapeChar{sequence: "\x1b[0;0m"}
)

// Any raw terminal input will yield a slice of cbrk.Char
type Char interface {
	String() string
	Equals(string) bool
}

// The type used for chars that can be represented in literal form, ie 'a', '-' or ')'
type LiteralChar struct {
	value byte
}

func (l LiteralChar) String() string {
	return string(l.value)
}

func (l LiteralChar) Equals(other string) bool {
  return l.String() == other
}


// The type used for chars that cannot be represented in literal form, ie '<Esc>', '\n' or '<CR>'
// This also includes sequences like '<Left>' and '<Clear>' that are actually printed as an escape sequence of multiple chars
type EscapeChar struct {
	sequence string
}

func (l EscapeChar) String() string {
	return l.sequence
}

func (l EscapeChar) Equals(other string) bool {
	return other == l.String() || l.sequence == other
}

