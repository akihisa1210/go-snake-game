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
	section   []Position // 末尾をスネークの先頭とする。
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
	return s.section[len(s.section)-1]
}

func (s *Snake) Move() {
	next := s.GetCurrentHeadPosition()

	switch s.direction {
	case Up:
		next.Y--
	case Right:
		next.X++
	case Down:
		next.Y++
	case Left:
		next.X--
	}

	// 末尾に next を追加し、先頭の要素を削除する
	s.section = append(s.section[1:], next)
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
