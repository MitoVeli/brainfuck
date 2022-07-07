# Brainfuck Library
## Import
Library can be imported by the following command
-   `go get github.com/MitoVeli/brainfuck/internal/service`

## Usage
No CLI executable is available. Library itsef can be tried with the example "main.go" file in the directory

Input.txt file, which is located in the root folder, is used as source by main.go. For trying the library with different inputs/commands, input.txt needs to be edited.

If library gets any undefined commands in the input.txt, it returns error as in the below example;
-   `Error while running brainfuck service: '*' not defined within brainfuck service commands`

In main.go, custom commands can be added as below
-   `bf.AddCustomCommand("/", func() {fmt.Print("custom command ")})`
-    i.e (returns Hello custom command World!)

Similary, any command can be removed as in the below example;
-   `bf.RemoveCommand(">")`
