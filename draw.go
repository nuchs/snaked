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

func render(s tcell.Screen, cfg Config, snake *Snake) {
	s.Clear()

	// Border if present
	borderStyle := tcell.StyleDefault.Foreground(tcell.ColorGreen)
	ox, oy, W, H, _ := gridLayout(s, cfg.Width, cfg.Height)
	DrawBox(s, ox-1, oy-1, ox+W, oy+H, borderStyle)

	// Draw snake: head bright, body dimmer
	headStyle := tcell.StyleDefault.Foreground(tcell.ColorYellow)
	bodyStyle := tcell.StyleDefault.Foreground(tcell.ColorLightYellow).Dim(true)

	for i, p := range snake.body {
		sx, sy := ToScreenCoords(ox, oy, p.x, p.y)
		ch := '●'
		if i == 0 {
			s.SetContent(sx, sy, ch, nil, headStyle)
		} else {
			s.SetContent(sx, sy, ch, nil, bodyStyle)
		}
	}

	// Status line
	tw, th := s.Size()
	status := fmt.Sprintf("Len=%d Dir=%v  (WASD/Arrows, ESC to quit)", len(snake.body), snake.direction)
	for x := 0; x < tw; x++ {
		s.SetContent(x, th-1, ' ', nil, tcell.StyleDefault)
	}
	DrawString(s, 0, th-1, status, tcell.StyleDefault.Dim(true))

	s.Show()
}
