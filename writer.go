// provide a class to be used from cbreakable terminal IO

package cbrk

import (
	"fmt"
	"os"
	"strings"
)

// cbrk.Stdout and cbrk.Stderr are instances of Writer than are tied to os.Stdout and os.Stderr
var (
	Stdout = Writer{out: os.Stdout}
	Stderr = Writer{out: os.Stderr}
)

// The type of all printers, including cbrk.Stdout
type Writer struct {
	out   *os.File
	lines int
}

// Print out a single element without a trailing newline
func (w *Writer) Print(elem interface{}) {
	toPrint := fmt.Sprint(elem)
	w.lines += strings.Count(toPrint, "\n")
	fmt.Fprint(w.out, toPrint)
}

// Print out a single element with a trailing newline
func (w *Writer) Println(elem interface{}) {
	toPrint := fmt.Sprint(elem)
	w.lines += strings.Count(toPrint, "\n") + 1
	fmt.Fprintln(w.out, toPrint)
}

// Clear all lines printed by this instance of Writer
func (w *Writer) Clear() {
	for i := 0; i < w.lines; i++ {
		w.Clearln()
	}
}

// Clear the entire terminal
func (w *Writer) ClearPage() {
	fmt.Fprint(w.out, Clear)
}

// Clear 1 line. Be sure to print a newline before running this
func (w *Writer) Clearln() {
	fmt.Fprint(w.out, UpLine)
	fmt.Fprint(w.out, ClearLine)
}
