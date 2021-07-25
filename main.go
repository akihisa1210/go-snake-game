package main

import (
	"time"

	"github.com/nsf/termbox-go"
)

var (
	height       = 15
	witdh        = 50
	defaultColor = termbox.ColorDefault
)

var keyEvents = make(chan termbox.Key)

func update(s *Snake) error {
	termbox.Clear(defaultColor, defaultColor)

	s.move()

	// render stage and snake
	for y := 0; y < height; y++ {
		for x := 0; x < witdh; x++ {
			if x == s.section.x && y == s.section.y {
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

	s := NewSnake(0, Position{15, 5})

mainloop:
	for {
		select {
		case k := <-keyEvents:
			switch k {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeyArrowUp:
				s.changeDirection(Up)
			case termbox.KeyArrowRight:
				s.changeDirection(Right)
			case termbox.KeyArrowDown:
				s.changeDirection(Down)
			case termbox.KeyArrowLeft:
				s.changeDirection(Left)
			}
		default:
			update(s)
			time.Sleep(time.Second / 10)
		}
	}
}
