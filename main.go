package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
)

func main() {
	cfg, err := FromFlags()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid config:", err)
		os.Exit(2)
	}

	log.Printf("Snake config -> %s", cfg)

	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintln(os.Stderr, "creating screen: ", err)
		os.Exit(1)
	}
	if err = screen.Init(); err != nil {
		fmt.Fprintln(os.Stderr, "initialising screen: ", err)
		os.Exit(1)
	}
	defer safeFini(screen)

	base := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	screen.SetStyle(base)
	screen.Clear()

	msg := "Hello! tcell - press ESC or Ctrl-C"
	drawCentered(screen, msg, base)
	screen.Show()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		screen.PostEvent((tcell.NewEventInterrupt(nil)))
	}()

	for {
		ev := screen.PollEvent()
		switch e := ev.(type) {
		case *tcell.EventKey:
			switch e.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return
			default:
				w, h := screen.Size()
				status := fmt.Sprintf("Last key: %v (rune=%q) - Esc/Ctrl+C to quit", e.Key(), e.Rune())
				for x := range w {
					screen.SetContent(x, h-1, ' ', nil, base)
				}
				drawAt := func(col, row int, s string, st tcell.Style) {
					for i, r := range s {
						if col+i >= w {
							break
						}
						screen.SetContent(col+i, row, r, nil, st)
					}
				}
				drawAt(0, h-1, status, base.Dim(true))
				screen.Show()
			}
		case *tcell.EventResize:
			screen.Clear()
			drawCentered(screen, msg, base)
			screen.Sync()
			screen.Show()

		case *tcell.EventInterrupt:
			return
		}
	}
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

func drawCentered(s tcell.Screen, msg string, st tcell.Style) {
	w, h := s.Size()
	if w <= 0 || h <= 0 {
		return
	}
	DrawBox(s, 0, 0, w-1, h-1, st.Bold(true))

	row := h / 2
	runeCount := utf8.RuneCountInString(msg)
	startCol := max((w-runeCount)/2, 0)

	DrawString(s, startCol, row, msg, st)
}
