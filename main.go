package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

var direction = "up"

func main() {
	// cfg, err := FromFlags()
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "Invalid config:", err)
	// 	os.Exit(2)
	// }

	screen := createScreen()
	defer safeFini(screen)

	base := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	screen.SetStyle(base)
	screen.Clear()
	screen.Show()

loop:
	for {
		ev := screen.PollEvent()
		switch e := ev.(type) {
		case *tcell.EventKey:
			if handlekey(e) {
				break loop
			}
			redraw(screen, base)
		case *tcell.EventResize:
			redraw(screen, base)
		case *tcell.EventInterrupt:
			break loop
		}
	}
}

func createScreen() tcell.Screen {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintln(os.Stderr, "creating screen: ", err)
		os.Exit(1)
	}
	if err = screen.Init(); err != nil {
		fmt.Fprintln(os.Stderr, "initialising screen: ", err)
		os.Exit(1)
	}

	return screen
}

func handlekey(key *tcell.EventKey) bool {
	switch key.Key() {
	case tcell.KeyEscape, tcell.KeyCtrlC:
		return true
	}

	switch key.Rune() {
	case 'a', 'A':
		direction = "left"
	case 'w', 'W':
		direction = "up"
	case 'd', 'D':
		direction = "right"
	case 's', 'S':
		direction = "down"
	}

	return false
}

func redraw(s tcell.Screen, st tcell.Style) {
	w, h := s.Size()
	status := fmt.Sprintf("Direction: %s (Esc to quit)", direction)

	for x := range w {
		DrawRune(s, x, h-1, ' ', st)
	}

	DrawString(s, 0, h-1, status, st)
	s.Show()
}

func safeFini(s tcell.Screen) {
	if s == nil {
		return
	}

	defer s.Fini()
	if r := recover(); r != nil {
		fmt.Fprintln(os.Stderr, "panic:", r)
	}
}
