package cmd

import (
	"fmt"
	"golang.org/x/term"
)

func GetTerminalSize() (int, int) {
	width, height, err := term.GetSize(0)
	if err != nil {
		fmt.Printf("error getting terminal size: %v", err)
	}
	return width, height
}

func ClearTerminal() {
	for {
		fmt.Println("\x1b[?25l")
		fmt.Println("\x1b[2J")
	}
}
