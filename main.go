package main

import (
	"time"

	"github.com/akihisa1210/go-snake-game/snake"
	"github.com/nsf/termbox-go"
)

var (
	height       = 15
	witdh        = 50
	defaultColor = termbox.ColorDefault
)

var keyEvents = make(chan termbox.Key)

func update(s *snake.Snake) error {
	termbox.Clear(defaultColor, defaultColor)

	s.Move()

	// render stage and snake
	for y := 0; y < height; y++ {
		for x := 0; x < witdh; x++ {
			if x == s.GetCurrentPosition().X && y == s.GetCurrentPosition().Y {
				termbox.SetCell(x, y, 'O', defaultColor, defaultColor)
			} else {
				termbox.SetCell(x, y, '_', defaultColor, defaultColor)
			}
		}
	}
	return termbox.Flush()
}

func listenToKey(events chan termbox.Key) {
	for {
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			events <- ev.Key
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

	go listenToKey(keyEvents)

	s := snake.NewSnake(0, snake.Position{X: 15, Y: 5})

mainloop:
	for {
		select {
		case k := <-keyEvents:
			switch k {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeyArrowUp:
				s.ChangeDirection(snake.Up)
			case termbox.KeyArrowRight:
				s.ChangeDirection(snake.Right)
			case termbox.KeyArrowDown:
				s.ChangeDirection(snake.Down)
			case termbox.KeyArrowLeft:
				s.ChangeDirection(snake.Left)
			}
		default:
			update(s)
			time.Sleep(time.Second / 10)
		}
	}
}
