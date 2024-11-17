package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
)

const endmedium = rune(25)

func fail(reason failure) {
	var msg string
	switch reason {
	case charlimit:
		msg = "char limit"
	case namelimit:
		msg = "word limit"
	default:
		msg = "unknown failure"
	}
	fmt.Println(msg)
	os.Exit(1)
}

var last_error_lineno int = -1

func note_error(lineno int, kind errorkind) {
	if lineno == last_error_lineno {
		return
	}
	last_error_lineno = lineno

	var msg string
	switch kind {
	case ambiguous3:
		msg = "ambiguous name"
	case declaration3:
		msg = "invalid declaration"
	case kind3:
		msg = "invalid name kind"
	case padding3:
		msg = "invalid padding"
	case range3:
		msg = "out of range"
	case syntax3:
		msg = "invalid syntax"
	case trap3:
		msg = "invalid trap"
	case undeclared3:
		msg = "undeclared name"
	default:
		msg = "unknown error"
	}
	fmt.Printf("%s:%d: %s\n", fn_in, lineno, msg)
	error_count++
}

var error_count int
var fn_in, fn_out string

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: assembler input.text output.code")
		os.Exit(1)
	}
	fn_in = os.Args[1]
	fn_out = os.Args[2]

	f_in, err := os.Open(fn_in)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f_in.Close()
	reader := bufio.NewReader(f_in)

	var out []uint16
	alva(
		func(value *rune) {
			ch, _, err := reader.ReadRune()
			if err != nil {
				if errors.Is(err, io.EOF) {
					ch = endmedium
				} else {
					fmt.Println(err)
					os.Exit(1)
				}
			// } else {
			// 	fmt.Printf("%c",ch)
			}
			*value = ch
		},
		func(value int) {
			out = append(out, uint16(value))
		},
		func() {
			if error_count != 0 {
				os.Exit(1)
			}
			_, err := f_in.Seek(0, 0)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			reader = bufio.NewReader(f_in)
		},
		note_error,
		fail,
	)

	if error_count != 0 {
		os.Exit(1)
	}

	f_out, err := os.Create(fn_out)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f_out.Close()

	err = binary.Write(f_out, binary.LittleEndian, out)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
