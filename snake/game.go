package snake

type Game struct {
	stage *Stage
}

func NewGame(stg *Stage) *Game {
	return &Game{stage: stg}
}

func (g *Game) IsOver() bool {
	// スネークの現在位置が壁かどうか（true ならゲームオーバー）
	return g.stage.IsWall(g.stage.snake.GetCurrentPosition())
}
