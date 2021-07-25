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
			"on wall (corner)",
			10,
			10,
			Position{0, 0},
			true,
		},
		{
			"on wall (edge)",
			10,
			10,
			Position{0, 5},
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
