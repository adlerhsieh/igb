package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

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
