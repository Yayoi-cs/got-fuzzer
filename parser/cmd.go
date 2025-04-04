package parser

import (
	"flag"
	"fmt"
	"os"
)

func CmdArgsParse() string {
	file := flag.String("f", "", "header file ex(:/usr/include/stdio.h")
	flag.Parse()
	if *file == "" {
		fmt.Println("Usage: <binary> -f <header file>")
		os.Exit(1)
	}
	_, err := os.Stat(*file)
	if err != nil {
		localError(err)
	}

	return *file
}
