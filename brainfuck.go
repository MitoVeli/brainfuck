package brainfuck

import (
	"bytes"
	"fmt"
	"io"
)

// func main() {

// 	content := strings.NewReader("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.o--------.>+.>.")

// 	bf, err := NewBrainfuckService()
// 	if err != nil {
// 		fmt.Println("Error while creating brainfuck service")
// 	}

// 	bf.AddCustomCommand("o", func() {
// 		fmt.Print("***", bf.GetIndex()/bf.GetPointer())
// 	})

// 	if err = bf.Run(content); err != nil {
// 		fmt.Println("Error while running brainfuck service", err)
// 	}
// }

type brainfuckService interface {
	Run(content io.Reader) error
	AddCustomCommand(symbol string, function func())
	GetStack() []int
	GetSize() int
	GetText() string
	GetIndex() int
	GetPointer() int
	GetStr() string
}

type brainfuck struct {
	stack    []int
	size     int
	text     string
	index    int
	pointer  int
	str      string
	commands map[string]func()
}

func NewBrainfuckService() (brainfuckService, error) {
	b := brainfuck{
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

func (b *brainfuck) AddCustomCommand(symbol string, function func()) {
	b.commands[symbol] = function
}

func (b *brainfuck) Run(content io.Reader) error {

	if err := b.Read(content); err != nil {
		return err
	}

	for b.index < len(b.text) {

		symbol := string(b.text[b.index])
		b.commands[symbol]()

		b.index++
	}

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

// TODO: to be checked!!
func (b *brainfuck) Print() {
	fmt.Print(string(b.stack[b.pointer]))
}

func (b *brainfuck) Plus() {
	b.stack[b.pointer] += 1
}

func (b *brainfuck) Minus() {
	b.stack[b.pointer] -= 1
}

func (b *brainfuck) MoreThan() {
	b.pointer += 1
}

func (b *brainfuck) LessThan() {
	b.pointer -= 1
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

// Getters
func (b *brainfuck) GetStack() []int {
	return b.stack
}

func (b *brainfuck) GetSize() int {
	return b.size
}

func (b *brainfuck) GetText() string {
	return b.text
}

func (b *brainfuck) GetIndex() int {
	return b.index
}

func (b *brainfuck) GetPointer() int {
	return b.pointer
}

func (b *brainfuck) GetStr() string {
	return b.str
}
