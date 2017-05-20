package main

import (
	"github.com/goby-lang/goby/bytecode"
	"github.com/goby-lang/goby/parser"
	"github.com/goby-lang/goby/vm"
	"os"
)

func main() {
	LoopInput()
	os.Remove(TmpFile())
	os.Exit(0)
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
