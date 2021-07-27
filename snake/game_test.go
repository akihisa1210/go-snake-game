package snake

import (
	"fmt"
	"testing"
)

func TestGameNotOnWall(t *testing.T) {
	snk := NewSnake(Down, []Position{{1, 1}})
	stg := NewStage(3, 3, snk)
	game := NewGame(stg)

	want := false
	result := game.IsOver()
	if result != want {
		t.Errorf("invalid game over! expected: %+v, actual: %+v\n", want, result)
	}
}

func TestGameOverByWall(t *testing.T) {
	snk := NewSnake(Down, []Position{{1, 1}})
	stg := NewStage(3, 3, snk)
	game := NewGame(stg)
	snk.Move()

	want := true
	result := game.IsOver()
	if result != want {
		t.Errorf("invalid game over! expected: %+v, actual: %+v\n", want, result)
	}
}

func TestGameOverBySnakeBody(t *testing.T) {
	snk := NewSnake(Down, []Position{
		{1, 1},
		{1, 2},
		{2, 2},
		{2, 1},
		{1, 1},
	})
	stg := NewStage(3, 3, snk)
	game := NewGame(stg)
	want := true
	result := game.IsOver()
	if result != want {
		t.Errorf("invalid game over! expected: %+v, actual: %+v\n", want, result)
	}
}

func TestGameNotOverBySnakeMove(t *testing.T) {
	snk := NewSnake(Down, []Position{
		{1, 1},
		{1, 2},
		{1, 3},
		{1, 4},
	})
	stg := NewStage(10, 10, snk)
	game := NewGame(stg)
	snk.Move()

	want := false
	result := game.IsOver()

	if result != want {
		t.Errorf("invalid game over! expected: %+v, actual: %+v\n", want, result)
	}
}

func TestGameOverBySnakeMove(t *testing.T) {
	snk := NewSnake(Left, []Position{
		{1, 1},
		{1, 2},
		{1, 3},
		{2, 3},
		{2, 2},
	})
	stg := NewStage(10, 10, snk)
	game := NewGame(stg)

	want := false
	result := game.IsOver()

	if result != want {
		t.Errorf("invalid game over before move! expected: %+v, actual: %+v\n", want, result)
	}

	snk.Move()

	want = true
	result = game.IsOver()

	if result != want {
		t.Errorf("invalid game over after move! expected: %+v, actual: %+v\n", want, result)
	}
}

func TestGamePlay(t *testing.T) {
	snk := NewSnake(Down, []Position{
		{1, 1},
		{1, 2},
		{1, 3},
		{1, 4},
	})
	stg := NewStage(10, 10, snk)
	game := NewGame(stg)
	fmt.Println("1", game.IsOver()) // debug
	snk.Move()
	fmt.Println("2", game.IsOver()) // debug
	stg.SnakeEatsFood()
	fmt.Println("3", game.IsOver()) // debug
	stg.PlaceFood()

	want := false
	result := game.IsOver()

	if result != want {
		t.Errorf("invalid game over after the first move! expected: %+v, actual: %+v\n", want, result)
	}

	snk.ChangeDirection(Right)
	snk.Move()
	stg.SnakeEatsFood()
	stg.PlaceFood()

	want = false
	result = game.IsOver()

	if result != want {
		t.Errorf("invalid game over after the second move! expected: %+v, actual: %+v\n", want, result)
	}
}
