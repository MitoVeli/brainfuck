package service

import (
	"bytes"
	"fmt"
	"io"
)

type BrainfuckService interface {
	Run(content io.Reader) error
	AddCustomCommand(symbol string, function func())
	RemoveCommand(symbol string)
	GetStack() []int
	GetSize() int
	GetText() string
	GetIndex() int
	GetPointer() int
	GetStr() string
}

type Brainfuck struct {
	stack    []int
	size     int
	text     string
	index    int
	pointer  int
	str      string
	commands map[string]func()
}

func NewBrainfuckService() (BrainfuckService, error) {
	b := Brainfuck{
		stack:    make([]int, 32768),
		size:     32768,
		text:     "",
		index:    0,
		pointer:  0,
		str:      "",
		commands: map[string]func(){},
	}

	b.commands["."] = b.Print
	b.commands[","] = b.Coma
	b.commands[">"] = b.MoreThan
	b.commands["<"] = b.LessThan
	b.commands["+"] = b.Plus
	b.commands["-"] = b.Minus
	b.commands["["] = b.LoopStart
	b.commands["]"] = b.LoopEnd

	return &b, nil
}

func (b *Brainfuck) AddCustomCommand(symbol string, function func()) {
	b.commands[symbol] = function
}

func (b *Brainfuck) RemoveCommand(symbol string) {
	delete(b.commands, symbol)
}

func (b *Brainfuck) Run(content io.Reader) error {

	// read content and store it in b.text
	if err := b.Read(content); err != nil {
		return err
	}

	// check if there is any undefined command received in b.text
	var idx int = 0
	for idx < len(b.text) {
		if _, ok := b.commands[string(b.text[idx])]; !ok {
			return fmt.Errorf("'%s' not defined within brainfuck service commands", string(b.text[idx]))
		}
		idx++
	}

	// run commands accordingly
	for b.index < len(b.text) {

		symbol := string(b.text[b.index])
		b.commands[symbol]()

		b.index++
	}

	return nil
}

func (b *Brainfuck) Read(content io.Reader) error {

	// generate new bytes buffer
	var buf bytes.Buffer

	// write content into bytes buffer
	if _, err := buf.ReadFrom(content); err != nil {
		return err
	}

	// stringify bytes buffer and store it in b.text
	b.text = buf.String()

	return nil
}

func (b *Brainfuck) Print() {
	fmt.Print(string(b.stack[b.pointer]))
}

func (b *Brainfuck) Plus() {
	b.stack[b.pointer] += 1
}

func (b *Brainfuck) Minus() {
	b.stack[b.pointer] -= 1
}

func (b *Brainfuck) MoreThan() {
	b.pointer += 1
}

func (b *Brainfuck) LessThan() {
	b.pointer -= 1
}

func (b *Brainfuck) Coma() {
	fmt.Scanln(&b.str)
	for _, c := range b.str {
		b.stack[b.index] = int(c)
	}
}

func (b *Brainfuck) LoopStart() {
	if b.stack[b.pointer] == 0 {
		for string(b.text[b.index]) != "]" {
			b.index += 1
		}
	}
}

func (b *Brainfuck) LoopEnd() {
	if b.stack[b.pointer] != 0 {
		for string(b.text[b.index]) != "[" {
			b.index -= 1
		}
	}
}

// Getters
func (b *Brainfuck) GetStack() []int {
	return b.stack
}

func (b *Brainfuck) GetSize() int {
	return b.size
}

func (b *Brainfuck) GetText() string {
	return b.text
}

func (b *Brainfuck) GetIndex() int {
	return b.index
}

func (b *Brainfuck) GetPointer() int {
	return b.pointer
}

func (b *Brainfuck) GetStr() string {
	return b.str
}
