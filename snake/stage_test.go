package snake

import "testing"

func TestStageWall(t *testing.T) {
	tests := []struct {
		name   string
		width  int
		height int
		target Position
		want   bool
	}{
		{
			"on floor",
			10,
			10,
			Position{1, 1},
			false,
		},
		{
			"on wall (top left corner)",
			10,
			10,
			Position{0, 0},
			true,
		},
		{
			"on wall (top right corner)",
			10,
			10,
			Position{9, 0},
			true,
		},
		{
			"on wall (bottom left corner)",
			10,
			10,
			Position{0, 9},
			true,
		},
		{
			"on wall (bottom right corner)",
			10,
			10,
			Position{9, 9},
			true,
		},
		{
			"on wall (top edge)",
			10,
			10,
			Position{3, 0},
			true,
		},
		{
			"on wall (left edge)",
			10,
			10,
			Position{0, 5},
			true,
		},
		{
			"on wall (bottom edge)",
			10,
			10,
			Position{2, 9},
			true,
		},
		{
			"on wall (right edge)",
			10,
			10,
			Position{9, 7},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			snk := NewSnake(Down, Position{2, 2})
			stg := NewStage(tt.width, tt.height, snk)
			result := stg.IsWall(tt.target)
			if result != tt.want {
				t.Errorf("invalid floor judgement! expected: %+v, actual: %+v\n", tt.want, result)
			}
		})
	}
}
