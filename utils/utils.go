package utils

import (
	"fmt"
	"os"
)

func PrintfSTDOUT(format string, a ...any) {
	fmt.Fprintf(os.Stdout, format+"\n", a...)
}

func PrintfSTDERR(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format, a...)
}
