package snake

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Position struct {
	X int
	Y int
}

type Snake struct {
	direction Direction
	section   Position
}

func NewSnake(d Direction, p Position) *Snake {
	return &Snake{
		direction: d,
		section:   p,
	}
}

func (s *Snake) GetCurrentPosition() Position {
	return s.section
}

func (s *Snake) Move() {
	switch s.direction {
	case Up:
		s.section.Y--
	case Right:
		s.section.X++
	case Down:
		s.section.Y++
	case Left:
		s.section.X--
	}
}

func (s *Snake) ChangeDirection(d Direction) {
	if s.direction == d || s.isOppositeDirection(d) {
		return
	}
	s.direction = d
}

func (s *Snake) isOppositeDirection(target Direction) bool {
	switch s.direction {
	case Up:
		if target == Down {
			return true
		}
	case Right:
		if target == Left {
			return true
		}
	case Down:
		if target == Up {
			return true
		}
	case Left:
		if target == Right {
			return true
		}
	}
	return false
}
