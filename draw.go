package main

import "github.com/gdamore/tcell/v2"

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
