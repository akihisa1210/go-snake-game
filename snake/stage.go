package snake

type Stage struct {
	width  int
	height int
	snake  *Snake
}

func NewStage(w int, h int, snake *Snake) *Stage {
	return &Stage{w, h, snake}
}

func (stg *Stage) IsWall(p Position) bool {
	if (p.X <= 0 || p.X >= stg.width) || (p.Y <= 0 || p.Y >= stg.height) {
		return true
	}
	return false
}
