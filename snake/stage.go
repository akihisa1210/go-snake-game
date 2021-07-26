package snake

import (
	"math/rand"
	"time"
)

type Stage struct {
	width  int
	height int
	snake  *Snake
	food   *Food
}

func NewStage(w int, h int, snake *Snake) *Stage {
	return &Stage{w, h, snake, nil}
}

func (stg *Stage) IsWall(p Position) bool {
	if (p.X <= 0 || p.X >= stg.width-1) || (p.Y <= 0 || p.Y >= stg.height-1) {
		return true
	}
	return false
}

func (stg *Stage) GetHeight() int {
	return stg.height
}

func (stg *Stage) GetWidth() int {
	return stg.width
}

type Food struct {
	position Position
}

func NewFood(p Position) *Food {
	return &Food{position: p}
}

func (f *Food) Where() Position {
	return f.position
}

func (stg *Stage) PlaceFood() {
	if stg.food != nil {
		return
	}

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(stg.GetWidth()-2) + 1
	y := rand.Intn(stg.GetHeight()-2) + 1
	f := NewFood(Position{X: x, Y: y})

	stg.food = f
}

func (stg *Stage) IsFood(p Position) bool {
	if stg.food != nil && stg.food.Where() == p {
		return true
	}
	return false
}

func (stg *Stage) SnakeEatsFood() {
	if stg.food != nil && stg.snake.GetCurrentHeadPosition() != stg.food.Where() {
		return
	}
	stg.snake.addSection(stg.snake.GetCurrentHeadPosition())
	stg.food = nil
}

func (stg *Stage) IsSnake(p Position) bool {
	for _, s := range stg.snake.sections {
		if s == p {
			return true
		}
	}
	return false
}
