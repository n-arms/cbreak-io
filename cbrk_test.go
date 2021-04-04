package cbrk

import (
	"os"
	"testing"
)

func TestCharEquality(t *testing.T) {
	var char Char
	char = LiteralChar{value: 'a'}
	if !char.Equals("a") {
		t.Error("Char('a') != 'a'")
	}

	char = EscapeChar{sequence: "\x1b[A"}
	if !char.Equals(Up) {
		t.Error("Char('\\x1b[A') != Left")
	}
}

func TestWriterPrint(t *testing.T) {
	w := Writer{out: os.Stdout}
	w.Println("hello, world")
}
