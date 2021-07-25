package main

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Position struct {
	x int
	y int
}

type Snake struct {
	direction Direction
	section   Position
}

func NewSnake(d Direction, p Position) *Snake {
	return &Snake{
		direction: d,
		section:   p,
	}
}

func (s *Snake) move() {
	switch s.direction {
	case Up:
		s.section.y--
	case Right:
		s.section.x++
	case Down:
		s.section.y++
	case Left:
		s.section.x--
	}
}

func (s *Snake) changeDirection(d Direction) {
	if s.direction == d || s.isOppositeDirection(d) {
		return
	}
	s.direction = d
}

func (s *Snake) isOppositeDirection(target Direction) bool {
	switch s.direction {
	case Up:
		if target == Down {
			return true
		}
	case Right:
		if target == Left {
			return true
		}
	case Down:
		if target == Up {
			return true
		}
	case Left:
		if target == Right {
			return true
		}
	}
	return false
}
