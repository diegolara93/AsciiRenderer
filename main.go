package main

import (
	"AsciiRenderer/cmd"
	"github.com/gdamore/tcell/v2"
	"log"
)

func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorRed)
	//boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	width, height := s.Size()

	cmd.DrawLine(s, 0, 0, width, height, defStyle)
	cmd.DrawLine(s, width, 0, 0, height, defStyle)
	quit := func() {
		// panics have to be caught or else program will crash without a trace
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	// Event loop
	for {
		// Update screen
		s.Show()

		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
			s.Clear()
			width, height := s.Size()

			cmd.DrawLine(s, 0, 0, width, height, defStyle)
			cmd.DrawLine(s, width, 0, 0, height, defStyle)
			cmd.DrawSquare(s, width/2-(width/20), height/3, 10, 10)
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				s.Clear()
			}
		case *tcell.EventMouse:

		}
	}
}
