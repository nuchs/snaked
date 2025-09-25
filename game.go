package main

import (
	"github.com/gdamore/tcell/v2"
)

type Direction int

const (
	None Direction = iota
	Up
	Down
	Left
	Right
)

type Point struct {
	x int
	y int
}

func gridLayout(s tcell.Screen, reqW, reqH int) (ox, oy, W, H int, border bool) {
	tw, th := s.Size()

	// Try to keep a 1-cell border if we have at least 3x3 space.
	if tw >= 3 && th >= 3 && reqW+2 <= tw && reqH+2 <= th {
		border = true
		W = reqW
		H = reqH
		needW, needH := W+2, H+2
		boxX := (tw - needW) / 2
		boxY := (th - needH) / 2
		ox, oy = boxX+1, boxY+1 // inside the border
		return
	}

	// Otherwise, drop the border and clamp grid to terminal size (or 0 if tiny).
	border = false
	if tw < 1 || th < 1 {
		return 0, 0, 0, 0, false
	}
	W = min(reqW, tw)
	H = min(reqH, th)
	boxX := (tw - W) / 2
	boxY := (th - H) / 2
	ox, oy = boxX, boxY // no border, so origin is the top-left of the grid
	return
}

func ToScreenCoords(oX, oY, gX, gY int) (int, int) {
	return oX + gX, oY + gY
}

func InBounds(w, h, gx, gy int) bool {
	return gx >= 0 && gx < w && gy >= 0 && gy < h
}

func update(screen tcell.Screen, cfg Config, snake *Snake) {
	_, _, w, h, _ := gridLayout(screen, cfg.Width, cfg.Height)
	next := snake.Step(w, h)
	snake.Advance(next)
	render(screen, cfg, snake)
}
