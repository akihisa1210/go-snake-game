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
	section   []Position
}

func NewSnake(d Direction, p Position) *Snake {
	return &Snake{
		direction: d,
		section: []Position{
			p,
		},
	}
}

func (s *Snake) GetCurrentHeadPosition() Position {
	return s.section[0]
}

func (s *Snake) Move() {
	switch s.direction {
	case Up:
		s.section[0].Y--
	case Right:
		s.section[0].X++
	case Down:
		s.section[0].Y++
	case Left:
		s.section[0].X--
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

func (s *Snake) addSection(p Position) {
	s.section = append(s.section, p)
}
