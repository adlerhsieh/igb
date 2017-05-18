package main

import (
	"github.com/goby-lang/goby/bytecode"
	"github.com/goby-lang/goby/parser"
	"github.com/goby-lang/goby/vm"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	code := "puts(123)"

	Execute(code)
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

func Execute(codeInput string) {
	WriteToTmpFile(codeInput)

	codeBytes := []byte(codeInput)
	program := parser.BuildAST(codeBytes)
	g := bytecode.NewGenerator(program)
	bytecodes := g.GenerateByteCode(program)

	v := vm.New(Dir(), []string{})
	v.ExecBytecodes(bytecodes, TmpFile())
}
