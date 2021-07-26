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
			snk := NewSnake(Down, []Position{{2, 2}})
			stg := NewStage(tt.width, tt.height, snk)
			result := stg.IsWall(tt.target)
			if result != tt.want {
				t.Errorf("invalid floor judgement! expected: %+v, actual: %+v\n", tt.want, result)
			}
		})
	}
}

func TestStagePlaceFood(t *testing.T) {
	snk := NewSnake(Down, []Position{{5, 5}})
	stg := NewStage(10, 10, snk)
	stg.PlaceFood()
	// food が1つだけ存在することを確認する
	cnt := 0
	for y := 0; y < stg.height; y++ {
		for x := 0; x < stg.height; x++ {
			if stg.IsFood(Position{X: x, Y: y}) {
				cnt++
			}
		}
	}
	if cnt != 1 {
		t.Errorf("invalid food number! expected: %+v, actual: %+v\n", 1, cnt)
	}
}

func TestStageNotPlaceFoodIfFoodExists(t *testing.T) {
	snk := NewSnake(Down, []Position{{5, 5}})
	stg := NewStage(10, 10, snk)
	stg.PlaceFood()
	stg.PlaceFood()
	cnt := 0
	for y := 0; y < stg.height; y++ {
		for x := 0; x < stg.height; x++ {
			if stg.IsFood(Position{X: x, Y: y}) {
				cnt++
			}
		}
	}
	if cnt != 1 {
		t.Errorf("invalid food number! expected: %+v, actual: %+v\n", 1, cnt)
	}
}

func TestStagePlaceFoodNotUpdateExistFood(t *testing.T) {
	snk := NewSnake(Down, []Position{{5, 5}})
	stg := NewStage(10, 10, snk)
	stg.PlaceFood()
	where := stg.food.Where()
	stg.PlaceFood()
	current := stg.food.Where()
	if where != current {
		t.Errorf("invalid food place! expected: %+v, actual: %+v\n", where, current)
	}
}

func TestStageIsSnakeBody(t *testing.T) {
	tests := []struct {
		name     string
		sections []Position
		target   Position
		want     bool
	}{
		{
			"1 section snake (not match)",
			[]Position{
				{1, 1},
			},
			Position{1, 2},
			false,
		},
		{
			"1 section snake (match head)",
			[]Position{
				{1, 1},
			},
			Position{1, 1},
			false,
		},
		{
			"2 section snake (not match)",
			[]Position{
				{10, 10},
				{10, 9},
			},
			Position{10, 11},
			false,
		},
		{
			"2 section snake (match head)",
			[]Position{
				{10, 10},
				{10, 9},
			},
			Position{10, 9},
			false,
		},
		{
			"2 section snake (match body)",
			[]Position{
				{10, 10},
				{10, 9},
			},
			Position{10, 10},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			snk := NewSnake(Up, tt.sections)
			stg := NewStage(100, 100, snk)
			result := stg.IsSnakeBody(tt.target)
			if result != tt.want {
				t.Errorf("invalid snake judgement! expected: %+v, actual: %+v\n", tt.want, result)
			}
		})
	}
}

func TestStageIsSnakeHead(t *testing.T) {
	tests := []struct {
		name     string
		sections []Position
		target   Position
		want     bool
	}{
		{
			"1 section snake (not match)",
			[]Position{
				{1, 1},
			},
			Position{1, 2},
			false,
		},
		{
			"1 section snake (match head)",
			[]Position{
				{1, 1},
			},
			Position{1, 1},
			true,
		},
		{
			"2 section snake (not match)",
			[]Position{
				{10, 10},
				{10, 9},
			},
			Position{10, 11},
			false,
		},
		{
			"2 section snake (match head)",
			[]Position{
				{10, 10},
				{10, 9},
			},
			Position{10, 9},
			true,
		},
		{
			"2 section snake (match body)",
			[]Position{
				{10, 10},
				{10, 9},
			},
			Position{10, 10},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			snk := NewSnake(Up, tt.sections)
			stg := NewStage(100, 100, snk)
			result := stg.IsSnakeHead(tt.target)
			if result != tt.want {
				t.Errorf("invalid snake judgement! expected: %+v, actual: %+v\n", tt.want, result)
			}
		})
	}
}
