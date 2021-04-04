// provide a class to be used from cbreakable terminal IO

package cbrk

import (
	"fmt"
	"os"
	"strings"
)

var (
  Stdout = Writer{out: os.Stdout}
  Stderr = Writer{out: os.Stderr}
)

type Writer struct {
	out   *os.File
	lines int
}

func (w *Writer) Print(elem interface{}) {
	toPrint := fmt.Sprint(elem)
	w.lines += strings.Count(toPrint, "\n")
	fmt.Fprint(w.out, toPrint)
}

func (w *Writer) Println(elem interface{}) {
	toPrint := fmt.Sprint(elem)
	w.lines += strings.Count(toPrint, "\n") + 1
	fmt.Fprintln(w.out, toPrint)
}

func (w *Writer) Clear() {
	fmt.Fprint(w.out, Clear)
}

func (w *Writer) Clearln() {
	fmt.Fprint(w.out, UpLine)
  fmt.Fprint(w.out, ClearLine)
}
