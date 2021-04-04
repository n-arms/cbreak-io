// Provide a new Char type to be used to talk about single characters

package cbrk

import (
	"fmt"
)

// Any raw terminal input will yield a slice of cbrk.Char
type Char interface {
	String() string
	Equals(fmt.Stringer) bool
}

// The type used for chars that can be represented in literal form, ie 'a', '-' or ')'
type LiteralChar struct {
	value byte
}

func (l *LiteralChar) String() string {
	return string(l.value)
}

func (l *LiteralChar) Equals(other fmt.Stringer) bool {
	return l.String() == other.String()
}

// The type used for chars that cannot be represented in literal form, ie '<Esc>', '\n' or '<CR>'
// This also includes sequences like '<Left>' and '<Clear>' that are actually printed as an escape sequence of multiple chars
type EscapeChar struct {
	sequence string
}

func (l *EscapeChar) String() string {
	return l.sequence
}

func (l *EscapeChar) Equals(other fmt.Stringer) bool {
	return other.String() == l.String() || l.sequence == other.String()
}
