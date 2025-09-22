package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func DrawRune(s tcell.Screen, x, y int, r rune, st tcell.Style) {
	s.SetContent(x, y, r, nil, st)
}

func DrawString(s tcell.Screen, x, y int, str string, st tcell.Style) {
	col := x
	for _, r := range str {
		s.SetContent(col, y, r, nil, st)
		col++
	}
}

func DrawBox(s tcell.Screen, x1, y1, x2, y2 int, st tcell.Style) {
	for col := x1; col < x2; col++ {
		DrawRune(s, col, y1, '─', st)
		DrawRune(s, col, y2, '─', st)
	}

	for row := y1; row < y2; row++ {
		DrawRune(s, x1, row, '│', st)
		DrawRune(s, x2, row, '│', st)
	}

	DrawRune(s, x1, y1, '┌', st)
	DrawRune(s, x1, y2, '└', st)
	DrawRune(s, x2, y1, '┐', st)
	DrawRune(s, x2, y2, '┘', st)
}

func DrawPlayfield(s tcell.Screen, W, H int, st tcell.Style) (ox, oy int) {
	ox, oy = GridOrigin(s, W, H)
	// Border coordinates surround the grid by 1 cell.
	x1, y1 := ox-1, oy-1
	x2, y2 := ox+W, oy+H
	DrawBox(s, x1, y1, x2, y2, st)
	return ox, oy
}

func render(s tcell.Screen, cfg Config, direction string) {
	s.Clear()

	playfieldStyle := tcell.StyleDefault.Foreground(tcell.ColorGreen)
	textStyle := tcell.StyleDefault

	ox, oy := DrawPlayfield(s, cfg.Width, cfg.Height, playfieldStyle)

	// Example: draw a test “head” at the center of the grid.
	hx, hy := cfg.Width/2, cfg.Height/2
	if InBounds(cfg.Width, cfg.Height, hx, hy) {
		sx, sy := ToScreenCoords(ox, oy, hx, hy)
		DrawRune(s, sx, sy, '●', tcell.StyleDefault.Foreground(tcell.ColorYellow))
	}

	// Status line (outside border if space allows)
	tw, th := s.Size()
	status := fmt.Sprintf("W=%d H=%d Dir=%s  (ESC to quit, WASD to change)", cfg.Width, cfg.Height, direction)
	for x := range tw {
		s.SetContent(x, th-1, ' ', nil, textStyle)
	}
	DrawString(s, 0, th-1, status, textStyle.Dim(true))

	s.Show()
}
