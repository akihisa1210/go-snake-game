package main

import (
	"time"

	"github.com/akihisa1210/go-snake-game/snake"
	"github.com/nsf/termbox-go"
)

var (
	defaultColor = termbox.ColorDefault
)

type EventType int

const (
	Key EventType = iota
	GameOver
)

// Event is key input or game over.
// If key input event, its content is termbox.Key.
// If game over event, its content is nil.
type Event struct {
	kind    EventType
	content interface{}
}

var events = make(chan Event)

func update(g *snake.Game, stg *snake.Stage, snk *snake.Snake) error {
	termbox.Clear(defaultColor, defaultColor)

	snk.Move()
	stg.PlaceFood()

	// render stage, snake, and food
	for y := 0; y < stg.GetHeight(); y++ {
		for x := 0; x < stg.GetWidth(); x++ {
			p := snake.Position{X: x, Y: y}
			if x == snk.GetCurrentPosition().X && y == snk.GetCurrentPosition().Y {
				termbox.SetCell(x, y, 'O', defaultColor, defaultColor)
			} else if stg.IsWall(p) {
				termbox.SetCell(x, y, 'X', defaultColor, defaultColor)
			} else if stg.IsFood(p) {
				termbox.SetCell(x, y, '@', defaultColor, defaultColor)
			} else {
				termbox.SetCell(x, y, '-', defaultColor, defaultColor)
			}
		}
	}

	return termbox.Flush()
}

func listenToKey(events chan Event) {
	for {
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			events <- Event{kind: Key, content: ev.Key}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	go listenToKey(events)

	snk := snake.NewSnake(0, snake.Position{X: 15, Y: 5})
	stg := snake.NewStage(50, 10, snk)
	game := snake.NewGame(stg)

mainloop:
	for {
		select {
		case e := <-events:
			{
				switch e.kind {
				case Key:
					switch e.content {
					case termbox.KeyEsc:
						break mainloop
					case termbox.KeyArrowUp:
						snk.ChangeDirection(snake.Up)
					case termbox.KeyArrowRight:
						snk.ChangeDirection(snake.Right)
					case termbox.KeyArrowDown:
						snk.ChangeDirection(snake.Down)
					case termbox.KeyArrowLeft:
						snk.ChangeDirection(snake.Left)
					}
				case GameOver:
					break mainloop
				}
			}
		default:
			// If you break mainloop in update(), terminal hangs.
			if game.IsOver() {
				break mainloop
			}
			update(game, stg, snk)
			time.Sleep(time.Second / 10)
		}
	}
}
