// Provide a new Char type to be used to talk about single characters

package cbrk

import (
	"fmt"
	"strconv"
)

var (
	Up        = EscapeChar{sequence: "\x1b[A"}
	Down      = EscapeChar{sequence: "\x1b[B"}
	Right     = EscapeChar{sequence: "\x1b[C"}
	Left      = EscapeChar{sequence: "\x1b[D"}
	Clear     = EscapeChar{sequence: "\x1b[2J"}
	ClearLine = EscapeChar{sequence: "\x1b[2K"}
	Red       = EscapeChar{sequence: "\x1b[1;31m"}
	Blue      = EscapeChar{sequence: "\x1b[1;34m"}
	Cyan      = EscapeChar{sequence: "\x1b[1;36m"}
	Green     = EscapeChar{sequence: "\x1b[0;32m"}
	Reset     = EscapeChar{sequence: "\x1b[0;0m"}
	UpLine    = EscapeChar{sequence: "\x1b[1A"}

	validChar = map[string]Char{
		"[":  SpecialChar{code: Partial},
		"[A": Up,
		"[B": Down,
		"[C": Right,
		"[D": Left,
	}
)

const (
	Literal = iota
	Escape  = iota
	Special = iota
)

const (
	Partial  = iota
	Reserved = iota
	Illegal  = iota
)

// Any raw terminal input will yield a slice of cbrk.Char
type Char interface {
	String() string
	Equals(interface{}) bool
	Repr() string
	Kind() int
}

// The type used for chars that can be represented in literal form, ie 'a', '-' or ')'
type LiteralChar struct {
	value byte
}

func (l LiteralChar) String() string {
	return string(l.value)
}

func (l LiteralChar) Equals(other interface{}) bool {
	return l.String() == fmt.Sprint(other)
}

func (l LiteralChar) Repr() string {
	return strconv.Itoa(int(l.value))
}

func (l LiteralChar) Kind() int {
	return Literal
}

// The type used for chars that cannot be represented in literal form, ie '<Esc>', '\n' or '<CR>'
// This also includes sequences like '<Left>' and '<Clear>' that are actually printed as an escape sequence of multiple chars
type EscapeChar struct {
	sequence string
}

func (l EscapeChar) String() string {
	return l.sequence
}

func (l EscapeChar) Equals(other interface{}) bool {
	return fmt.Sprint(other) == l.String() || l.sequence == fmt.Sprint(other)
}

func (l EscapeChar) Repr() string {
	output := ""
	for _, i := range []byte(l.sequence) {
		output += strconv.Itoa(int(i)) + " "
	}
	return output
}

func (l EscapeChar) Kind() int {
	return Escape
}

type SpecialChar struct {
	code int16
}

func (s SpecialChar) String() string {
	return strconv.Itoa(int(s.code))
}

func (s SpecialChar) Equals(other interface{}) bool {
	return fmt.Sprint(other) == strconv.Itoa(int(s.code))
}

func (s SpecialChar) Repr() string {
	return s.String()
}

func (s SpecialChar) Kind() int {
	return Special
}

func CharString(chars []Char) string {
	output := []byte{}

	for _, i := range chars {
		output = append(output, byte(i.String()[0]))
	}

	return string(output)
}
