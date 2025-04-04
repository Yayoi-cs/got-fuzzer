package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	C_CHAR int = iota
	C_SIGNED_CHAR
	C_UNSIGNED_CHAR
	C_SHORT
	C_UNSIGNED_SHORT
	C_INT
	C_UNSIGNED_INT
	C_LONG
	C_UNSIGNED_LONG
	C_LONG_LONG
	C_UNSIGNED_LONG_LONG
	C_FLOAT
	C_DOUBLE
	C_LONG_DOUBLE
	C_BOOL
	C_VOID
	C_SIZE_T
	C_PTRDIFF_T
	C_WCHAR_T
	C_WINT_T
	C_SIG_ATOMIC_T
	C_TIME_T
	C_CLOCK_T
	C_VA_LIST
	C_JMPBUF
	C_FILE
	C_FPOS_T
	C_INTMAX_T
	C_UINTMAX_T
	C_INTPTR_T
	C_UINTPTR_T
	C_INT8_T
	C_UINT8_T
	C_INT16_T
	C_UINT16_T
	C_INT32_T
	C_UINT32_T
	C_INT64_T
	C_UINT64_T
)

var CTypes = []cType{
	{name: "char", ident: C_CHAR},
	{name: "signed char", ident: C_SIGNED_CHAR},
	{name: "unsigned char", ident: C_UNSIGNED_CHAR},
	{name: "short", ident: C_SHORT},
	{name: "unsigned short", ident: C_UNSIGNED_SHORT},
	{name: "int", ident: C_INT},
	{name: "unsigned int", ident: C_UNSIGNED_INT},
	{name: "long", ident: C_LONG},
	{name: "unsigned long", ident: C_UNSIGNED_LONG},
	{name: "long long", ident: C_LONG_LONG},
	{name: "unsigned long long", ident: C_UNSIGNED_LONG_LONG},
	{name: "float", ident: C_FLOAT},
	{name: "double", ident: C_DOUBLE},
	{name: "long double", ident: C_LONG_DOUBLE},
	{name: "_Bool", ident: C_BOOL},
	{name: "void", ident: C_VOID},
	{name: "size_t", ident: C_SIZE_T},
	{name: "ptrdiff_t", ident: C_PTRDIFF_T},
	{name: "wchar_t", ident: C_WCHAR_T},
	{name: "wint_t", ident: C_WINT_T},
	{name: "sig_atomic_t", ident: C_SIG_ATOMIC_T},
	{name: "time_t", ident: C_TIME_T},
	{name: "clock_t", ident: C_CLOCK_T},
	{name: "va_list", ident: C_VA_LIST},
	{name: "jmp_buf", ident: C_JMPBUF},
	{name: "FILE", ident: C_FILE},
	{name: "fpos_t", ident: C_FPOS_T},
	{name: "intmax_t", ident: C_INTMAX_T},
	{name: "uintmax_t", ident: C_UINTMAX_T},
	{name: "intptr_t", ident: C_INTPTR_T},
	{name: "uintptr_t", ident: C_UINTPTR_T},
	{name: "int8_t", ident: C_INT8_T},
	{name: "uint8_t", ident: C_UINT8_T},
	{name: "int16_t", ident: C_INT16_T},
	{name: "uint16_t", ident: C_UINT16_T},
	{name: "int32_t", ident: C_INT32_T},
	{name: "uint32_t", ident: C_UINT32_T},
	{name: "int64_t", ident: C_INT64_T},
	{name: "uint64_t", ident: C_UINT64_T},
}

type cType struct {
	name  string
	ident int
}

type cFuncParam struct {
	paramType    string
	paramName    string
	isPointer    bool
	pointerDepth int
}

type cFunction struct {
	name         string
	returnType   string
	isRetPointer bool
	params       []cFuncParam
	isExtern     bool
}

func exportParser(filename string) []string {

	return nil
}

func includeParser(filename string) []string {
	fd, err := os.Open(filename)
	if err != nil {
		localError(err)
	}
	defer fd.Close()

	var incList []string

	sc := bufio.NewScanner(fd)
	for sc.Scan() {
		line := sc.Text()
		if strings.HasPrefix(line, "#include") {
			logger(line)
			incList = append(incList, line)
		}
	}
	return incList
}

func logger(message string) {
	fmt.Println("[+]", message)
}

func localError(e error) {
	fmt.Println("Error:", e)
	os.Exit(1)
}

func HeaderParser(filename string) {
	res := includeParser(filename)
	res = exportParser(filename)

	fmt.Println(res)
}
