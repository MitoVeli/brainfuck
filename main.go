package main

import (
	"fmt"
	"io"
	"os"

	bfLibrary "github.com/MitoVeli/brainfuck/internal/service"
)

func main() {

	// read file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var content io.Reader = file

	bf, err := bfLibrary.NewBrainfuckService()
	if err != nil {
		panic(err)
	}

	bf.AddCustomCommand("/", func() {
		fmt.Print(" & custom command divides the index by the pointer, the result: ", bf.GetIndex()/bf.GetPointer(), " & ")
	})

	// bf.RemoveCommand(">")

	if err = bf.Run(content); err != nil {
		fmt.Println("Error while running brainfuck service:", err)
	}
}
