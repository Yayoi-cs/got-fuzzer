package main

import (
	"fmt"
	"got-fuzzer/parser"
)

func banner() {
	fmt.Println("got-fuzzer")
	fmt.Println("           m    m\"\"")
	fmt.Println(" mmmm mmmmm#mmmm#mmm   mmmmmmmmmmm mmm m mm")
	fmt.Println("#\" \"##\" \"# #    #  #   #   m\"   m\"#\"  ##\"  \"")
	fmt.Println("#   ##   # #    #  #   # m\"   m\"  #\"\"\"\"#")
	fmt.Println("\"#m\"#\"#m#\" \"mm  #  \"mm\"##mmmm#mmmm\"#mm\"#")
	fmt.Println(" m  #")
	fmt.Println("  \"\"")
}

func main() {
	banner()
	filename := parser.CmdArgsParse()
	parser.HeaderParser(filename)
}
