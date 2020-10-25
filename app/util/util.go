package util

import (
	"fmt"
	"os"
)

// Er prints one or more message to stderr and exits the program
func Er(msg ...interface{}) {
	fmt.Fprintln(os.Stderr, msg...)
	os.Exit(1)
}
