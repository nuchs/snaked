package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

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
	snake := NewSnake(Point{cfg.Width / 2, cfg.Height / 2}, 3, Right)

loop:
	for {
		select {
		case ev := <-ch:
			if handleEvent(screen, cfg, ev, snake) {
				break loop
			}
		case <-ticker.C:
			update(screen, cfg, snake)
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

func key2Dir(e *tcell.EventKey) (Direction, bool) {
	switch e.Rune() {
	case 'w', 'W':
		return Up, true
	case 's', 'S':
		return Down, true
	case 'a', 'A':
		return Left, true
	case 'd', 'D':
		return Right, true
	}
	return None, false
}

func handleEvent(s tcell.Screen, cfg Config, ev tcell.Event, snake *Snake) bool {
	switch e := ev.(type) {
	case *tcell.EventKey:
		if dir, ok := key2Dir(e); ok {
			snake.SetDirection(dir)
		}
		if e.Rune() == 'g' || e.Rune() == 'G' {
			snake.PendingGrowth += 3
		}
		if e.Key() == tcell.KeyEscape || e.Key() == tcell.KeyCtrlC {
			return true
		}

	case *tcell.EventResize:
		s.Sync()
		render(s, cfg, snake)
	}

	return false
}
