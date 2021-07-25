package main

import "fmt"

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
	position  Position
}

func NewSnake(d Direction, x int, y int) *Snake {
	return &Snake{
		direction: d,
		position: Position{
			x,
			y,
		},
	}
}

func (s *Snake) move() {
	switch s.direction {
	case Up:
		s.position.y--
	case Right:
		s.position.x++
	case Down:
		s.position.y++
	case Left:
		s.position.x--
	}
}

func (s *Snake) changeDirection(d Direction) {
	if s.direction == d || s.isOppositeDirection(s.direction, d) {
		return
	}

	fmt.Println(d) // debug
	s.direction = d
	fmt.Println(s.direction) // debug
}

func (s *Snake) isOppositeDirection(current Direction, target Direction) bool {
	switch current {
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
	return true
}
