// Provide a way to get cbreak! input easily

package cbrk

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

var Stdin = Reader{bufio.NewReader(os.Stdin)}

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

type Reader struct {
	in *bufio.Reader
}

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
		char2, err := r.in.ReadByte()
		if err != nil {
			panic(err)
		}
		return EscapeChar{sequence: "\x1b" + string(char) + string(char2)}
	}
	return LiteralChar{value: char}
}

func (r *Reader) Getln() []Char {
	var char Char = LiteralChar{value: ' '}
	output := []Char{}
	for !char.Equals("\n") {
		char = r.Get()
		output = append(output, char)
	}
	return output
}
