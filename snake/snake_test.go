package snake

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
			s.Move()
			if s.section[0] != tt.want {
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
		current_position Position
		want             Position
	}{
		{
			"(1, 1)",
			Position{1, 1},
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
	snk := NewSnake(Down, Position{5, 5})
	stg := NewStage(10, 10, snk)
	stg.food = &Food{position: Position{5, 5}}
	stg.SnakeEatsFood()

	result := stg.food
	if result != nil {
		t.Errorf("food is not eaten! position: %+v\n", result)
	}

	l := len(stg.snake.section)
	if l != 2 {
		t.Errorf("invalid length! expected: %+v, actual: %+v\n", 2, l)
	}
}
