package repl

import (
	"Interpreter/evaluator"
	"Interpreter/lexer"
	"Interpreter/object"
	"Interpreter/parser"
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		if filepath.Ext(line) == ".Steins" {
			executeFile(line, out, env)
			continue
		}
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}

}
func executeFile(filename string, out io.Writer, env *object.Environment) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(out, "Failed to read file %s: %v\n", filename, err)
		return
	}

	sourceCode := string(content)
	fmt.Fprintf(out, "Processing from file %s:\n", filename)
	fmt.Println(sourceCode)
	fmt.Fprintln(out, "Output:")

	l := lexer.New(sourceCode)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		fmt.Fprintln(out, evaluated.Inspect())
	}
}

const MONKEY = `           __,__ 
  .--.  .-"     "-.  .--.
 / .. \/  .-. .-.  \/ .. \
| |  '|  /   Y   \  |'  | |
| \   \ \ 0 | 0 / /   / | |
 \ '- ,\.-"""""""-./, -' /
  ''-' /_   ^ ^   _\ '-''
      |  \._   _./  |
      \   \ '~' /   /
       '._ '-=-' _.'
          '-----'       
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, "parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
