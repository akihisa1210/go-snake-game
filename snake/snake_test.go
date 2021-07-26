package snake

import (
	"testing"
)

func TestSnakeMove(t *testing.T) {
	tests := []struct {
		name              string
		current_direction Direction
		current_position  []Position
		want              Position
	}{
		{
			"move to up",
			Up,
			[]Position{
				{10, 10},
			},
			Position{10, 9},
		},
		{
			"move to right",
			Right,
			[]Position{
				{10, 10},
			},
			Position{11, 10},
		},
		{
			"move to down",
			Down,
			[]Position{
				{10, 10},
			},
			Position{10, 11},
		},
		{
			"move to left",
			Left,
			[]Position{
				{10, 10},
			}, Position{9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSnake(tt.current_direction, tt.current_position)
			s.Move()
			if s.sections[0] != tt.want {
				t.Errorf("invalid movement! expected: %+v, actual: %+v\n", tt.want, s.sections)
			}
		})
	}
}

func TestSnakeMoveWithMultipleSections(t *testing.T) {
	tests := []struct {
		name              string
		current_direction Direction
		current_sections  []Position
		want              []Position
	}{
		{
			"2 sections and up",
			Up,
			[]Position{
				{10, 11},
				{10, 10},
			},
			[]Position{
				{10, 10},
				{10, 9},
			},
		},
		{
			"3 sections and right",
			Right,
			[]Position{
				{4, 6},
				{4, 7},
				{5, 7},
			},
			[]Position{
				{4, 7},
				{5, 7},
				{6, 7},
			},
		},
		{
			"4 sections and down",
			Down,
			[]Position{
				{22, 7},
				{21, 7},
				{21, 8},
				{21, 9},
			},
			[]Position{
				{21, 7},
				{21, 8},
				{21, 9},
				{21, 10},
			},
		},
		{
			"5 sections and left",
			Left,
			[]Position{
				{9, 7},
				{8, 7},
				{8, 8},
				{8, 7},
				{7, 7},
			},
			[]Position{
				{8, 7},
				{8, 8},
				{8, 7},
				{7, 7},
				{6, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSnake(tt.current_direction, tt.current_sections)
			s.Move()
			for i, w := range tt.want {
				if s.sections[i] != w {
					t.Errorf("invalid position! index: %+v, expected: %+v, actual: %+v\n", i, w, s.sections[i])
				}
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
			s := NewSnake(tt.current_direction, []Position{{0, 0}})
			s.ChangeDirection(tt.target_direction)
			if s.direction != tt.want {
				t.Errorf("invalid direction! expected: %+v, actual: %+v\n", tt.want, s.direction)
			}
		})
	}
}

func TestSnakeGetPosition(t *testing.T) {
	tests := []struct {
		name             string
		current_position []Position
		want             Position
	}{
		{
			"(1, 1)",
			[]Position{
				{1, 1},
			},
			Position{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSnake(Down, tt.current_position)
			cp := s.GetCurrentHeadPosition()
			if cp != tt.want {
				t.Errorf("invalid position! expected: %+v, actual: %+v\n", tt.want, cp)
			}
		})
	}
}

func TestSnakeEatFood(t *testing.T) {
	snk := NewSnake(Down, []Position{{5, 5}})
	stg := NewStage(10, 10, snk)
	stg.food = &Food{position: Position{5, 5}}
	stg.SnakeEatsFood()

	result := stg.food
	if result != nil {
		t.Errorf("food is not eaten! position: %+v\n", result)
	}

	l := len(stg.snake.sections)
	if l != 2 {
		t.Errorf("invalid length! expected: %+v, actual: %+v\n", 2, l)
	}
}
