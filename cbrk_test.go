package cbrk

import (
  "testing"
)


func TestCharEquality(t *testing.T) {
  var char Char
  char = LiteralChar{value: 'a'}
  if ! char.Equals("a") {
    t.Error("Char('a') != 'a'")
  }

  char = EscapeChar{sequence: "\x1b[A"}
  if ! char.Equals(Up.String()) {
    t.Error("Char('\\x1b[A') != Left")
  }
}
