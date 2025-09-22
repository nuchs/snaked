package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

func GridOrigin(s tcell.Screen, wantW, wantH int) (int, int) {
	w, h := s.Size()
	needW := wantW + 2
	needH := wantH + 2
	if needW > w {
		needW = w - 2
	}
	if needH > h {
		needH = h - 2
	}
	oX := (w - needW) / 2
	oY := (h - needH) / 2

	return oX + 1, oY + 1
}

func ToScreenCoords(oX, oY, gX, gY int) (int, int) {
	return oX + gX, oY + gY
}

func InBounds(w, h, gx, gy int) bool {
	return gx >= 0 && gx < w && gy >= 0 && gy < h
}

func update() {
	log.Printf("tick, direction = %s\n", direction)
}
