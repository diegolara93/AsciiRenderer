package main

import (
	"AsciiRenderer/cmd"
	"fmt"
)

func main() {
	cmd.TestVecs()
	fmt.Println()
	termWidth, termHeight := cmd.GetTerminalSize()
	s := cmd.NewScreen(termWidth, termHeight)

}
