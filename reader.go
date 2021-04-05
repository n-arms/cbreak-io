// Provide a way to get cbreak! input easily

package cbrk

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// A cbrk.Reader object tied to os.Stdin
var Stdin = Reader{bufio.NewReader(os.Stdin)}

// Call Cbreak(true) to enter cbreak mode, and call Cbreak(false) to exit
func Cbreak(on bool) {
	var err error
	var output []byte
	var command *exec.Cmd

	if !on {
		command = exec.Command("stty", "sane")
		command.Stdin = os.Stdin
		err = command.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}

	command = exec.Command("stty", "cbreak")
	command.Stdin = os.Stdin
	err = command.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	command = exec.Command("stty", "-echo")
	command.Stdin = os.Stdin
	err = command.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		panic(err)
	}
}

// The type of cbrk.Stdin
type Reader struct {
	in *bufio.Reader
}

// Get a single cbrk.Char, block until present
// If it recieves \x1b it will block until it recieves one of the programmed valid escape codes.
func (r *Reader) Get() Char {
	char, err := r.in.ReadByte()
	if err != nil {
		fmt.Println(err)
		return LiteralChar{}
	}

	if char == '\x1b' {
		char, err = r.in.ReadByte()
		if err != nil {
			panic(err)
		}
		output := []byte{char}
		for validChar[string(output)].Kind() == Special && validChar[string(output)].String() == strconv.Itoa(Partial) {
			char, err = r.in.ReadByte()
			if err != nil {
				panic(err)
			}
			output = append(output, char)
		}

		return EscapeChar{sequence: "\x1b" + string(output)}
	}
	return LiteralChar{value: char}
}

// Block until the given reader has a \n in it and return the entire line, including the newline 
// Using cbrk.Reader.Get
func (r *Reader) Getln() []Char {
	var char Char = LiteralChar{value: ' '}
	output := []Char{}
	for !char.Equals("\n") {
		char = r.Get()
		output = append(output, char)
	}
	return output
}
