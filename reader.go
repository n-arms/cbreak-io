// Provide a way to get cbreak! input easily

package cbrk

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

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

func (r *Reader) getch() Char {
	char, err := r.in.ReadByte()
	if err != nil {
		panic(err)
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
