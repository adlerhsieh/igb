package main

import (
	"bufio"
	"fmt"
	"github.com/goby-lang/goby/bytecode"
	"github.com/goby-lang/goby/parser"
	"github.com/goby-lang/goby/vm"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var LineNumber int

func PrintPrompt() {
	LineNumber = LineNumber + 1
	l := strconv.Itoa(LineNumber)
	if LineNumber < 10 {
		l = "00" + l
	} else if LineNumber < 100 {
		l = "0" + l
	}
	fmt.Print("igb(", l, ")>")
}

func main() {
	LoopInput()
	os.Remove(TmpFile())
	os.Exit(0)
}

func LoopInput() {
	for {
		codeInput := UserInput()
		if codeInput == "exit" {
			break
		}
		if codeInput == "" {
			continue
		}
		responseObject := Execute(codeInput)
		fmt.Println("# =>", responseObject.Inspect())
	}
}

func UserInput() string {
	reader := bufio.NewReader(os.Stdin)
	PrintPrompt()
	codeInput, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(codeInput)
}

func Dir() string {
	directory, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return directory
}

func TmpFile() string {
	return Dir() + "/tmp"
}

func WriteToTmpFile(codeInput string) {
	err := ioutil.WriteFile(TmpFile(), []byte(codeInput), 0644)
	if err != nil {
		panic(err)
	}
}

func Execute(codeInput string) vm.Object {
	WriteToTmpFile(codeInput)

	codeBytes := []byte(codeInput)
	program := parser.BuildAST(codeBytes)
	g := bytecode.NewGenerator(program)
	bytecodes := g.GenerateByteCode(program)

	v := vm.New(Dir(), []string{})
	v.ExecBytecodes(bytecodes, TmpFile())
	return v.GetExecResult()
}
