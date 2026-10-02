package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var stdinReader = bufio.NewReader(os.Stdin)

func readLine() string {
	line, _ := stdinReader.ReadString('\n')
	return strings.TrimSpace(line)
}

func readChoice() int {
	choice, err := strconv.Atoi(readLine())
	if err != nil {
		return 0
	}
	return choice
}
