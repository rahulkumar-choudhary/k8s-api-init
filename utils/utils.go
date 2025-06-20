package utils

import (
	"fmt"
	"os"
	"strings"
)

const (
	patternLength = 66
)

func UserHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("unable to determine user home directory")
	}
	return homeDir
}

func ExitWithError(msg string) {
	os.Stderr.WriteString(msg + "\n")
	os.Exit(1)
}

func repeatPattern(pattern string, n int) {
	fmt.Println(strings.Repeat(pattern, n))
}

func logEmptyMessage(n int) {
	if n > 0 {
		return
	}
	fmt.Println(strings.Repeat(" ", patternLength/2) + "N/A")
}
