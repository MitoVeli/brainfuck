package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

const SIZE int = 32768

func main() {

	content := strings.NewReader("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.")

	bf, err := NewBrainfuckService()
	if err != nil {
		fmt.Println("Error while creating brainfuck service")
	}

	if err = bf.Run(content); err != nil {
		fmt.Println("Error while running brainfuck service", err)
	}

}

type brainfuckService interface {
	Run(content io.Reader) error
	Read(content io.Reader) error
	Print()
	Plus()
	Minus()
	MoreThan()
	LessThan()
	Coma()
	LoopStart()
	LoopEnd()
}

type brainfuck struct {
	stack   []int
	size    int
	text    string
	index   int
	pointer int
	str     string
}

func NewBrainfuckService() (brainfuckService, error) {
	return &brainfuck{
		size:    32768,
		stack:   make([]int, 32768),
		text:    "",
		index:   0,
		pointer: 0,
		str:     "",
	}, nil
}

func (b *brainfuck) Run(content io.Reader) error {

	if err := b.Read(content); err != nil {
		return err
	}

	for b.index < len(b.text) {

		switch string(b.text[b.index]) {
		case ".":
			b.Print()
		case ",":
			b.Coma()
		case ">":
			b.MoreThan()
		case "<":
			b.LessThan()
		case "+":
			b.Plus()
		case "-":
			b.Minus()
		case "[":
			b.LoopStart()
		case "]":
			b.LoopEnd()
		}

		b.index++
	}

	fmt.Println("\nbefore print", b.stack)

	return nil
}

func (b *brainfuck) Read(content io.Reader) error {

	// generate new bytes buffer
	var buf bytes.Buffer

	// write content into bytes buffer
	if _, err := buf.ReadFrom(content); err != nil {
		return err
	}

	// assign buf.string to b.text
	b.text = buf.String()

	return nil
}

func (b *brainfuck) Print() {
	fmt.Print(string(b.stack[b.pointer]))
}

func (b *brainfuck) Plus() {
	b.stack[b.pointer] = (b.stack[b.pointer] + 255 + 1) % 255
}

func (b *brainfuck) Minus() {
	b.stack[b.pointer] = (b.stack[b.pointer] + 255 - 1) % 255
}

func (b *brainfuck) MoreThan() {
	b.pointer = (b.pointer + b.size + 1) % b.size
}

func (b *brainfuck) LessThan() {
	b.pointer = (b.pointer + b.size - 1) % b.size
}

// TODO: to be checked!!
func (b *brainfuck) Coma() {
	fmt.Scanln(&b.str)
	for _, c := range b.str {
		b.stack[b.index] = int(c)
	}
}

func (b *brainfuck) LoopStart() {
	if b.stack[b.pointer] == 0 {
		for string(b.text[b.index]) != "]" {
			b.index += 1
		}
	}
}

func (b *brainfuck) LoopEnd() {
	if b.stack[b.pointer] != 0 {
		for string(b.text[b.index]) != "[" {
			b.index -= 1
		}
	}
}
