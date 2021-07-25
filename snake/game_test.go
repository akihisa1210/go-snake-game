package snake

import (
	"fmt"
	"testing"
)

func TestGameNotOnWall(t *testing.T) {
	snk := NewSnake(Down, Position{1, 1})
	stg := NewStage(3, 3, snk)
	game := NewGame(stg)
	fmt.Printf("%+v\n", snk)
	want := false
	result := game.IsOver()
	if result != want {
		t.Errorf("invalid game over! expected: %+v, actual: %+v\n", want, result)
	}
}

func TestGameOverByWall(t *testing.T) {
	snk := NewSnake(Down, Position{1, 1})
	stg := NewStage(3, 3, snk)
	game := NewGame(stg)
	snk.Move()
	fmt.Printf("%+v\n", snk)
	want := true
	result := game.IsOver()
	if result != want {
		t.Errorf("invalid game over! expected: %+v, actual: %+v\n", want, result)
	}
}
