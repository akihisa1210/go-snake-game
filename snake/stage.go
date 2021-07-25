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
