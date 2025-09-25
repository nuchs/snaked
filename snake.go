package main

type Snake struct {
	direction     Direction
	body          []Point
	PendingGrowth int
}

func NewSnake(start Point, length int, dir Direction) *Snake {
	if length < 1 {
		length = 1
	}

	body := make([]Point, length)
	for i := 1; i < length; i++ {
		switch dir {
		case Up:
			body[i] = Point{start.x, start.y + i}
		case Down:
			body[i] = Point{start.x, start.y - i}
		case Left:
			body[i] = Point{start.x + i, start.y}
		case Right:
			body[i] = Point{start.x - i, start.y}
		}
	}

	return &Snake{body: body, direction: dir}
}

func (s *Snake) SetDirection(d Direction) {
	if d == None {
		return
	}
	s.direction = d
}

func (s *Snake) Head() Point {
	return s.body[0]
}

func (s *Snake) Step(width, height int) Point {
	h := s.Head()
	switch s.direction {
	case Up:
		h.y++
	case Down:
		h.y--
	case Right:
		h.x++
	case Left:
		h.x--
	}
	return h
}

func (s *Snake) Advance(nextHead Point) {
	s.body = append(s.body, nextHead)
	if s.PendingGrowth > 0 {
		s.PendingGrowth--
		return
	}
	s.body = s.body[:len(s.body)-1]
}
