package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

var direction = "up"

func Run(cfg Config) error {
	screen, err := createScreen()
	if err != nil {
		return err
	}
	defer safeFini(screen)
	screen.Clear()
	screen.Show()

	ch := startEventHandler(screen)
	ticker := startTicker(cfg)

loop:
	for {
		select {
		case ev := <-ch:
			if handleEvent(screen, cfg, ev) {
				break loop
			}
		case <-ticker.C:
			update()
			render(screen, cfg, direction)
		}
	}

	return nil
}

func createScreen() (tcell.Screen, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, fmt.Errorf("Failed to create scree: %w", err)
	}
	if err = screen.Init(); err != nil {
		return nil, fmt.Errorf("Failed to create scree: %w", err)
	}

	return screen, nil
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

func startEventHandler(s tcell.Screen) <-chan tcell.Event {
	eventCh := make(chan tcell.Event, 10)
	go func() {
		for {
			ev := s.PollEvent()
			eventCh <- ev
		}
	}()

	return eventCh
}

func startTicker(cfg Config) *time.Ticker {
	ticker := time.NewTicker(time.Second / time.Duration(cfg.Speed))

	return ticker
}

func handleEvent(s tcell.Screen, cfg Config, ev tcell.Event) bool {
	switch e := ev.(type) {
	case *tcell.EventKey:
		switch e.Rune() {
		case 'w', 'W':
			direction = "up"
		case 's', 'S':
			direction = "down"
		case 'a', 'A':
			direction = "left"
		case 'd', 'D':
			direction = "right"
		}
		if e.Key() == tcell.KeyEscape || e.Key() == tcell.KeyCtrlC {
			return true
		}

	case *tcell.EventResize:
		s.Sync()
		render(s, cfg, direction)
	}

	return false
}
