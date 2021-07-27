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
	sections  []Position // 末尾をスネークの先頭とする。
	length    int
}

func NewSnake(d Direction, s []Position) *Snake {
	return &Snake{
		direction: d,
		sections:  s,
		length:    len(s),
	}
}

func (s *Snake) GetCurrentHeadPosition() Position {
	return s.sections[len(s.sections)-1]
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

	if s.length != len(s.sections) {
		// フードを食べたあと、最初の移動
		// sections 末尾（スネーク先頭）に next を追加するだけ
		s.sections = append(s.sections, next)
	} else {
		// フードを食べていない場合の移動
		// sections 末尾（スネーク先頭）に next を追加し、section 先頭（スネーク末尾）の要素を削除する
		s.sections = append(s.sections[1:], next)
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
	s.length++
}
