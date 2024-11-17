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

type failure = int

const ( /* failure enum */
	blocklimit = iota
	charlimit
	inputlimit
	labellimit
	namelimit
	processlimit
	wordlimit
)

type errorkind = int

const ( /* errorkind enum */
	ambiguous3 = iota
	call3
	cobegin3
	constructor3
	funcval3
	incomplete3
	numeral3
	range3
	split3
	syntax3
	type3
	undeclared3
)

func fail(reason failure) {
	var msg string
	switch reason {
	case blocklimit:
		msg = "block limit"
	case charlimit:
		msg = "char limit"
	case inputlimit:
		msg = "input limit"
	case labellimit:
		msg = "label limit"
	case namelimit:
		msg = "name limit"
	case processlimit:
		msg = "process limit"
	case wordlimit:
		msg = "word limit"
	default:
		msg = "unknown failure"
	}
	fmt.Println(msg)
	os.Exit(1)
}

func note_error(lineno int, kind errorkind) {
	var msg string
	switch kind {
	case ambiguous3:
		msg = "ambiguous name"
	case call3:
		msg = "invalid procedure call"
	case cobegin3:
		msg = "invalid concurrent statement"
	case constructor3:
		msg = "invalid constructor"
	case funcval3:
		msg = "invalid use of function variable"
	case incomplete3:
		msg = "invalid recursive use of name"
	case numeral3:
		msg = "numeral out of range"
	case range3:
		msg = "invalid range"
	case split3:
		msg = "invalid split procedure"
	case syntax3:
		msg = "invalid syntax"
	case type3:
		msg = "invalid type"
	case undeclared3:
		msg = "undeclared name"
	default:
		msg = "unknown error"
	}
	fmt.Printf("%s:%d: %s\n", fn_in, lineno, msg)
	error_count++
}

type Checksums struct {
	sum1, sum2 int
}

func (c *Checksums) Update(v int16) {
	const n = 8191
	c.sum1 = (c.sum1 + (int(v) % n)) % n
	c.sum2 = (c.sum2 + c.sum1) % n
}

var error_count int
var fn_in, fn_out string

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: compile input.text output.code")
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

	showChecksums := false
	checksums := Checksums{}

	var out1 []int16
	pass1(
		func(value *rune) {
			ch, _, err := reader.ReadRune()
			if err != nil {
				if errors.Is(err, io.EOF) {
					ch = endmedium
				} else {
					fmt.Println(err)
					os.Exit(1)
				}
			}
			*value = ch
		},
		func(value symbol) {
			checksums.Update(int16(value))
			out1 = append(out1, int16(value))
		},
		fail,
	)
	if showChecksums {
		fmt.Println("Pass 1 checksum:", checksums)
	}

	var out2 []int16
	i := 0
	pass2(
		func(value *symbol) {
			*value = int(out1[i])
			i++
		},
		func(value symbol) {
			checksums.Update(int16(value))
			out2 = append(out2, int16(value))
		},
		fail,
	)
	if showChecksums {
		fmt.Println("Pass 2 checksum:", checksums)
	}

	var out3 []int16
	i = 0
	pass3(
		func(value *symbol) {
			*value = int(out2[i])
			i++
		},
		func(value symbol) {
			checksums.Update(int16(value))
			out3 = append(out3, int16(value))
		},
		fail,
	)
	if showChecksums {
		fmt.Println("Pass 3 checksum:", checksums)
	}

	var out4 []int16
	i = 0
	pass4(
		true,
		func(value *operator) {
			*value = int(out3[i])
			i++
		},
		func(value int) {
			checksums.Update(int16(value))
			out4 = append(out4, int16(value))
		},
		note_error,
		func() {
			i = 0
			out4 = nil
		},
		fail,
	)
	if showChecksums {
		fmt.Println("Pass 4 checksum:", checksums)
	}

	if error_count != 0 {
		os.Exit(1)
	}

	f_out, err := os.Create(fn_out)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f_out.Close()

	err = binary.Write(f_out, binary.LittleEndian, out4)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
