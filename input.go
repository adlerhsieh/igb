package main

import (
	"bufio"
	"fmt"
	"os"
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
