package main

import "testing"

func TestSnakeMove(t *testing.T) {
	tests := []struct {
		name              string
		current_direction Direction
		current_position  Position
		want              Position
	}{
		{
			"move to up",
			Up,
			Position{10, 10},
			Position{10, 9},
		},
		{
			"move to right",
			Right,
			Position{10, 10},
			Position{11, 10},
		},
		{
			"move to down",
			Down,
			Position{10, 10},
			Position{10, 11},
		},
		{
			"move to left",
			Left,
			Position{10, 10},
			Position{9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSnake(tt.current_direction, tt.current_position)
			s.move()
			if s.section != tt.want {
				t.Errorf("invalid movement! expected: %+v, actual: %+v\n", tt.want, s.section)
			}
		})
	}
}

func TestSnakeChangeDirection(t *testing.T) {
	tests := []struct {
		name              string
		current_direction Direction
		target_direction  Direction
		want              Direction
	}{
		{
			"down to right",
			Down,
			Right,
			Right,
		},
		{
			"same direction (down to down)",
			Down,
			Down,
			Down,
		},
		{
			"opposite direction (right to left)",
			Right,
			Left,
			Right,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSnake(tt.current_direction, Position{0, 0})
			s.changeDirection(tt.target_direction)
			if s.direction != tt.want {
				t.Errorf("invalid direction! expected: %+v, actual: %+v\n", tt.want, s.direction)
			}
		})
	}
}
